package handler

import (
    "fmt"
    "log"
    "net/http"
)

func (h *Handler) logIdentify(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        data := fmt.Sprintf("New request  IP:%s  method:%s  URI:%s", r.RemoteAddr, r.Method, r.URL.Path)

        log.Println(data)

        next.ServeHTTP(w, r.WithContext(r.Context()))
    })

}

func (h *Handler) JsonContent(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        next.ServeHTTP(w, r.WithContext(r.Context()))
    })

}

func (h *Handler) corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token, Authorization")
        w.WriteHeader(http.StatusOK)

        if r.Method == "OPTIONS" {
            return
        }

        next.ServeHTTP(w, r)
    })
}
