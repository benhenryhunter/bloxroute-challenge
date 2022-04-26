package utils

import (
	"encoding/json"
	"log"

	"github.com/dickmanben/bloxroute-challenge/client/types"
)

func AddItem(key string) {
	clientReqest := types.ClientRequest{
		Action: "AddItem",
		Item: types.Item{
			Key: key,
		},
	}
	b, err := json.Marshal(&clientReqest)
	if err != nil {
		log.Fatal(err)
	}
	if err := SendMessage(string(b)); err != nil {
		log.Fatal(err)
	}
}

func GetItem(key string) {
	clientReqest := types.ClientRequest{
		Action: "GetItem",
		Key:    key,
	}
	b, err := json.Marshal(&clientReqest)
	if err != nil {
		log.Fatal(err)
	}
	if err := SendMessage(string(b)); err != nil {
		log.Fatal(err)
	}
}

func GetItems() {
	clientReqest := types.ClientRequest{
		Action: "GetItems",
	}
	b, err := json.Marshal(&clientReqest)
	if err != nil {
		log.Fatal(err)
	}
	if err := SendMessage(string(b)); err != nil {
		log.Fatal(err)
	}
}

func DeleteItem(key string) {
	clientReqest := types.ClientRequest{
		Action: "DeleteItem",
		Key:    key,
	}
	b, err := json.Marshal(&clientReqest)
	if err != nil {
		log.Fatal(err)
	}
	if err := SendMessage(string(b)); err != nil {
		log.Fatal(err)
	}
}
