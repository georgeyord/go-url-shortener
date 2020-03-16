package actions

import (
	"net/http"

	"github.com/georgeyord/go-url-shortener/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateUrlPairInput struct {
	Short string `json:"short"`
	Long  string `json:"long" binding:"required"`
}

type UpdateUrlPairInput struct {
	Short string `json:"short"`
	Long  string `json:"long"`
}

func CreateUrlPair(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input CreateUrlPairInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create UrlPair
	urlPair := models.UrlPair{Short: input.Short, Long: input.Long}
	if err := db.Create(&urlPair).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": urlPair})
}

func ListUrlPairs(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var urlPairs []models.UrlPair
	if err := db.Find(&urlPairs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": urlPairs})
}

func Redirect(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var urlPair models.UrlPair
	short := c.Param("short")
	if err := db.Where("short = ?", short).First(&urlPair).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusPermanentRedirect, urlPair.Long)
}
