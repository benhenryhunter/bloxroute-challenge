package worker

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/dickmanben/bloxroute-challenge/server/types"
)

//
// Begins the worker and will automatically restart it on a panic
//
func InitWorker() {
	defer func() {
		if rec := recover(); rec != nil {
			go restartWorker()
		}
	}()
	out := make(chan *sqs.Message)
	go listen(out)
	err := Subscribe(os.Getenv("QUEUE_NAME"), out)
	if err != nil {
		fmt.Println(err)
	}
}

func restartWorker() {
	fmt.Println("server crashed but is restarting")
	go InitWorker()
}

func listen(out chan *sqs.Message) {
	for {
		select {
		case message := <-out:
			clientRequest := types.ClientRequest{}
			if err := json.Unmarshal([]byte(*message.Body), &clientRequest); err != nil {
				DeleteMessage(os.Getenv("QUEUE_NAME"), *message.ReceiptHandle)
				fmt.Println("Improper message received")
				continue
			}
			f, ok := handlers[clientRequest.Action]
			if !ok {
				fmt.Println("No handler for specified action")
				continue
			}
			f(clientRequest)
			DeleteMessage(os.Getenv("QUEUE_NAME"), *message.ReceiptHandle)
		}
	}
}
