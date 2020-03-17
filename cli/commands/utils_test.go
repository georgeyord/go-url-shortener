package commands

import (
	"log"

	"github.com/georgeyord/go-url-shortener/pkg/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/spf13/viper"
)

// Grab the Gin router with registered routes
func initTestDb() *gorm.DB {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err.Error())
	}
	models.SetupModels(db)
	viper.Set("db", db)
	return db
}
