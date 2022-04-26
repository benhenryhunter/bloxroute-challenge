package types

type ClientRequest struct {
	Action string `json:"action"`
	Key    string `json:"key"`
	Item   Item   `json:"item"`
}
