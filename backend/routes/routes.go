package routes

import (
	"cascloud/db"
	"cascloud/helpers"
	"cascloud/models"
	"cascloud/storage"

	"context"
	"fmt"
	"strconv"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

type HandlerClient struct {
	DBClient db.DBInterface
	S3Client storage.S3Interface
}

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

type Handler interface {
	RegisterUser(c echo.Context) error
	LoginUser(c echo.Context) error
	UploadFile(c echo.Context) error
	CreateFolder(c echo.Context) error
	GetDirectory(c echo.Context) error
	GetFilesByFolderID(c echo.Context) error
	GetWorkspace(c echo.Context) error
	GetUser(c echo.Context) error
}

func (h *HandlerClient) RegisterUser(c echo.Context) error {
	var user models.User
	bindErr := c.Bind(&user)
	if bindErr != nil {
		return bindErr
	}
	// check if the user already exists
	if h.DBClient.UserExists(user.Email) {
		return c.JSON(400, "User already exists")
	}

	createErr := h.DBClient.CreateUser(&user)
	if createErr != nil {
		return createErr
	}

	workspaceID := uuid.New()

	workspace := models.Workspace{
		ID:      workspaceID,
		Name:    fmt.Sprintf("%s's Workspace", user.FirstName),
		OwnerID: user.ID,
		Users:   pq.StringArray{user.ID.String()},
	}

	// create a workspace for the user
	workspaceErr := h.DBClient.CreateWorkspace(&workspace, &user)
	if workspaceErr != nil {
		return workspaceErr
	}

	return c.JSON(200, user)

}

// a function to login a user
func (h *HandlerClient) LoginUser(c echo.Context) error {
	var creds models.UserLogin
	bindErr := c.Bind(&creds)
	if bindErr != nil {
		return bindErr
	}
	// check if the user already exists
	if !h.DBClient.UserExists(creds.Email) {
		return c.JSON(400, "User does not exist")
	}
	// get the user from the database
	user, userErr := h.DBClient.GetUserByEmail(creds.Email)
	if userErr != nil {
		return userErr
	}
	log.Info().Msg("User found")
	// check if the password is correct
	if !helpers.ComparePasswords(user.PasswordHash, creds.Password) {
		return c.JSON(401, "Incorrect password")
	}

	jwtToken, jwtErr := helpers.GenerateJWT(*user)
	if jwtErr != nil {
		return jwtErr
	}
	return c.JSON(200, map[string]interface{}{
		"token": jwtToken,
		"user":  user,
	})

}

func (h *HandlerClient) GetUser(c echo.Context) error {
	userID := c.QueryParam("user_id")
	if userID == "" {
		log.Error().Msg("User ID not provided")
		return c.JSON(400, "User ID not provided")
	}
	// Get the user from the database
	user, userErr := h.DBClient.GetUserByID(userID)
	if userErr != nil {
		log.Error().Err(userErr).Msg("Error getting user from database")
		return c.JSON(400, "Error getting user from database")
	}

	return c.JSON(200, user)
}

// a function to create a folder
func (h *HandlerClient) CreateFolder(c echo.Context) error {
	// We need to get the user id and the folder name from the request body
	var (
		folderReq models.CreateFolderRequest
	)

	bindErr := c.Bind(&folderReq)
	if bindErr != nil {
		return bindErr
	}

	folder := models.Folder{
		Name:        folderReq.Name,
		WorkspaceID: uuid.MustParse(folderReq.WorkspaceID),
		ParentID:    uuid.MustParse(folderReq.ParentID),
	}

	// create the folder in the database
	folderErr := h.DBClient.CreateFolder(&folder)
	if folderErr != nil {
		return folderErr
	}

	return c.JSON(200, folder)
}

// a function to upload a file
func (h *HandlerClient) UploadFile(c echo.Context) error {
	// Get the file from the form data
	file, err := c.FormFile("file")
	// get the folder id from the form data
	if err != nil {
		log.Error().Err(err).Msg("Error getting file from form data")
		return c.JSON(400, "Error getting file from form data")
	}
	folderID := c.FormValue("folder_id")
	if folderID == "" {
		log.Error().Msg("Folder ID not provided")
		return c.JSON(400, "Folder ID not provided")
	}

	fileData, openErr := file.Open()
	if openErr != nil {
		log.Error().Err(openErr).Msg("Error opening file")
		return c.JSON(400, "Error opening file")
	}
	// we need to calculate the size of the file and convert to a readable format in bytes
	fileSize := file.Size

	defer fileData.Close()

	// fileX, fileY := c.FormValue("x_coordinate"), c.FormValue("y_coordinate")
	// if fileX != "" && fileY != "" {
	// 	log.Error().Msg("No coordinate given")
	// 	return c.JSON(400, "No coordinate given")
	// }

	//we need to do the above but we need to convert the string form values to floats
	fileX, err := strconv.ParseFloat(c.FormValue("x_coordinate"), 64)
	if err != nil {
		log.Error().Err(err).Msg("Error converting x coordinate to float")
		return c.JSON(400, "Error converting x coordinate to float")
	}
	fileY, err := strconv.ParseFloat(c.FormValue("y_coordinate"), 64)
	if err != nil {
		log.Error().Err(err).Msg("Error converting y coordinate to float")
		return c.JSON(400, "Error converting y coordinate to float")
	}

	// Get the folder from the database
	folder, folderErr := h.DBClient.GetFolderByID(folderID)
	if folderErr != nil {
		log.Error().Err(folderErr).Msg("Error getting folder from database")
		return c.JSON(400, "Error getting folder from database")
	}

	path := fmt.Sprintf("%s/%s", folder.Path, file.Filename)
	fmt.Println(path)
	// Upload the file to s3
	uploadErr := h.S3Client.UploadFile(context.Background(), path, fileData)
	if uploadErr != nil {
		log.Error().Err(uploadErr).Msg("Error uploading file to s3")
		return c.JSON(400, "Error uploading file to s3")
	}

	// Create the file in the database
	fileModel := models.File{
		Name:     file.Filename,
		FolderID: folder.ID,
		Size:     fileSize,
		Path:     path,
		X:        fileX,
		Y:        fileY,
	}
	fileErr := h.DBClient.CreateFile(&fileModel)
	if fileErr != nil {
		log.Error().Err(fileErr).Msg("Error creating file in database")
		return c.JSON(400, "Error creating file in database")
	}

	return c.JSON(200, fileModel)
}

// a function to edit a file
func (h *HandlerClient) EditFile(c echo.Context) error {
	// The user will be able to edit the name , coordinates and folder of the file
	var (
		fileReq models.EditFileRequest
	)

	bindErr := c.Bind(&fileReq)
	if bindErr != nil {
		return bindErr
	}

	file := models.File{
		ID:       uuid.MustParse(fileReq.ID),
		Name:     fileReq.Name,
		FolderID: uuid.MustParse(fileReq.FolderID),
		X:        fileReq.X,
		Y:        fileReq.Y,
	}



// a function to get all files from a folder
func (h *HandlerClient) GetFilesByFolderID(c echo.Context) error {
	folderID := c.QueryParam("folder_id")
	if folderID == "" {
		log.Error().Msg("Folder ID not provided")
		return c.JSON(400, "Folder ID not provided")
	}
	// Get the files from the database
	files, filesErr := h.DBClient.GetFilesByFolderID(folderID)
	if filesErr != nil {
		log.Error().Err(filesErr).Msg("Error getting files from database")
		return c.JSON(400, "Error getting files from database")
	}

	return c.JSON(200, files)
}

func (h *HandlerClient) GetDirectory(c echo.Context) error {
	folderID := c.QueryParam("folder_id")
	if folderID == "" {
		log.Error().Msg("Folder ID not provided")
		return c.JSON(400, "Folder ID not provided")
	}
	// Get the folder from the database
	folders, files, folderErr := h.DBClient.GetFoldersAndFilesInFolder(folderID)
	if folderErr != nil {
		log.Error().Err(folderErr).Msg("Error getting folder from database")
		return c.JSON(400, "Error getting folder from database")
	}

	return c.JSON(200, map[string]interface{}{
		"folders": folders,
		"files":   files,
	})
}

func (h *HandlerClient) GetUsersWorkspaces(c echo.Context) error {
	userID := c.QueryParam("user_id")
	if userID == "" {
		log.Error().Msg("User ID not provided")
		return c.JSON(400, "User ID not provided")
	}
	// Get the workspaces from the database
	workspaces, workspacesErr := h.DBClient.GetWorkspacesAvailableWorkspaces(userID)
	if workspacesErr != nil {
		log.Error().Err(workspacesErr).Msg("Error getting workspaces from database")
		return c.JSON(400, "Error getting workspaces from database")
	}

	return c.JSON(200, workspaces)
}

// a function to download a file
func (h *HandlerClient) DownloadFile(c echo.Context) error {
	filePath := c.QueryParam("file_id")
	if filePath == "" {
		log.Error().Msg("File path not provided")
		return c.JSON(400, "File path not provided")
	}
	file, err := h.DBClient.GetFileByID(filePath)
	if err != nil {
		log.Error().Err(err).Msg("Error getting file from database")
		return c.JSON(400, "Error getting file from database")
	}
	// Download the file from s3
	fileData, fileErr := h.S3Client.DownloadFile(context.Background(), file.Path)
	if fileErr != nil {
		log.Error().Err(fileErr).Msg("Error downloading file from s3")
		return c.JSON(400, "Error downloading file from s3")
	}

	return c.Stream(200, "application/octet-stream", fileData)
}
