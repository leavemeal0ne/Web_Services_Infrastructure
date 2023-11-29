package handler

import (
	"github.com/gorilla/mux"
	"lab4/internal/service"
	"net/http"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router.Use(h.logIdentify, h.JsonContent, h.corsMiddleware)

	clients := router.PathPrefix("/clients").Subrouter()

	{
		clients.HandleFunc("", h.getAllClients()).Methods(http.MethodGet, http.MethodOptions)
		clients.HandleFunc("", h.addClient()).Methods(http.MethodPost, http.MethodOptions)
		clients.HandleFunc("/{id:[0-9]+}", h.updateClient()).Methods(http.MethodPut, http.MethodOptions)
		clients.HandleFunc("/{id:[0-9]+}", h.getClientById()).Methods(http.MethodGet, http.MethodOptions)
		clients.HandleFunc("/{id:[0-9]+}", h.deleteClient()).Methods(http.MethodDelete, http.MethodOptions)
	}
	workers := router.PathPrefix("/workers").Subrouter()

	{
		workers.HandleFunc("", h.getAllWorkers()).Methods(http.MethodGet, http.MethodOptions)
		workers.HandleFunc("", h.addWorker()).Methods(http.MethodPost, http.MethodOptions)
		workers.HandleFunc("/{id:[0-9]+}", h.updateWorker()).Methods(http.MethodPut, http.MethodOptions)
		workers.HandleFunc("/{id:[0-9]+}", h.getWorkerById()).Methods(http.MethodGet, http.MethodOptions)
		workers.HandleFunc("/{id:[0-9]+}", h.deleteWorker()).Methods(http.MethodDelete, http.MethodOptions)
	}

	positions := router.PathPrefix("/positions").Subrouter()

	{
		positions.HandleFunc("", h.getAllPositions()).Methods(http.MethodGet, http.MethodOptions)
		positions.HandleFunc("", h.addPosition()).Methods(http.MethodPost, http.MethodOptions)
		positions.HandleFunc("/{id:[0-9]+}", h.updatePosition()).Methods(http.MethodPut, http.MethodOptions)
		positions.HandleFunc("/{id:[0-9]+}", h.getPositionById()).Methods(http.MethodGet, http.MethodOptions)
		positions.HandleFunc("/{id:[0-9]+}", h.deletePosition()).Methods(http.MethodDelete, http.MethodOptions)
	}

	return router
}
