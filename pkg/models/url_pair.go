package models

import "github.com/jinzhu/gorm"

type UrlPair struct {
	gorm.Model
	Short string
	Long  string
}

/*
 * Just called New because external users will
 * use urlPair.New()
 */
func New(short string, long string) *UrlPair {
	urlPair := new(UrlPair)
	urlPair.Short = short
	urlPair.Long = long
	return urlPair
}
