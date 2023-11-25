package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"lab4/internal/models"
	"net/http"
	"strconv"
)

func (h *Handler) getAllWorkers() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := h.service.GetAllWorkers()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		h.respond(w, http.StatusOK, data)
	}
}

func (h *Handler) addWorker() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &models.Worker{}
		err := json.NewDecoder(r.Body).Decode(data)
		if err != nil {
			http.Error(w, ErrWhileDecodingData.Error(), http.StatusInternalServerError)
			return
		}
		err = h.service.CreateWorker(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		h.respond(w, http.StatusOK, data)
	}
}

func (h *Handler) updateWorker() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data := &models.Worker{}
		data.Id = id
		err = json.NewDecoder(r.Body).Decode(data)
		if err != nil {
			http.Error(w, ErrWhileDecodingData.Error(), http.StatusInternalServerError)
			return
		}
		err = h.service.UpdateWorker(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		h.respond(w, http.StatusOK, data)
	}
}

func (h *Handler) getWorkerById() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data := &models.Worker{}
		data.Id = id
		err = h.service.GetWorkerById(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		h.respond(w, http.StatusOK, data)
	}
}

func (h *Handler) deleteWorker() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data := &models.Worker{}
		data.Id = id
		err = h.service.DeleteWorker(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		h.respond(w, http.StatusOK, data)
	}
}
