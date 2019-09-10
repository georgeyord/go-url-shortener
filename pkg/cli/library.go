package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/logrusorgru/aurora"
)

func GetInput(label string) (name string) {
	PrintMessage(fmt.Sprintf("Please enter your %s: ", label), aurora.Cyan)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		return strings.TrimSpace(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return
}

func PrintMessage(message string, auroraFn func(arg interface{}) aurora.Value) {
	fmt.Print(auroraFn(message))
}
