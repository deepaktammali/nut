package handlers

import (
	"fmt"
	"net/http"
	"nut/internal/config"
	"nut/internal/stores"
	"nut/internal/stores/postgres"
)

type Handler struct {
	http.Handler
	app   *config.AppConfig
	store *stores.Store
}

func NewHandler(app *config.AppConfig) *Handler {
	handler := new(Handler)
	handler.app = app

	pgStore := postgres.NewStore(app.Db)
	handler.store = pgStore

	var router = http.NewServeMux()
	router.Handle("GET /health", http.HandlerFunc(handler.healthHandler))
	router.Handle("POST /tickets", http.HandlerFunc(handler.createTicket))

	handler.Handler = router
	return handler
}

func (h *Handler) healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}
