package main

import (
	"hermes/internal/api/v1/handlers"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	port := os.Getenv("HERMES_PORT")
	if port == "" {
		port = "3200"
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Register GitHub handler
	githubHandler := handlers.NewGitHubHandler()
	e.GET("/github/repos", githubHandler.GetRepositories)

	// Health check
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	log.Printf("Starting server on port %s...\n", port)
	if err := e.Start(":" + port); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
