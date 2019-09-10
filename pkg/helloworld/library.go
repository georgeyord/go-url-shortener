package helloworld

import (
	"fmt"
	"log"
	"strings"
)

func GetHelloWorldMessage(name string) string {
	if name == "" {
		log.Fatal("Name is required")
	}

	return fmt.Sprintf("Hello %s!", strings.Title(name))
}
