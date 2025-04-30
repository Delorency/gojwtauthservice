package http

import (
	"auth/internal/container"
	"log"

	_ "auth/docs"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"

	ah "auth/internal/transport/http/handlers"
)

func AddMiddleware(router *chi.Mux) *chi.Mux {
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	return router
}

func NewRouter(cont *container.Container, logger *log.Logger) *chi.Mux {
	router := AddMiddleware(chi.NewRouter())
	handlers := ah.NewAuthHandler(cont.AuthService, logger)

	router.Get("/swagger/*", httpSwagger.WrapHandler)

	router.Post("/access/{guid}", handlers.Access)

	router.Post("/refresh", handlers.Refresh)

	return router
}
