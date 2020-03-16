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

var listUrlPairCmd = &cobra.Command{
	Use: "list",

	Short: "list all url pairs",

	Version: viper.GetString("boot.timestamp"),

	Run: listUrlPairs,
}

func init() {
	rootCmd.AddCommand(listUrlPairCmd)
}

func listUrlPairs(cmd *cobra.Command, args []string) {
	viperDb := viper.Get("db")
	db := viperDb.(*gorm.DB)

	var urlPairs []models.UrlPair
	if err := db.Find(&urlPairs).Error; err != nil {
		cli.PrintMessage(err.Error(), aurora.Red)
		return
	}

	for _, urlPair := range urlPairs {
		cli.PrintMessage(fmt.Sprintf("%s\t%s", urlPair.Short, urlPair.Long), aurora.Cyan)
	}
}
