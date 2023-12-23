package models

import (
	"cascloud/types"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type User struct {
	ID           uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	FirstName    string    `json:"first_name" gorm:"not null"`
	LastName     string    `json:"last_name" gorm:"not null"`
	Email        string    `json:"email" gorm:"not null;unique"`
	UserName     string    `json:"username" gorm:"not null;unique"`
	PasswordHash string    `json:"password" gorm:"not null"`
	// why is uuid.UUID not working here?
	Workspaces pq.StringArray  `json:"workspaces" gorm:"type:uuid[]"`
	CreatedAt  types.Timestamp `json:"created_at" gorm:"type:timestamptz;autoCreateTime"`
}

type Workspace struct {
	ID           uuid.UUID       `json:"id" gorm:"primaryKey;type:uuid;"`
	HomeFolderID uuid.UUID       `json:"home_folder_id" gorm:"not null"`
	Name         string          `json:"name" gorm:"not null"`
	OwnerID      uuid.UUID       `json:"owner_id" gorm:"not null"`
	Users        pq.StringArray  `json:"users" gorm:"type:uuid[]"`
	CreatedAt    types.Timestamp `json:"created_at" gorm:"type:timestamptz;autoCreateTime"`
}

type Collaborations struct {
	ID          uuid.UUID       `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID      uuid.UUID       `json:"user_id" gorm:"not null"`
	RoleID      uuid.UUID       `json:"role_id"`
	WorkspaceID uuid.UUID       `json:"workspace_id"`
	CreatedAt   types.Timestamp `json:"created_at" gorm:"type:timestamptz;autoCreateTime"`
}

type Role struct {
	ID        uuid.UUID       `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name      string          `json:"name" gorm:"not null"`
	CreatedAt types.Timestamp `json:"created_at" gorm:"type:timestamptz;autoCreateTime"`
}

// Hierarchical file system
type File struct {
	ID        uuid.UUID       `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name      string          `json:"name" gorm:"not null"`
	Path      string          `json:"path" gorm:"not null"`
	Size      int64           `json:"size" gorm:"not null"`
	X         float64         `json:"x" gorm:"not null"`
	Y         float64         `json:"y" gorm:"not null"`
	FolderID  uuid.UUID       `json:"folder_id" gorm:"not null"`
	CreatedAt types.Timestamp `json:"created_at" gorm:"type:timestamptz;autoCreateTime"`
}

type Folder struct {
	ID          uuid.UUID       `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name        string          `json:"name" gorm:"not null"`
	WorkspaceID uuid.UUID       `json:"workspace_id" gorm:"not null"`
	ParentID    uuid.UUID       `json:"parent_id" gorm:"not null"`
	X           float64         `json:"x" gorm:"not null"`
	Y           float64         `json:"y" gorm:"not null"`
	CreatedAt   types.Timestamp `json:"created_at" gorm:"type:timestamptz;autoCreateTime"`
	Children    []uuid.UUID     `json:"children" gorm:"type:uuid[]"`
	Path        string          `json:"path" gorm:"not null"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateFolderRequest struct {
	Name        string `json:"name"`
	WorkspaceID string `json:"workspace_id"`
	ParentID    string `json:"parent_id"`
}

type EditFileRequest struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	FolderID string `json:"folder_id"`
	X        string `json:"x"`
	Y        string `json:"y"`
}
