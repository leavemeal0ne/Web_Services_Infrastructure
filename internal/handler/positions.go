package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"lab4/internal/models"
	"net/http"
	"strconv"
)

func (h *Handler) getAllPositions() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := h.service.GetAllPositions()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		h.respond(w, http.StatusOK, data)
	}
}

func (h *Handler) addPosition() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &models.Position{}
		err := json.NewDecoder(r.Body).Decode(data)
		if err != nil {
			http.Error(w, ErrWhileDecodingData.Error(), http.StatusInternalServerError)
			return
		}
		err = h.service.CreatePosition(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		h.respond(w, http.StatusOK, data)
	}
}

func (h *Handler) updatePosition() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data := &models.Position{}
		data.Id = id
		err = json.NewDecoder(r.Body).Decode(data)
		if err != nil {
			http.Error(w, ErrWhileDecodingData.Error(), http.StatusInternalServerError)
			return
		}
		err = h.service.UpdatePosition(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		h.respond(w, http.StatusOK, data)
	}
}

func (h *Handler) getPositionById() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data := &models.Position{}
		data.Id = id
		err = h.service.GetPositionById(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		h.respond(w, http.StatusOK, data)
	}
}

func (h *Handler) deletePosition() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data := &models.Position{}
		data.Id = id
		err = h.service.DeletePosition(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		h.respond(w, http.StatusOK, data)
	}
}
