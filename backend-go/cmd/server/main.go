package main

import (
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	//"gorm.io/gorm/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/fmmaks666/yoevent-backend/internal/api"
	"github.com/fmmaks666/yoevent-backend/internal/models"
)

func main() {
	db, err := gorm.Open(sqlite.Open("yo.db"), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Failed to open the DB")
	}

	err = godotenv.Load()
	if err != nil {
		panic("Failed to open .env file")
	}
	adminPass := os.Getenv("ADMIN_PASSWORD")
	frontendUrl := os.Getenv("FRONTEND_URL")
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	salt := os.Getenv("SALT")

	models.Setup(db)

	// ctx := context.Background()
	// TODO: Make a config?
	r := api.Setup(db, adminPass, frontendUrl, salt)
	fmt.Println("Listening...")
	r.Run()
}
