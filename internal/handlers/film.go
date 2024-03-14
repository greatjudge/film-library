package handlers

import (
	"filmlibr/internal/entity"
	"filmlibr/internal/sending"
	"fmt"
	"net/http"
	"strconv"
)

type filmService interface {
	GetByID(id int) (entity.Film, error)
}

type FilmHandler struct {
	service filmService
}

func (h *FilmHandler) List(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.listPost(w, r)
	case http.MethodGet:
		h.listGet(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *FilmHandler) Film(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
	case http.MethodPatch:
		h.patch(w, r)
	case http.MethodPut:
		h.put(w, r)
	case http.MethodDelete:
		h.delete(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *FilmHandler) FilmActor(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.addActor(w, r)
	case http.MethodDelete:
		h.removeActor(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *FilmHandler) get(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		sending.JSONError(
			w,
			fmt.Errorf(badIDMes, idStr),
			http.StatusBadRequest,
		)
		return
	}

	a, err := h.service.GetByID(id)
	if err != nil {
		handleServiceError(w, err)
		return
	}
	sending.JSONMarshalAndSend(w, a, http.StatusOK)
}

func (h *FilmHandler) put(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		sending.JSONError(
			w,
			fmt.Errorf(badIDMes, idStr),
			http.StatusBadRequest,
		)
		return
	}
}

func (h *FilmHandler) patch(w http.ResponseWriter, r *http.Request) {

}

func (h *FilmHandler) delete(w http.ResponseWriter, r *http.Request) {

}
