package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) respond(w http.ResponseWriter, sCode int, data interface{}) {
	if data != nil {
		_ = json.NewEncoder(w).Encode(data)
	}
}
