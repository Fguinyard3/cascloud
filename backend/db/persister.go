package db

import (
	"cascloud/models"
)

type DBInterface interface {
	CreateUser(user *models.User) error
	UserExists(email string) bool
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
	CreateWorkspace(workspace *models.Workspace, user *models.User) error
	CreateCollaboration(collaboration *models.Collaborations) error
	CreateFolder(folder *models.Folder) error
	CreateFile(file *models.File) error
	EditFile(file *models.File) error
	GetFolderByID(id string) (*models.Folder, error)
	GetFolderPath(folder *models.Folder) (string, error)
	GetFilesByFolderID(folderID string) ([]models.File, error)
	GetFileByID(fileID string) (*models.File, error)
	GetFoldersAndFilesInFolder(folderID string) ([]models.Folder, []models.File, error)
	GetWorkspaceByID(id string) (*models.Workspace, error)
	GetWorkspacesAvailableWorkspaces(userID string) (*[]models.Workspace, error)
}
