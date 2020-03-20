package config

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/common-nighthawk/go-figure"
	kafka "github.com/georgeyord/go-url-shortener/pkg/kafka"
	"github.com/georgeyord/go-url-shortener/pkg/models"
	kafkalib "github.com/segmentio/kafka-go"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/spf13/viper"
)

const STAGING = "staging"
const PRODUCTION = "production"

func Init(role string) {
	initApplicationEnv(role)
	initConfig()
	initLogger()
}

func IsEnv(env string) bool {
	return viper.GetString("env") == env
}

func initConfig() {
	loadConfigFile("config")

	_, isEnvSet := os.LookupEnv("IS_DOCKER")
	if isEnvSet {
		loadConfigFile("config.docker")
	}

	loadConfigFile(fmt.Sprintf("config.%s", viper.GetString("env")))
	loadConfigFile(fmt.Sprintf("config.%s", viper.GetString("role")))
	loadConfigFile(fmt.Sprintf("config.%s.%s", viper.GetString("role"), viper.GetString("env")))

	viper.SetEnvPrefix("SHORTENER")
	viper.AutomaticEnv()

	if err := viper.WriteConfigAs("../log/config.yaml"); err != nil {
		log.Warn().Err(err).Msg("Writing config backup failed")
	}
}

func InitDb() *gorm.DB {
	dbType := viper.GetString("db.type")
	dbPath := viper.GetString("db.path")

	db, err := gorm.Open(dbType, dbPath)

	if err != nil {
		log.Fatal().Str("type", dbType).Str("path", dbPath).Err(err).Msg("Database failed to connect")
	}

	log.Info().Msg("Database loaded")
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
		log.Info().Str("topic", topic).Msg("Closing kafka writer")
		writer.Close()
	}
}

func PrintIntro() {
	appFigure := figure.NewFigure(viper.GetString("application.name"), viper.GetString("application.asciiart.theme"), true)
	appFigure.Print()

	if viper.IsSet("role") {
		roleFigure := figure.NewFigure(fmt.Sprintf("Role: %s", viper.GetString("role")), viper.GetString("application.asciiart.subtheme"), true)
		roleFigure.Print()
	}
}

func initApplicationEnv(role string) {
	now := time.Now()
	viper.SetDefault("env", STAGING)
	viper.SetDefault("role", role)
	viper.SetDefault("boot.timestamp", now.Unix())
	log.Debug().Str("Boot timestamp", viper.GetString("boot.timestamp")).Msg("")

	potentialEnv, isEnvSet := os.LookupEnv("APPLICATION_ENV")
	if isEnvSet {
		if potentialEnv == STAGING || potentialEnv == PRODUCTION {
			viper.Set("env", potentialEnv)
			log.Info().Str("env", potentialEnv).Msg("Application environment")
		} else {
			log.Fatal().Str("env", potentialEnv).Msg("Unsupported environment")
		}
	} else {
		log.Info().Str("env", viper.GetString("env")).Msg("Falling back to default environment")
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
		log.Debug().Str("config", configName).Msg("Loading config file (merge)")
		err = viper.MergeInConfig()
	} else {
		loadConfigFileInitialed = true
		log.Debug().Str("config", configName).Msg("Loading config file")
		err = viper.ReadInConfig()
	}

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Warn().Str("path", configPath).Str("config", configName).Str("type", configType).Msg("Config file not found")
		} else {
			// Config file was found but another error was produced
			log.Fatal().Str("config", configName).Err(err).Msg("Fatal error config file")
		}
	}
}
