package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type API struct {
	r    *chi.Mux
	port string
}

type APIOptions struct {
	Port string
}

func (a *API) Mount() {
	a.r.Use(middleware.RequestID)
	a.r.Use(middleware.RealIP)
	a.r.Use(middleware.Logger)
	a.r.Use(middleware.Recoverer)

	a.r.Route("/v1", func(r chi.Router) {
		r.Get("/heartbeat", heartbeat)
	})
}

func (a *API) Server() *http.Server {
	return &http.Server{
		Addr:    fmt.Sprintf(":%s", a.port),
		Handler: a.r,
	}
}

func (a *API) Start() error {
	log.Println(fmt.Sprintf("starting server on port %s", a.port))
	return a.Server().ListenAndServe()
}

func NewAPI(r *chi.Mux, opts APIOptions) *API {
	return &API{
		r:    r,
		port: opts.Port,
	}
}
