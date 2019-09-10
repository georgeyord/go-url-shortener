package errors

import (
	"errors"
	"fmt"
)

func Error(message string) {
	panic(errors.New(fmt.Sprintf("Error: %s", message)))
}
