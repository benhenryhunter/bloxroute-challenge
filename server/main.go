package main

import (
	"github.com/dickmanben/bloxroute-challenge/server/worker"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()

	worker.InitWorker()
}
