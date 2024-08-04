package handler

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/pavr1/people/models"
	repohandler "github.com/pavr1/people/repoHandler"
)

type HttpHandler struct {
	repoHandler *repohandler.RepoHandler
}

func NewHttpHandler(repo *repohandler.RepoHandler) *HttpHandler {
	return &HttpHandler{
		repoHandler: repo,
	}
}

func (h *HttpHandler) GetPersonList(w http.ResponseWriter, r *http.Request) {
	person := models.NewPerson(h.repoHandler.Config)

	list, err := person.GetPersonList()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

		return
	}

	bytes, err := json.Marshal(list)
	if err != nil {
		log.WithError(err).Fatal("Failed to marshal person list")

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func (h *HttpHandler) GetPerson(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, world!"))
}

func (h *HttpHandler) CreatePerson(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, world!"))
}

func (h *HttpHandler) UpdatePerson(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, world!"))
}

func (h *HttpHandler) DeletePerson(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, world!"))
}
