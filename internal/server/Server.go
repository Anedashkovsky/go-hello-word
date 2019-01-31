package server

import (
	"fmt"
	"go-hello-word/internal/config"
	"log"
	"net/http"
)

// Server Struct for start server
type Server struct {
	config *config.Server
}

// Init Initializes server with config
func (server *Server) Init(config *config.Server) {
	server.config = config
}

// Start Start server with config fron Config
func (server *Server) Start() {
	address := server.config.GetHost() + ":" + server.config.GetPort()
	fmt.Println("Starting server at " + address)
	http.HandleFunc("/api/", server.createHandler())
	http.HandleFunc("/", server.createNotFoundHandler())

	log.Fatal(http.ListenAndServe(address, nil))
}

func (server *Server) createHandler() http.HandlerFunc {
	return func(responceWriter http.ResponseWriter, request *http.Request) {
		fmt.Println(request.URL)
		fmt.Fprint(responceWriter, "Hello there!")
	}
}

func (server *Server) createNotFoundHandler() http.HandlerFunc {
	return func(responceWriter http.ResponseWriter, request *http.Request) {
		fmt.Println(request.URL)
		http.NotFound(responceWriter, request)
	}
}
