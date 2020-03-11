package models

import (
	"github.com/jinzhu/gorm"
)

func SetupModels(db *gorm.DB) {
	db.AutoMigrate(&UrlPair{})
}
