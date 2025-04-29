package handlers

import (
	"auth/internal/service"
	"log"
	"net/http"
)

type AuthHandlerI interface {
	Access(w http.ResponseWriter, r *http.Request)
	Refresh(w http.ResponseWriter, r *http.Request)
}

type authHandler struct {
	service service.AuthServiceI
	logger  *log.Logger
}

func NewAuthHandler(service service.AuthServiceI, logger *log.Logger) AuthHandlerI {
	return &authHandler{service, logger}
}
