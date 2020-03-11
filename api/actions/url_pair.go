package actions

import (
	"net/http"
	"regexp"

	"github.com/georgeyord/go-url-shortener/pkg/models"
	"github.com/georgeyord/go-url-shortener/pkg/urlshort"
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
	// Authenticate
	// if json.User != "manu" || json.Password != "123" {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
	// 	return
	// }
	// c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})

	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input CreateUrlPairInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Move this to BeforeSave gorm hook
	matched, _ := regexp.MatchString(`^https?://.`, input.Long)
	if !matched {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Long url should start with 'http://' or 'https://'!"})
		return
	}
	if len(input.Short) < 1 {
		input.Short = urlshort.GenerateRandomShortUrl()
	}

	// Create UrlPair
	urlPair := models.UrlPair{Short: input.Short, Long: input.Long}
	db.Create(&urlPair)

	c.JSON(http.StatusOK, gin.H{"data": urlPair})
}

func FindUrlPairs(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var urlPairs []models.UrlPair
	db.Find(&urlPairs)

	c.JSON(http.StatusOK, gin.H{"data": urlPairs})
}

func Redirect(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var urlPair models.UrlPair
	short := c.Param("short")
	if err := db.Where("short = ?", short).First(&urlPair).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, urlPair.Long)
}
