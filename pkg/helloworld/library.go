package helloworld

import (
	"fmt"
	"log"
	"strings"
)

func GetHelloWorldMessage(name string) string {
	const msgEmptyName = "Name is required"

	if name == "" {
		log.Print(msgEmptyName)
		name = "world"
	}

	return fmt.Sprintf("Hello %s!", strings.Title(name))
}
