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
	router.Use(h.logIdentify, h.JsonContent)

	clients := router.PathPrefix("/clients").Subrouter()

	{
		clients.HandleFunc("", h.getAllClients()).Methods(http.MethodGet)
		clients.HandleFunc("", h.addClient()).Methods(http.MethodPost)
		clients.HandleFunc("/{id:[0-9]+}", h.updateClient()).Methods(http.MethodPut)
		clients.HandleFunc("/{id:[0-9]+}", h.getClientById()).Methods(http.MethodGet)
		clients.HandleFunc("/{id:[0-9]+}", h.deleteClient()).Methods(http.MethodDelete)
	}
	workers := router.PathPrefix("/workers").Subrouter()

	{
		workers.HandleFunc("", h.getAllWorkers()).Methods(http.MethodGet)
		workers.HandleFunc("", h.addWorker()).Methods(http.MethodPost)
		workers.HandleFunc("/{id:[0-9]+}", h.updateWorker()).Methods(http.MethodPut)
		workers.HandleFunc("/{id:[0-9]+}", h.getWorkerById()).Methods(http.MethodGet)
		workers.HandleFunc("/{id:[0-9]+}", h.deleteWorker()).Methods(http.MethodDelete)
	}

	positions := router.PathPrefix("/positions").Subrouter()

	{
		positions.HandleFunc("", h.getAllPositions()).Methods(http.MethodGet)
		positions.HandleFunc("", h.addPosition()).Methods(http.MethodPost)
		positions.HandleFunc("/{id:[0-9]+}", h.updatePosition()).Methods(http.MethodPut)
		positions.HandleFunc("/{id:[0-9]+}", h.getPositionById()).Methods(http.MethodGet)
		positions.HandleFunc("/{id:[0-9]+}", h.deletePosition()).Methods(http.MethodDelete)
	}

	return router
}
