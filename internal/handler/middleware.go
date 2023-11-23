package handler

import (
    "fmt"
    "log"
    "net/http"
)

func (h *Handler) logIdentify(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        data := fmt.Sprintf("New request\tIP:%s\tmethod:%s\tURI:%s", r.RemoteAddr, r.Method, r.URL.Path)

        log.Println(data)

        next.ServeHTTP(w, r.WithContext(r.Context()))
    })

}
