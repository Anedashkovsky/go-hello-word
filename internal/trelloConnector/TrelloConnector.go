package trelloconnector

import (
	"go-hello-word/internal/config"
	"go-hello-word/internal/startupChecker"
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
	return trelloConnector.loadBoardData(trelloConnector.boardID)
}

func (trelloConnector *TrelloConnector) loadBoardData(boardID string) Board {
	return Board{
		ID:              "123",
		Description:     "Description",
		DescriptionData: "DescriptionData",
		Name:            "Name",
	}
}
