package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	slogfiber "github.com/samber/slog-fiber"

	"github.com/usysrc/ausleihe/db"
	"github.com/usysrc/ausleihe/handler"
)

func main() {
	// create the logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// create the database
	database, err := db.CreateDatabase()
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer db.CloseDatabase(database)

	// Initialize standard Go html template engine
	engine := html.New("./views", ".html")
	engine.Reload(true)

	// Start fiber
	app := fiber.New(fiber.Config{
		Views:                 engine,
		DisableStartupMessage: true,
	})

	// Add structured logging middleware
	app.Use(slogfiber.New(logger))

	// Define routes
	app.Get("/", handler.IndexHandler)
	app.Post("/add-item", handler.ItemHandler)

	// Start server
	port := 3000
	logger.Info(fmt.Sprintf("Starting server on localhost:%d", port))
	err = app.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		logger.Error(err.Error())
	}
}
