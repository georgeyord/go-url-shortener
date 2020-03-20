package commands

import (
	"testing"

	"github.com/georgeyord/go-url-shortener/pkg/test/cli"
	"github.com/rs/zerolog/log"

	"github.com/stretchr/testify/assert"
)

func TestHelloworldShouldReturnValidOutputWhenProvidedValidName(t *testing.T) {
	args := cli.ParseShellArgs("helloworld -n foo")
	rootCmd.SetArgs(args)

	captured := cli.CaptureOutput(func() {
		_, err := rootCmd.ExecuteC()
		if err != nil {
			log.Fatal().Err(err).Msg("")
		}
	})

	// Check message exists
	assert.Contains(t, captured, "Hello Foo!")
	// Check correct color hex exists
	assert.Contains(t, captured, "[35m")
}
