package main

import (
	"fmt"
	"net/http"

	"github.com/pavr1/people/config"
	log "github.com/sirupsen/logrus"
)

func main() {
	setupLogger()
	config, err := config.NewConfig()
	if err != nil {
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

func setupLogger() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}
