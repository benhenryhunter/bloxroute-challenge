package worker

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

//
// Subscribe subscribes to a queue
//
func Subscribe(queueName string, out chan *sqs.Message) error {
	mySession := session.Must(session.NewSession(&aws.Config{Region: aws.String("us-east-1")}))
	sqsClient := sqs.New(mySession)
	input := sqs.ReceiveMessageInput{
		QueueUrl:        aws.String(queueName),
		WaitTimeSeconds: aws.Int64(20),
	}
	for {
		output, err := sqsClient.ReceiveMessage(&input)
		if err != nil {
			return err
		}
		for _, message := range output.Messages {
			out <- message
		}
	}
}

//
// DeleteMessage deletes a message from a queue
//
func DeleteMessage(queueName, handle string) error {
	mySession := session.Must(session.NewSession(&aws.Config{Region: aws.String("us-east-1")}))
	sqsClient := sqs.New(mySession)
	_, err := sqsClient.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      aws.String(queueName),
		ReceiptHandle: aws.String(handle),
	})
	return err
}
