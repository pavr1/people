package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pavr1/people/config"
	_http "github.com/pavr1/people/handlers/http"
	"github.com/pavr1/people/handlers/repo"
	log "github.com/sirupsen/logrus"
)

func main() {
	router := mux.NewRouter()

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

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	log.Info("*")
	// })

	router.HandleFunc("/person/list", httpHandler.GetPersonList).Methods("GET")
	router.HandleFunc("/person/create", httpHandler.CreatePerson).Methods("POST")
	router.HandleFunc("/person/update", httpHandler.UpdatePerson).Methods("PUT")
	router.HandleFunc("/person/delete", httpHandler.DeletePerson).Methods("DELETE")
	router.HandleFunc("/person/:id", httpHandler.GetPerson).Methods("GET")

	log.WithField("port", config.Server.Port).Info("Starting Server...")
	// Start the HTTP server
	log.Error(http.ListenAndServe(fmt.Sprintf(":%d", config.Server.Port), router))
}

func setupLogger() *log.Entry {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	log.SetReportCaller(true)

	return log.NewEntry(log.StandardLogger())
}
