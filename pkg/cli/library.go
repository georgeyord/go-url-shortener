package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/logrusorgru/aurora"
	"github.com/rs/zerolog/log"
)

func GetInput(label string) (name string) {
	PrintMessage(fmt.Sprintf("Please enter your %s: ", label), aurora.Cyan)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		return strings.TrimSpace(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal().Err(err).Msg("Reading standard input failed")
	}
	return
}

func PrintMessage(message string, auroraFn func(arg interface{}) aurora.Value) {
	fmt.Println(auroraFn(message))
}
