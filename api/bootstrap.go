package main

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func bootstrap() {
	setupEnv()
	setupConfig()
}

func setupEnv() {
	viper.SetDefault("env", "staging")

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

func setupConfig() {
	loadConfigFile("config")

	_, isEnvSet := os.LookupEnv("IS_DOCKER")
	if isEnvSet {
		loadConfigFile("config.docker")
	}

	loadConfigFile("config." + viper.GetString("env"))

	viper.SetEnvPrefix("SCRUMPOKER")
	viper.AutomaticEnv()

	if err := viper.WriteConfigAs("log/config.yaml"); err != nil {
		log.Printf("Writing config backup failed: %s", err)
	}
}

var loadConfigFileInitialed = false

func loadConfigFile(configName string) {
	const configPath = "./config"
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
