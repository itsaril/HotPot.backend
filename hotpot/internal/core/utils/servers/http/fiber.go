// Package http provides an HTTP server implementation using the Fiber framework.
// It allows for easy initialization, route setup, and server management.
package http

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

// Server represents an HTTP server using the Fiber framework.
type Server struct {
	Port      string               // The port on which the server will run.
	App       *fiber.App           // The Fiber app instance.
	InitRoute func(app *fiber.App) // A function to initialize routes for the Fiber app.
}

// NewFiber creates a new HTTP server instance using the Fiber framework.
//
// Arguments:
//
//	port - The port on which the server will run.
//	app - The Fiber app instance to be used.
//	initRoute - A function for initializing routes in the Fiber app (can be nil).
//
// Returns:
//
//	A pointer to a new Server instance.
func NewFiber(port string, app *fiber.App, initRoute func(app *fiber.App)) *Server {
	if initRoute != nil {
		initRoute(app) // Initialize routes if a function is provided.
	}
	return &Server{
		Port:      port,
		App:       app,
		InitRoute: initRoute,
	}
}

// Start starts the HTTP server on the specified port.
//
// Returns:
//
//	An error if the server fails to start.
func (s *Server) Start() error {
	log.Print("\033[32m") // Print green text for server start message.
	log.Printf("HTTP server running on port %s", s.Port)
	log.Print("\033[0m") // Reset terminal color.
	return s.App.Listen(":" + s.Port)
}

// Stop stops the HTTP server.
//
// Note:
//
//	The Fiber framework does not provide a direct Stop() function.
//	Use App.Shutdown() for graceful shutdown if needed.
//
// Returns:
//
//	Always returns nil as there is no actual stop logic implemented here.
func (s *Server) Stop() error {
	log.Println("HTTP server stopping")
	// Placeholder for stopping logic. Implement app.Shutdown() if necessary.
	return nil
}
