package actions

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/georgeyord/go-url-shortener/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	kafkalib "github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
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

	defer writeStats(c, urlPair)
	c.Redirect(http.StatusPermanentRedirect, urlPair.Long)
}

func writeStats(c *gin.Context, urlPair models.UrlPair) {
	type Message struct {
		Short string `json:"short"`
		Long  string `json:"long"`
		Host  string `json:"host"`
	}

	var message Message
	headers := []kafkalib.Header{
		kafkalib.Header{
			Key:   "requestURI",
			Value: []byte(c.Request.RequestURI),
		},
	}

	statsTopic := viper.GetString("kafka.writers.stats.topic")
	writerValue, writerExists := c.Get(statsTopic)
	if writerExists {
		writer := writerValue.(*kafkalib.Writer)

		message.Short = urlPair.Short
		message.Long = urlPair.Long
		message.Host = c.Request.Host

		jsonMessage, errJson := json.Marshal(message)
		if errJson != nil {
			log.Println(errJson)
			return
		}

		errWriter := writer.WriteMessages(context.Background(), kafkalib.Message{
			Headers: headers,
			Value:   jsonMessage,
		})
		if errWriter != nil {
			log.Println(errWriter)
		}
	}
}
