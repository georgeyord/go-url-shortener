package common

import (
	"log"

	"github.com/georgeyord/go-url-shortener/pkg/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func InitTestDb() *gorm.DB {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err.Error())
	}
	models.SetupModels(db)
	return db
}
