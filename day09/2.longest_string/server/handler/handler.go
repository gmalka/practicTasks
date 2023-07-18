package handler

import (
	"io"
	"net/http"
)

type StringHandler interface {
	ServeString(string) string
}

func Newhandler(sh StringHandler) Handler {
	return Handler{sh: sh}
}

type Handler struct {
	sh StringHandler
}

func (h Handler) InitRouter() *http.ServeMux {
	handler := http.NewServeMux()

	handler.HandleFunc("/api/substring", h.HandleString)

	return handler
}

func (h Handler) HandleString(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}
	defer r.Body.Close()

	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := h.sh.ServeString(string(b))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}
