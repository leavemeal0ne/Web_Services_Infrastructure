package handler

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"lab4/internal/models"
	"net/http"
	"strconv"
)

var (
	ErrInvalidData       = errors.New("Invalid data input")
	ErrWhileDecodingData = errors.New("Failed to decode input data")
)

func (h *Handler) getAllClients() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := h.service.GetAllClients()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		h.respond(w, http.StatusOK, data)
	}
}

func (h *Handler) addClient() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &models.Client{}
		err := json.NewDecoder(r.Body).Decode(data)
		if err != nil {
			http.Error(w, ErrWhileDecodingData.Error(), http.StatusInternalServerError)
			return
		}
		err = h.service.CreateClient(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		h.respond(w, http.StatusOK, data)
	}
}

func (h *Handler) updateClient() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data := &models.Client{}
		data.Id = id
		err = json.NewDecoder(r.Body).Decode(data)
		if err != nil {
			http.Error(w, ErrWhileDecodingData.Error(), http.StatusInternalServerError)
			return
		}
		err = h.service.UpdateClient(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		h.respond(w, http.StatusOK, data)
	}
}

func (h *Handler) getClientById() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data := &models.Client{}
		data.Id = id
		err = h.service.GetClientById(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		h.respond(w, http.StatusOK, data)
	}
}

func (h *Handler) deleteClient() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data := &models.Client{}
		data.Id = id
		err = h.service.DeleteClient(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		h.respond(w, http.StatusOK, data)
	}
}
