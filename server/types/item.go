package types

import (
	"errors"
	"sync"
)

type Item struct {
	Key  string `json:"key"`
	Data []byte `json:"data"`
}

type Items struct {
	sync.RWMutex
	items []Item
}

func (i *Items) AddItem(item Item) {
	i.Lock()
	defer i.Unlock()
	i.items = append(i.items, item)
}

func (i *Items) GetItems() []Item {
	return i.items
}

func (i *Items) GetItem(key string) (*Item, error) {
	i.RLock()
	defer i.RUnlock()
	for _, item := range i.items {
		if item.Key == key {
			return &item, nil
		}
	}
	return nil, errors.New("key not found")
}

func (i *Items) DeleteItem(key string) error {
	i.Lock()
	defer i.Unlock()
	for index, item := range i.items {
		if item.Key == key {
			i.items = append(i.items[:index], i.items[index+1:]...)
			return nil
		}
	}
	return nil
}
