package api

import (
	db "github.com/development/simple-bank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	server.router = router
	return server
}

// Start runs the HTTP server on the specified address.
func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

// errorResponse formats the error message as JSON.
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
