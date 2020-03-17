package commands

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"

	"github.com/logrusorgru/aurora"

	"github.com/georgeyord/go-url-shortener/pkg/cli"
	"github.com/georgeyord/go-url-shortener/pkg/models"
)

var createUrlPairCmd = &cobra.Command{
	Use: "create",

	Short: "creates a new short url from the long one given",

	Version: viper.GetString("boot.timestamp"),

	Run: createUrlPair,
}

func init() {
	rootCmd.AddCommand(createUrlPairCmd)
	createUrlPairCmd.Flags().StringP(longUrlLabel, "l", "", "Long url (required)")
	createUrlPairCmd.Flags().StringP(shortUrlLabel, "s", "", "Short url (optional)")
}

func createUrlPair(cmd *cobra.Command, args []string) {
	viperDb := viper.Get("db")
	db := viperDb.(*gorm.DB)

	long, _ := cmd.Flags().GetString(longUrlLabel)
	short, _ := cmd.Flags().GetString(shortUrlLabel)

	if long == "" {
		long = cli.GetInput(longUrlLabel)
	}

	// Create UrlPair
	urlPair := models.UrlPair{Short: short, Long: long}
	if err := db.Create(&urlPair).Error; err != nil {
		cli.PrintMessage(err.Error(), aurora.Red)
		return
	}

	cli.PrintMessage(fmt.Sprintf("New short url created: %s", urlPair.Short), aurora.Green)
	cli.PrintMessage(fmt.Sprintf("Long url: %s", urlPair.Long), aurora.BrightGreen)
}
