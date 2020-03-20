package helloworld

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/spf13/viper"
)

// Modifies input string 'name' to a 'Hello World!' pattern
func GetHelloWorldMessage(name string) string {
	if name == "" {
		if viper.IsSet("cmd.helloworld.name.default") {
			log.Debug().Msg("Falling back to default 'name' from configuration")
			name = viper.GetString("cmd.helloworld.name.default")
		} else {
			log.Info().Msg("Name is required")
			name = "world"
		}
	}

	return fmt.Sprintf("Hello %s!", strings.Title(name))
}
