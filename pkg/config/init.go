package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/common-nighthawk/go-figure"
	kafka "github.com/georgeyord/go-url-shortener/pkg/kafka"
	"github.com/georgeyord/go-url-shortener/pkg/models"
	kafkalib "github.com/segmentio/kafka-go"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/spf13/viper"
)

func Init() {
	initApplicationEnv()
	initConfig()
}

func initConfig() {
	loadConfigFile("config")

	_, isEnvSet := os.LookupEnv("IS_DOCKER")
	if isEnvSet {
		loadConfigFile("config.docker")
	}

	loadConfigFile("config." + viper.GetString("env"))

	viper.SetEnvPrefix("SCRUMPOKER")
	viper.AutomaticEnv()

	if err := viper.WriteConfigAs("../log/config.yaml"); err != nil {
		log.Printf("Writing config backup failed: %s", err)
	}
}

func InitDb() *gorm.DB {
	dbType := viper.GetString("db.type")
	dbPath := viper.GetString("db.path")

	db, err := gorm.Open(dbType, dbPath)

	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to connect to '%s' database '%s'\nError: %s", dbType, dbPath, err.Error()))
	}

	log.Printf("DB of type %s loaded from: %s", dbType, dbPath)
	models.SetupModels(db)

	return db
}

func InitKafkaWriters() map[string]*kafkalib.Writer {
	statsTopic := viper.GetString("kafka.writers.stats.topic")
	statsWriter := kafka.NewWriter(statsTopic)

	return map[string]*kafkalib.Writer{
		statsTopic: statsWriter,
	}
}

func CloseKafkaWriters(writers map[string]*kafkalib.Writer) {
	for topic, writer := range writers {
		log.Printf("Closing kafka writer for topic '%s'", topic)
		writer.Close()
	}
}

func PrintIntro(role string) {
	appFigure := figure.NewFigure(viper.GetString("application.name"), viper.GetString("application.asciiart.theme"), true)
	appFigure.Print()

	if role != "" {
		roleFigure := figure.NewFigure(fmt.Sprintf("Role: %s", role), viper.GetString("application.asciiart.subtheme"), true)
		roleFigure.Print()
	}
}

func initApplicationEnv() {
	now := time.Now()
	viper.SetDefault("env", "staging")
	viper.SetDefault("boot.timestamp", now.Unix())
	log.Printf("Boot timestamp: %s", viper.GetString("boot.timestamp"))

	potentialEnv, isEnvSet := os.LookupEnv("APPLICATION_ENV")
	if isEnvSet {
		if potentialEnv == "staging" || potentialEnv == "production" {
			viper.Set("env", potentialEnv)
			log.Printf("Application environment: %s", viper.GetString("env"))
		} else {
			log.Fatalf("Unsupported environment: %s", potentialEnv)
		}
	} else {
		log.Printf("Falling back to default environment: %s", viper.GetString("env"))
	}
}

var loadConfigFileInitialed = false

func loadConfigFile(configName string) {
	const configPath = "../config"
	const configType = "yaml"
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	// Handle errors reading the config file
	var err error
	if loadConfigFileInitialed {
		log.Printf("Loading config file (merge): %s", configName)
		err = viper.MergeInConfig()
	} else {
		loadConfigFileInitialed = true
		log.Printf("Loading config file: %s", configName)
		err = viper.ReadInConfig()
	}

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Printf("Config file not found: %s/%s.%s", configPath, configName, configType)
		} else {
			// Config file was found but another error was produced
			log.Fatalf("Fatal error config file '%s': %s \n", configName, err)
		}
	}
}
