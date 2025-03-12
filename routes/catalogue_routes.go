package routes

import (
	"catalogue/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(router *gin.Engine, db *gorm.DB) {
	handlers.DB = db

	router.POST("/catalogue", handlers.CreateCatalogue)
	router.GET("/getcatalogue", handlers.GetCatalogue)
	router.GET("/catalogue", handlers.GetAllCatalogue)
	router.PUT("/catalogue", handlers.UpdateCatalogue)
	router.DELETE("/catalogue", handlers.DeleteCatalogue)
}
