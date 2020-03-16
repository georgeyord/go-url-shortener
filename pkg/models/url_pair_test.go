package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUrlPairWillCreateAUrlPairWithTheCorrectAttributes(t *testing.T) {
	const long = "bar"
	const short = "foo"
	got := NewUrlPair(long, short)

	assert.Equal(t, got.Short, short)
	assert.Equal(t, got.Long, long)
}

func TestBeforeSaveWhileCreatingNewModelShouldReturnNoErrorsOnValidModel(t *testing.T) {
	const long = "https://www.google.com"
	const short = "a7sy9d8"

	db := initTestDb()
	urlPair := NewUrlPair(long, short)
	if err := db.Create(&urlPair).Error; err != nil {
		panic(err.Error())
	}

	assert.Equal(t, long, urlPair.Long)
	assert.Equal(t, short, urlPair.Short)
	assert.NotEmpty(t, urlPair.ID)
	assert.NotEmpty(t, urlPair.CreatedAt)
	assert.NotEmpty(t, urlPair.UpdatedAt)
}

func TestBeforeSaveWhileCreatingNewModelShouldCreateRandomShortUrlWhenShortUrlIsEmpty(t *testing.T) {
	const long = "https://www.google.com"
	var short string

	db := initTestDb()
	urlPair := NewUrlPair(long, short)
	if err := db.Create(&urlPair).Error; err != nil {
		panic(err.Error())
	}

	assert.Equal(t, long, urlPair.Long)
	assert.NotEmpty(t, urlPair.Short)
	assert.NotEmpty(t, urlPair.ID)
	assert.NotEmpty(t, urlPair.CreatedAt)
	assert.NotEmpty(t, urlPair.UpdatedAt)
}

func TestBeforeSaveWhileCreatingNewModelShouldReturnErrorWhenLongUrlIsEmpty(t *testing.T) {
	var long string
	var short string

	db := initTestDb()
	urlPair := NewUrlPair(long, short)
	err := db.Create(&urlPair).Error

	assert.NotEmpty(t, err)
	assert.Contains(t, err.Error(), "Long url should start with 'http://' or 'https://'!")
}

func TestBeforeSaveWhileCreatingNewModelShouldReturnErrorWhenShortUrlAlreadyExists(t *testing.T) {
	const long = "https://www.google.com"
	const short = "a7sy9d8"

	db := initTestDb()
	urlPair1 := NewUrlPair(long, short)
	if err := db.Create(&urlPair1).Error; err != nil {
		panic(err.Error())
	}

	urlPair2 := NewUrlPair(long, short)
	err := db.Create(&urlPair2).Error

	assert.NotEmpty(t, err)
	assert.Contains(t, err.Error(), "Short url is already in use!")
}

func TestBeforeSaveWhileCreatingNewModelShouldReturnSuccessfullyWhenShortUrlAlreadyExistsButWePerformAnUpdateOnThatSpecificRecord(t *testing.T) {
	const long = "https://www.google.com"
	const short = "a7sy9d8"

	db := initTestDb()
	urlPair := NewUrlPair(long, short)
	if err := db.Create(&urlPair).Error; err != nil {
		panic(err.Error())
	}

	assert.Equal(t, long, urlPair.Long)
	assert.NotEmpty(t, urlPair.Short)
	assert.NotEmpty(t, urlPair.ID)
	assert.NotEmpty(t, urlPair.CreatedAt)
	assert.NotEmpty(t, urlPair.UpdatedAt)

	oldShort := urlPair.Short
	oldLong := urlPair.Long
	oldCreatedAt := urlPair.CreatedAt
	oldUpdatedAt := urlPair.UpdatedAt

	urlPair.Long = "https://www.google.de"
	if err := db.Save(&urlPair).Error; err != nil {
		panic(err.Error())
	}

	assert.Equal(t, oldShort, urlPair.Short)
	assert.NotEqual(t, oldLong, urlPair.Long)
	assert.Equal(t, oldCreatedAt, urlPair.CreatedAt)
	assert.NotEqual(t, oldUpdatedAt, urlPair.UpdatedAt)
}

func TestShortExistsWithExistingShortUrlShouldReturnTrue(t *testing.T) {
	db := initTestDb()
	urlPair := NewUrlPair("http://www.google.com", "123")
	if err := db.Create(&urlPair).Error; err != nil {
		panic(err.Error())
	}

	got, err := ShortExists("123", db)

	assert.Nil(t, err)
	assert.True(t, got)
}
