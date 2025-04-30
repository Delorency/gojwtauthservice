package main

import (
	"auth/internal/app"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	app.Start()
}
