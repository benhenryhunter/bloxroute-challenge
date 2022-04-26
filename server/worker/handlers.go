package worker

import (
	"fmt"

	"github.com/dickmanben/bloxroute-challenge/server/types"
)

var Items = types.Items{}

var handlers = map[string]func(res types.ClientRequest){
	"AddItem":    AddItem,
	"GetItem":    GetItem,
	"GetItems":   GetItems,
	"DeleteItem": DeleteItem,
}

func AddItem(req types.ClientRequest) {
	if _, err := Items.GetItem(req.Item.Key); err == nil {
		fmt.Printf("item already exists with key %v\n", req.Item.Key)
		return
	}
	fmt.Printf("adding item with key %v\n", req.Item.Key)
	Items.AddItem(req.Item)
}

func GetItem(req types.ClientRequest) {
	fmt.Printf("getting item with key %v\n", req.Key)
	item, err := Items.GetItem(req.Key)
	if err != nil {
		fmt.Printf("item not found with key %v", req.Key)
		return
	}
	fmt.Println(item)
}

func GetItems(req types.ClientRequest) {
	fmt.Println("getting all items\n")
	fmt.Println(Items.GetItems())
}

func DeleteItem(req types.ClientRequest) {
	fmt.Printf("deleting item with key %v\n", req.Key)
	err := Items.DeleteItem(req.Key)
	if err != nil {
		fmt.Printf("item not found with key %v", req.Key)
		return
	}
	fmt.Printf("deleted item with key %v\n", req.Key)
}
