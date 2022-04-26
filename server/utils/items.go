package utils

import (
	"github.com/dickmanben/bloxroute-challenge/server/types"
)

var Items = types.Items{}

func GetItem(key string) (*types.Item, error) {
	return Items.FindByKey(key)
}

func GetItems() types.Items {
	return Items
}

func AddItem(item types.Item) {
	// mutex.Lock()
	Items = append(Items, item)
	// mutex.UnLock()
}

func DeleteItem(key string) error {
	return Items.DeleteByKey(key)
}
