package helloworld

import (
	"fmt"
	"strings"

	"github.com/georgeyord/go-scrumpoker-api/pkg/errors"
)

func GetHelloWorldMessage(name string) string {
	if name == "" {
		errors.Error("Name is required")
	}

	return fmt.Sprintf("Hello %s!", strings.Title(name))
}
