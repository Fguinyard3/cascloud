package main

import (
	"cascloud/config"
	"cascloud/db"
	"cascloud/routes"
	"cascloud/storage"
	"context"
	"fmt"
	"net/http"

	s3cfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		panic(err)
	}

	dbInstance, err := db.DB(cfg)
	if err != nil {
		panic(err)
	}

	sqlDB, err := dbInstance.DB()
	if err != nil {
		fmt.Println("Error getting underlying DB:", err)
		return
	}
	defer sqlDB.Close()

	s3config, err := s3cfg.LoadDefaultConfig(context.TODO(), s3cfg.WithRegion(cfg.S3Region))
	if err != nil {
		log.Error().Err(err).Msg("Error loading S3 config")
		return
	}

	s3Service := storage.S3Client{
		Client:     s3.NewFromConfig(s3config),
		BucketName: cfg.S3BucketName,
	}

	handler := &routes.HandlerClient{
		DBClient: db.NewClient(dbInstance),
		S3Client: &s3Service,
	}

	// Define routes and handlers here
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/users", handler.RegisterUser)
	e.POST("/login", handler.LoginUser)
	e.POST("/upload", handler.UploadFile)
	e.POST("/create-folder", handler.CreateFolder)
	e.GET("/get-directory", handler.GetDirectory)
	e.GET("/get-workspaces", handler.GetUsersWorkspaces)
	e.GET("/get-files", handler.GetFilesByFolderID)
	e.GET("/download", handler.DownloadFile)
	e.GET("/get-user", handler.GetUser)

	// Start the Echo server
	e.Start(":8080")

}
