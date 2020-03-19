package commands

import (
	"log"
	"testing"

	"github.com/georgeyord/go-url-shortener/pkg/models"
	"github.com/georgeyord/go-url-shortener/pkg/test/cli"
	"github.com/georgeyord/go-url-shortener/pkg/test/common"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestListShouldReturnValidOutputWhenNoUrlPairsAreFound(t *testing.T) {
	db := common.InitTestDb()
	viper.Set("db", db)

	args := cli.ParseShellArgs("list")
	rootCmd.SetArgs(args)

	captured := cli.CaptureOutput(func() {
		_, err := rootCmd.ExecuteC()
		if err != nil {
			log.Fatal(err)
		}
	})

	// Checek message exists
	assert.Contains(t, captured, "No records found!")
	// Check correct color hex exists
	assert.Contains(t, captured, "[33m")
}

func TestListShouldReturnValidOutputWhenAUrlPairIsFound(t *testing.T) {
	db := common.InitTestDb()
	viper.Set("db", db)

	urlPair := models.NewUrlPair("http://www.google.com", "123")
	if err := db.Create(&urlPair).Error; err != nil {
		log.Fatal(err.Error())
	}

	args := cli.ParseShellArgs("list")
	rootCmd.SetArgs(args)

	captured := cli.CaptureOutput(func() {
		_, err := rootCmd.ExecuteC()
		if err != nil {
			log.Fatal(err)
		}
	})

	// Checek message exists
	assert.Contains(t, captured, "http://www.google.com")
	assert.Contains(t, captured, "123")
	// Check correct color hex exists
	assert.Contains(t, captured, "[36m")
}
