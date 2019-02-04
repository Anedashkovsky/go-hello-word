package server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-hello-word/internal/config"
	"go-hello-word/internal/trelloConnector"
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

	router := mux.NewRouter()
	router.HandleFunc("/api/hello", server.createHandler()).Methods("GET")
	router.HandleFunc("/api/board", server.createGetBoardHandler()).Methods("GET")

	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(address, nil))
}

type response struct {
	Data string `json:"data"`
}

func (server *Server) createHandler() http.HandlerFunc {
	return func(responceWriter http.ResponseWriter, request *http.Request) {
		responceWriter.Header().Set("Content-type", "application/json")
		responceWriter.WriteHeader(http.StatusOK)

		fmt.Println(request.URL)

		responceData := response{Data: "Hello there!"}
		json.NewEncoder(responceWriter).Encode(responceData)
	}
}

func (server *Server) createGetBoardHandler() http.HandlerFunc {
	return func(responceWriter http.ResponseWriter, request *http.Request) {
		var connector *trelloconnector.TrelloConnector
		connector = trelloconnector.NewTrelloConnector()

		responceWriter.Header().Set("Content-type", "application/json")
		responceWriter.WriteHeader(http.StatusOK)

		fmt.Println(request.URL)

		responceData := connector.GetBoardData()
		json.NewEncoder(responceWriter).Encode(responceData)
	}
}

func (server *Server) createNotFoundHandler() http.HandlerFunc {
	return func(responceWriter http.ResponseWriter, request *http.Request) {
		fmt.Println(request.URL)
		http.NotFound(responceWriter, request)
	}
}
