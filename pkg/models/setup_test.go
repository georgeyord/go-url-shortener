package models

import (
	"github.com/rs/zerolog/log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func initTestDb() *gorm.DB {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
	SetupModels(db)
	return db
}
