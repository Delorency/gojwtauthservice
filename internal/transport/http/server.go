package http

import (
	cont "auth/internal/container"
	"fmt"
	"log"
	"net/http"
)

func NewHTTPServer(addr, port string, cont *cont.Container, logger *log.Logger) *http.Server {
	router := NewRouter(cont, logger)

	server := http.Server{
		Addr:    fmt.Sprintf("%s:%s", addr, port),
		Handler: router,
	}

	return &server
}
