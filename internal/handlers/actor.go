package handlers

import (
	"encoding/json"
	"filmlibr/internal/entity"
	"filmlibr/internal/sending"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

const BadIDMes = "id must be int, get %v"

type actorService interface {
	GetAll() ([]entity.ActorWithFilms, error)
	Add(a entity.Actor) (entity.Actor, error)
	GetByID(id int) (entity.Actor, error)
	UpdateCompletely(a entity.Actor) (entity.Actor, error)
	UpdatePartial(a entity.Actor) (entity.Actor, error)
	Delete(id int) error
}

type ActorHandler struct {
	service actorService
}

func (h *ActorHandler) List(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.ListGet(w, r)
	case http.MethodPost:
		h.ListPost(w, r)
	default:
		// TODO metod not allowed
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *ActorHandler) Actor(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.Get(w, r)
	case http.MethodPut:
		h.Put(w, r)
	case http.MethodPatch:
		h.Patch(w, r)
	case http.MethodDelete:
		h.Delete(w, r)
	default:
		// TODO metod not allowed
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *ActorHandler) ListGet(w http.ResponseWriter, r *http.Request) {
	acts, err := h.service.GetAll()
	if err != nil {
		handleServiceError(w, err)
		return
	}
	sending.JSONMarshalAndSend(w, acts, http.StatusOK)
}

func (h *ActorHandler) ListPost(w http.ResponseWriter, r *http.Request) {
	a := entity.Actor{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &a)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	a, err = h.service.Add(a)
	if err != nil {
		handleServiceError(w, err)
		return
	}
	sending.JSONMarshalAndSend(w, a, http.StatusCreated)
}

func (h *ActorHandler) Get(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		sending.JSONError(
			w,
			fmt.Errorf(BadIDMes, idStr),
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

func (h *ActorHandler) Put(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		sending.JSONError(
			w,
			fmt.Errorf(BadIDMes, idStr),
			http.StatusBadRequest,
		)
		return
	}

	a := entity.Actor{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &a)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	a.ID = id
	a, err = h.service.UpdateCompletely(a)
	if err != nil {
		handleServiceError(w, err)
		return
	}
	sending.JSONMarshalAndSend(w, a, http.StatusOK)
}

func (h *ActorHandler) Patch(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		sending.JSONError(
			w,
			fmt.Errorf(BadIDMes, idStr),
			http.StatusBadRequest,
		)
		return
	}

	a := entity.Actor{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &a)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	a.ID = id
	a, err = h.service.UpdatePartial(a)
	if err != nil {
		handleServiceError(w, err)
		return
	}
	sending.JSONMarshalAndSend(w, a, http.StatusOK)
}

func (h *ActorHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		sending.JSONError(
			w,
			fmt.Errorf(BadIDMes, idStr),
			http.StatusBadRequest,
		)
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		handleServiceError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func handleServiceError(w http.ResponseWriter, err error) {
	// TODO
}