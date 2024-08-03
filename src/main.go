package main

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
}

func main() {
	setupLogger()
	config, err := readConfig()
	if err != nil {
		log.WithField("error", err).Error("Failed to read config")
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Info("*")
		fmt.Fprintf(w, "*")
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	})

	log.WithField("port", config.Server.Port).Info("Starting Server...")
	// Start the HTTP server
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Server.Port), nil))
}

func readConfig() (*Config, error) {
	// Set the configuration file name and type
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	// Set the configuration file path
	viper.AddConfigPath(".")

	// Read the configuration file
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	// Unmarshal the configuration into a struct
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func setupLogger() {
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}
