package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"

	library "github.com/georgeyord/go-scrumpoker-api/pkg/helloworld"
	"github.com/logrusorgru/aurora"
)

var cfgFile string

const nameLabel = "name"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "helloworld",

	Short: "Modifies input string 'name' to a 'Hello World!' pattern",

	Example: `helloworld -n foo --> Hello Foo!
helloworld -n foo bar --> Hello Foo Bar!`,

	Version: viper.GetString("boot.timestamp"),

	Run: run,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Here you will define your flags and configuration settings.
func init() {
	// Persistent flags, which will be global for your application.
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./.cmd-config.yaml)")

	// Local flags, which will only run when this action is called directly.
	rootCmd.Flags().StringP(nameLabel, "n", "", "Set your name")
}

func run(cmd *cobra.Command, args []string) {
	name, _ := cmd.Flags().GetString(nameLabel)

	if name == "" {
		name = GetInput(nameLabel)
	}

	message := library.GetHelloWorldMessage(name)
	PrintMessage(message, aurora.Magenta)
}
