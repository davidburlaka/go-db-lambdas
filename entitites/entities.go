package entitites

var TableNameBase = "user_"

type ReadEvent struct {
	EntityId string `json:"entity_id"`
	Username string `json:"username"`
}

type ReadResponse struct {
	Username string `json:"username"`
	Data     string `json:"data"`
}

type WriteEvent struct {
	EntityId string `json:"entity_id"`
	Username string `json:"username"`
	Data     string `json:"data"`
}
