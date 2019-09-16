package helloworld

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

// Modifies input string 'name' to a 'Hello World!' pattern
func GetHelloWorldMessage(name string) string {
	if name == "" {
		if viper.IsSet("cmd.name.default") {
			log.Print("Falling back to default 'name' from configuration")
			name = viper.GetString("cmd.name.default")
		} else {
			log.Print("Name is required")
			name = "world"
		}
	}

	return fmt.Sprintf("Hello %s!", strings.Title(name))
}
