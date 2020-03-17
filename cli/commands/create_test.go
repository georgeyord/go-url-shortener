package commands

import (
	"log"
	"testing"

	"github.com/georgeyord/go-url-shortener/pkg/test/cli"
	"github.com/stretchr/testify/assert"
)

func TestCreateShouldReturnValidOutputWhenAValidUrlPairIsProvided(t *testing.T) {
	initTestDb()

	args := cli.ParseShellArgs("create -l http://www.google.com -s 123")
	rootCmd.SetArgs(args)

	captured := cli.CaptureOutput(func() {
		_, err := rootCmd.ExecuteC()
		if err != nil {
			log.Fatal(err)
		}
	})

	// Checek message exists
	assert.Contains(t, captured, "New short url created")
	assert.Contains(t, captured, "123")
	assert.Contains(t, captured, "Long url")
	assert.Contains(t, captured, "http://www.google.com")
}
