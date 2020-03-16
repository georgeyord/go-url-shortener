package models

import (
	"errors"
	"regexp"

	"github.com/georgeyord/go-url-shortener/pkg/urlshort"

	"github.com/jinzhu/gorm"
)

type UrlPair struct {
	gorm.Model
	Short string `json:"short"`
	Long  string `json:"long"`
}

func (urlPair *UrlPair) BeforeSave(db *gorm.DB) (err error) {
	matched, _ := regexp.MatchString(`^https?://.`, urlPair.Long)
	if !matched {
		err = errors.New("Long url should start with 'http://' or 'https://'!")
		return err
	}

	if len(urlPair.Short) < 1 {
		urlPair.Short = urlshort.GenerateRandomShortUrl()
	}

	if db.NewRecord(urlPair) {
		exists, err := ShortExists(urlPair.Short, db)
		if err != nil {
			return err
		} else if exists {
			err = errors.New("Short url is already in use!")
			return err
		}
	}

	return
}

/*
 * Just called New because external users will
 * use urlPair.New()
 */
func NewUrlPair(long string, short string) *UrlPair {
	urlPair := new(UrlPair)
	urlPair.Short = short
	urlPair.Long = long
	return urlPair
}

func ShortExists(short string, db *gorm.DB) (bool, error) {
	var count int
	err := db.Model(&UrlPair{}).Where("short = ?", short).Count(&count).Error
	if err != nil {
		return true, err
	}

	return count > 0, nil
}
