package handlers

import (
	"fmt"
	"net/http"
	"nut/internal/config"
)

type Handler struct {
	http.Handler
	app *config.AppConfig
}

func NewHandler(app *config.AppConfig) *Handler {
	handler := new(Handler)
	handler.app = app

	var router = http.NewServeMux()
	router.Handle("/ping", http.HandlerFunc(handler.pingHandler))

	handler.Handler = router
	return handler
}

func (h *Handler) pingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Pong")
}
