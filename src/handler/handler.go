package handler

import (
	"net/http"

	repohandler "github.com/pavr1/people/src/repoHandler"
)

type HttpHandler struct {
	repoHandler *repohandler.RepoHandler
}

func NewHttpHandler(repo *repohandler.RepoHandler) *HttpHandler {
	return &HttpHandler{
		repoHandler: repo,
	}
}

func (h *HttpHandler) GetPeople(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, world!"))
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
