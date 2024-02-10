package server

import (
	"bid2bless/src/database"
	"bid2bless/src/routes"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
)

type Server struct {
	app        *fiber.App
	mainLogger *log.Logger
	db         *database.DB
}

func NewServer(store *database.DB) (*Server, error) {
	server := &Server{
		app: fiber.New(fiber.Config{
			// Server configuration
		}),
		mainLogger: log.New(os.Stdout, "Main: ", log.LstdFlags|log.Lmsgprefix),
		db:         store,
	}

	routes.SetupRoutes(server.app)
	err := server.app.Listen(os.Getenv("PORT"))
	server.mainLogger.Fatal(err)
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.app.Listen(address)

}
