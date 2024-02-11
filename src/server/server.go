package server

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

var appLogger *log.Logger = log.New(os.Stdout, "Server: ", log.LstdFlags|log.Lmsgprefix)

type Server struct {
	*fiber.App
}

func New() *Server {
	server := &Server{
		fiber.New(fiber.Config{
			// Server configuration
		}),
	}
	// server.Use(swagger.Config{
	// 	BasePath: "/",
	// 	FilePath: "./docs/openapi.json",
	// 	Path:     "openapi",
	// 	Title:    "Bid2Bless API Docs",
	// })
	return server
}

func (s *Server) Start(address string) error {
	if err := s.Listen(address); err != nil {
		appLogger.Println("Server listen failed!")
		return err
	}
	return nil
}
