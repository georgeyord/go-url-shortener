package commands

import (
	"context"
	"fmt"
	"os"

	"github.com/georgeyord/go-url-shortener/pkg/cli"
	"github.com/georgeyord/go-url-shortener/pkg/models"
	"github.com/jinzhu/gorm"
	"github.com/logrusorgru/aurora"
	kafkalib "github.com/segmentio/kafka-go"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "get",
	Short:   "Given the short url, print the long url",
	Version: viper.GetString("boot.timestamp"),
	Run:     printUrlPair,
	Args:    cobra.ExactArgs(1),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func printUrlPair(cmd *cobra.Command, args []string) {
	viperDb := viper.Get("db")
	db := viperDb.(*gorm.DB)

	short := args[0]

	var urlPair models.UrlPair
	if err := db.Where("short = ?", short).First(&urlPair).Error; err != nil {
		cli.PrintMessage(err.Error(), aurora.Red)
		return
	}

	var writer *kafkalib.Writer
	statsTopic := viper.GetString("kafka.topics.stats")
	viperStatsWriter := viper.Get(statsTopic)
	if viperStatsWriter != nil {
		writer = viperStatsWriter.(*kafkalib.Writer)
		writer.WriteMessages(context.Background(), kafkalib.Message{
			Key:   []byte(urlPair.Short),
			Value: []byte(urlPair.Long),
		})
	}
	cli.PrintMessage(urlPair.Long, aurora.Cyan)
}
