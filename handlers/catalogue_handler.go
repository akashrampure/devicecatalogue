package handlers

import (
	"catalogue/cache"
	"catalogue/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

func CreateCatalogue(c *gin.Context) {
	var catalogue models.DeviceCatalogue
	if err := c.ShouldBindJSON(&catalogue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := DB.Omit("slno").Create(&catalogue)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	message := fmt.Sprintf("%s Catalogue created successfully", catalogue.SkuCode)
	c.JSON(http.StatusCreated, gin.H{"message": message})
}

func GetCatalogue(c *gin.Context) {
	skuCode := c.Query("skucode")

	if skuCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Sku code is required"})
		return
	}

	cacheData, err := cache.GetCache(skuCode)
	if err != nil {
		fmt.Println("Cache not found for sku code:", skuCode)
	}
	if cacheData != nil {
		c.JSON(http.StatusOK, cacheData)
		return
	}

	var catalogue models.DeviceCatalogue
	result := DB.First(&catalogue, "skucode = ?", skuCode)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Catalogue not found"})
		return
	}

	cache.SetCache(skuCode, catalogue)
	c.JSON(http.StatusOK, catalogue)
}

func GetAllCatalogue(c *gin.Context) {
	cacheKey := "getallcatalogue"
	cacheData, err := cache.GetCache(cacheKey)
	if err != nil {
		fmt.Println("Cache not found for:", cacheKey)
	}
	if cacheData != nil {
		c.JSON(http.StatusOK, cacheData)
		return
	}

	var catalogues []models.DeviceCatalogue
	result := DB.Find(&catalogues)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No catalogues found"})
		return
	}

	cache.SetCache(cacheKey, catalogues)
	c.JSON(http.StatusOK, catalogues)
}

func UpdateCatalogue(c *gin.Context) {
	skuCode := c.Query("skucode")

	if skuCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Sku code is required"})
		return
	}

	var catalogue models.DeviceCatalogue
	result := DB.First(&catalogue, "skucode = ?", skuCode)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Catalogue not found"})
		return
	}

	if err := c.ShouldBindJSON(&catalogue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result = DB.Model(&catalogue).Where("skucode = ?", skuCode).Omit("slno", "skucode").Updates(catalogue)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	message := fmt.Sprintf("%s Catalogue updated successfully", skuCode)
	c.JSON(http.StatusOK, gin.H{"message": message})
}

func DeleteCatalogue(c *gin.Context) {
	skuCode := c.Query("skucode")
	if skuCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Sku code is required"})
		return
	}

	var catalogue models.DeviceCatalogue
	result := DB.First(&catalogue, "skucode = ?", skuCode)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Catalogue not found"})
		return
	}

	result = DB.Delete(&catalogue, "skucode = ?", skuCode)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	ResetSequence()

	message := fmt.Sprintf("%s Catalogue deleted successfully", skuCode)
	c.JSON(http.StatusOK, gin.H{"message": message})
}

func ResetSequence() error {
	if err := DB.Exec(`
		SELECT setval(
			pg_get_serial_sequence('"intellicar"."devicecatalogue"', 'slno'), 
			COALESCE(MAX(slno), 1)
		) FROM "intellicar"."devicecatalogue";
	`).Error; err != nil {
		return err
	}
	return nil
}
