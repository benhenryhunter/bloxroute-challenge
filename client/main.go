package main

import (
	"os"

	"github.com/dickmanben/bloxroute-challenge/client/utils"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()

	args := os.Args[1:]
	switch args[0] {
	case "AddItem":
		key := args[1]
		utils.AddItem(key)
	case "GetItem":
		key := args[1]
		utils.GetItem(key)
	case "GetItems":
		utils.GetItems()
	case "DeleteItem":
		key := args[1]
		utils.DeleteItem(key)
	}

}
