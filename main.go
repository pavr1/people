package main

import (
	"fmt"
	"net/http"

	"github.com/pavr1/people/config"
	_http "github.com/pavr1/people/handlers/http"
	"github.com/pavr1/people/handlers/repo"
	log "github.com/sirupsen/logrus"
)

func main() {
	log := setupLogger()
	config, err := config.NewConfig()
	if err != nil {
		return
	}

	repoHandler, err := repo.NewRepoHandler(log, config)
	if err != nil {
		log.WithError(err).Error("Failed to create repo handler")

		return
	}

	httpHandler := _http.NewHttpHandler(repoHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Info("*")
	})

	http.HandleFunc("/person", httpHandler.GetPersonList)
	http.HandleFunc("/person/create", httpHandler.CreatePerson)
	http.HandleFunc("/person/update", httpHandler.UpdatePerson)
	http.HandleFunc("/person/delete", httpHandler.DeletePerson)
	http.HandleFunc("/person/:id", httpHandler.GetPerson)

	log.WithField("port", config.Server.Port).Info("Starting Server...")
	// Start the HTTP server
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Server.Port), nil))
}

func setupLogger() *log.Entry {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	return log.NewEntry(log.StandardLogger())
}
