package commands

import (
	"github.com/spf13/cobra"

	"github.com/spf13/viper"

	"github.com/logrusorgru/aurora"

	"github.com/georgeyord/go-url-shortener/pkg/cli"
	library "github.com/georgeyord/go-url-shortener/pkg/helloworld"
)

const nameLabel = "name"

var helloworldCmd = &cobra.Command{
	Use: "helloworld",

	Short: "Modifies input string 'name' to a 'Hello World!' pattern",

	Example: `helloworld -n foo --> Hello Foo!
	helloworld -n foo bar --> Hello Foo Bar!`,

	Version: viper.GetString("boot.timestamp"),

	Run: helloworld,
}

func init() {
	rootCmd.AddCommand(helloworldCmd)
	helloworldCmd.Flags().StringP(nameLabel, "n", "", "Set your name")
}

func helloworld(cmd *cobra.Command, args []string) {
	name, _ := cmd.Flags().GetString(nameLabel)

	if name == "" {
		name = cli.GetInput(nameLabel)
	}

	message := library.GetHelloWorldMessage(name)
	cli.PrintMessage(message, aurora.Magenta)
}
