package trelloconnector

import (
	"encoding/json"
	"fmt"
	"go-hello-word/internal/config"
	"go-hello-word/internal/startupChecker"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// TrelloConnector Connector to trello public api. It can get information from board.
// Requres TRELLO_KEY TRELLO_TOKEN and TRELLO_BOARD env vars defined
type TrelloConnector struct {
	key        string
	token      string
	boardID    string
	envChecker *checker.EnvChecker
	config     *config.TrelloConfig
}

// NewTrelloConnector Create trello connector and init it with env vars
func NewTrelloConnector() *TrelloConnector {
	var key string
	var token string
	var boardID string
	var envChecker *checker.EnvChecker
	var trelloConfig *config.TrelloConfig

	envChecker = checker.NewEnvChecker()
	trelloConfig = config.NewTrelloConfig()
	key = envChecker.GetEnv("TRELLO_KEY")
	token = envChecker.GetEnv("TRELLO_TOKEN")
	boardID = envChecker.GetEnv("TRELLO_BOARD")

	return &TrelloConnector{key: key, token: token, envChecker: envChecker, boardID: boardID, config: trelloConfig}
}

// GetBoardData Get data from board with id dined in TRELLO_BOARD env variable
func (trelloConnector *TrelloConnector) GetBoardData() Board {
	channel := make(chan Board)
	var board Board

	go trelloConnector.loadBoardData(channel, trelloConnector.boardID)
	board = <-channel

	return board
}

func (trelloConnector *TrelloConnector) loadBoardData(channel chan Board, boardID string) {
	fmt.Println(trelloConnector.generateBoardRequest())
	response, err := http.Get(trelloConnector.generateBoardRequest())

	if err != nil {
		log.Fatalln(err)
	}

	data, _ := ioutil.ReadAll(response.Body)

	var boardData Board
	json.Unmarshal(data, &boardData)
	channel <- boardData

	defer response.Body.Close()
}

func (trelloConnector *TrelloConnector) generateBoardRequest() string {
	const DELIMITER = "?"
	query := url.Values{}
	query.Set("token", trelloConnector.token)
	query.Set("key", trelloConnector.key)

	var stringBuilder strings.Builder

	stringBuilder.WriteString(trelloConnector.config.GetBaseUrl())
	stringBuilder.WriteString(trelloConnector.config.GetBoardApiUrl())
	stringBuilder.WriteString("/")
	stringBuilder.WriteString(trelloConnector.boardID)
	stringBuilder.WriteString(DELIMITER)
	stringBuilder.WriteString(query.Encode())

	return stringBuilder.String()
}
