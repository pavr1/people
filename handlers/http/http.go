package http

import (
	"encoding/json"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"

	repohandler "github.com/pavr1/people/handlers/repo"
	"github.com/pavr1/people/models"
	"github.com/pavr1/people/models/request"
)

type HttpHandler struct {
	repo *repohandler.RepoHandler
}

func NewHttpHandler(repo *repohandler.RepoHandler) *HttpHandler {
	return &HttpHandler{
		repo: repo,
	}
}

func (h *HttpHandler) GetPersonList(w http.ResponseWriter, r *http.Request) {
	people, err := h.repo.GetPersonList()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

		return
	}

	bytes, err := json.Marshal(people)
	if err != nil {
		log.WithError(err).Fatal("Failed to marshal person list")

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func (h *HttpHandler) GetPerson(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.WithError(err).Fatal("Failed to read request body")

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	idParam := request.IDParam{}

	err = json.Unmarshal(body, &idParam)
	if err != nil {
		log.WithError(err).Fatal("Failed to unmarshal request body")

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	person, err := h.repo.GetPerson(idParam.ID)
	if err != nil {
		//will need to check for not found
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

		return
	}

	bytes, err := json.Marshal(person)
	if err != nil {
		log.WithError(err).Fatal("Failed to marshal person")

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func (h *HttpHandler) CreatePerson(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.WithError(err).Fatal("Failed to read request body")

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	person := models.Person{}

	err = json.Unmarshal(body, &person)
	if err != nil {
		log.WithError(err).Fatal("Failed to unmarshal request body")

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.repo.CreatePerson(&person)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

		return
	}

	w.WriteHeader(http.StatusOK)
	//might need to change this to return the created person
	w.Write([]byte("Person successfully created"))
}

func (h *HttpHandler) UpdatePerson(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.WithError(err).Fatal("Failed to read request body")

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	person := models.Person{}

	err = json.Unmarshal(body, &person)
	if err != nil {
		log.WithError(err).Fatal("Failed to unmarshal request body")

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.repo.UpdatePerson(&person)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Person successfully updated"))
}

func (h *HttpHandler) DeletePerson(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.WithError(err).Fatal("Failed to read request body")

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	idParam := request.IDParam{}

	err = json.Unmarshal(body, &idParam)
	if err != nil {
		log.WithError(err).Fatal("Failed to unmarshal request body")

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.repo.DeletePerson(idParam.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Person successfully deleted"))
}
