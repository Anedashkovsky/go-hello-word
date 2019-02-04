package trelloconnector

// Board - Board type which return fron trello api
type Board struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"desc"`
	DescriptionData string `json:"descData"`
}
