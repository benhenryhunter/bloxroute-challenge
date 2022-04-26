package types

import "errors"

type Item struct {
	Key  string `json:"key"`
	Data []byte `json:"data"`
}

type Items []Item

func (i Items) FindByKey(key string) (*Item, error) {
	for _, item := range i {
		if item.Key == key {
			return &item, nil
		}
	}
	return nil, errors.New("key not found")
}

func (i *Items) DeleteByKey(key string) error {
	for index, item := range *i {
		if item.Key == key {
			*i = append((*i)[:index], (*i)[index+1:]...)
			return nil
		}
	}
	return nil
}
