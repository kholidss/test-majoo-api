package main

import (
	"test-majoo-api/config"
	"test-majoo-api/internal/server"

	"github.com/labstack/echo/v4"

	"test-majoo-api/pkg/database"
	// "pkg/database"

	"test-majoo-api/pkg/firebase"
)

func main() {
	cfg := config.LoadConfig()

	db := database.DBInit(cfg)

	firebaseAuth := firebase.SetupFirebase()

	server := server.NewServer(cfg, db, firebaseAuth)

	server.Router.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("firebaseAuth", firebaseAuth)
			return next(c)
		}
	})

	server.ListenAndServe(cfg.ServerPort)
}
