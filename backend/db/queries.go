package db

import (
	model "cascloud/models"
	"fmt"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type DBClient struct {
	gorm *gorm.DB
}

func NewClient(gormDB *gorm.DB) *DBClient {
	return &DBClient{gorm: gormDB}
}

func (c *DBClient) CreateUser(user *model.User) error {
	log.Info().Msg("Creating user")
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg("Error hashing password")
		return err
	}
	user.PasswordHash = string(bcryptPassword)

	return c.gorm.Create(user).Error
}

// a function to check if a user exists in the database
func (c *DBClient) UserExists(email string) bool {
	log.Info().Msg("Checking if user exists")
	var user model.User
	result := c.gorm.Where("email = ?", email).First(&user)
	return result.RowsAffected != 0
}

// a function to get a user by email
func (c *DBClient) GetUserByEmail(email string) (*model.User, error) {
	log.Info().Msg("Getting user by email")
	var user model.User
	err := c.gorm.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *DBClient) GetUserByID(id string) (*model.User, error) {
	log.Info().Msg("Getting user by id")
	var user model.User
	err := c.gorm.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *DBClient) CreateWorkspace(workspace *model.Workspace, user *model.User) error {
	log.Info().Msg("Creating workspace")
	err := c.gorm.Create(workspace).Error
	if err != nil {
		return err
	}
	// make a new uuid
	folderID := uuid.New()
	folder := &model.Folder{
		ID:          folderID,
		Name:        fmt.Sprintf("home@%s", user.UserName),
		ParentID:    uuid.Nil,
		WorkspaceID: workspace.ID,
		Path:        fmt.Sprintf("home@%s", user.UserName),
	}
	// create the folder but we need to get the folder Id thats created
	err = c.CreateFolder(folder)
	if err != nil {
		return err
	}
	workspace.HomeFolderID = folderID
	// update the workspace with the folder id
	workspace.Users = append(workspace.Users, user.ID.String())
	err = c.gorm.Model(workspace).Updates(map[string]interface{}{
		"Users":        gorm.Expr("ARRAY_APPEND(users, ?)", user.ID),
		"HomeFolderID": folderID,
	}).Error
	if err != nil {
		return err
	}

	// update the user with the workspace id
	user.Workspaces = append(user.Workspaces, workspace.ID.String())
	err = c.gorm.Model(user).Updates(map[string]interface{}{
		"Workspaces": gorm.Expr("ARRAY_APPEND(workspaces, ?)", workspace.ID),
	}).Error
	if err != nil {
		return err
	}

	return nil

}

func (c *DBClient) GetWorkspaceByID(id string) (*model.Workspace, error) {
	var workspace model.Workspace
	err := c.gorm.Where("id = ?", id).First(&workspace).Error
	if err != nil {
		return nil, err
	}
	return &workspace, nil
}

func (c *DBClient) GetWorkspacesAvailableWorkspaces(userID string) (*[]model.Workspace, error) {
	var workspaces []model.Workspace
	// first get the user
	user, err := c.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	// user.Workspaces is a slice of uuids
	err = c.gorm.Where("id = ANY(?)", user.Workspaces).Find(&workspaces).Error
	if err != nil {
		return nil, err
	}
	return &workspaces, nil
}

func (c *DBClient) CreateCollaboration(collaboration *model.Collaborations) error {
	return c.gorm.Create(collaboration).Error
}

func (c *DBClient) CreateFolder(folder *model.Folder) error {
	if folder.ParentID != uuid.Nil {
		// get the parent folder
		parentFolder, err := c.GetFolderByID(folder.ParentID.String())
		if err != nil {
			return err
		}
		folder.Path = parentFolder.Path + "/" + folder.Name
	} else {
		folder.Path = folder.Name
	}
	return c.gorm.Create(folder).Error
}

func (c *DBClient) CreateFile(file *model.File) error {
	files, err := c.GetFilesByFolderID(file.FolderID.String())
	if err != nil {
		return err
	}
	// check if the name already exists
	for _, f := range files {
		if f.Name == file.Name {
			// the name already exists
			file.Name = fmt.Sprintf("%s%d", file.Name, len(files))
		}
	}
	return c.gorm.Create(file).Error
}

// a function to edit a file
func (c *DBClient) EditFile(file *model.File) error {
	return c.gorm.Model(file).Updates(map[string]interface{}{
		"Name":     file.Name,
		"Path":     file.Path,
		"X":        file.X,
		"Y":        file.Y,
		"FolderID": file.FolderID,
	}).Error
}

// a function to get a folder by id
func (c *DBClient) GetFolderByID(id string) (*model.Folder, error) {
	log.Info().Msg("Getting folder by id")
	var folder model.Folder
	err := c.gorm.Where("id = ?", id).First(&folder).Error
	if err != nil {
		return nil, err
	}
	return &folder, nil
}

// a function to get the path of a folder
func (c *DBClient) GetFolderPath(folder *model.Folder) (string, error) {
	err := c.gorm.Where("id = ?", folder.ID).First(&folder).Error
	if err != nil {
		return "", err
	}
	return folder.Path, nil
}

// A function to get the files by folder id
func (c *DBClient) GetFilesByFolderID(folderID string) ([]model.File, error) {
	var files []model.File
	err := c.gorm.Where("folder_id = ?", folderID).Find(&files).Error
	if err != nil {
		return nil, err
	}
	if len(files) == 0 {
		return files, fmt.Errorf("no files found in folder")
	}
	fmt.Println(files)
	return files, nil
}

// a function to get a file by id
func (c *DBClient) GetFileByID(id string) (*model.File, error) {
	var file model.File
	err := c.gorm.Where("id = ?", id).First(&file).Error
	if err != nil {
		return nil, err
	}
	return &file, nil
}

func (c *DBClient) GetFoldersAndFilesInFolder(folderID string) ([]model.Folder, []model.File, error) {
	var folders []model.Folder
	var files []model.File
	err := c.gorm.Where("parent_id = ?", folderID).Find(&folders).Error
	if err != nil {
		return nil, nil, err
	}
	err = c.gorm.Where("folder_id = ?", folderID).Find(&files).Error
	if err != nil {
		return nil, nil, err
	}
	return folders, files, nil
}

// a fuction that will get a workspace a find out that w
