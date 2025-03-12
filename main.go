package main

import (
	"catalogue/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connectDB(dbUrl string) {
	var err error
	DB, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	log.Println("Connected to the database successfully!")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	dbUrl := os.Getenv("DB_URL")
	port := os.Getenv("PORT")

	connectDB(dbUrl)

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	routes.Routes(router, DB)
	router.Run(":" + port)
}
