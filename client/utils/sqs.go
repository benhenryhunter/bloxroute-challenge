package utils

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func SendMessage(body string) error {
	mySession := session.Must(session.NewSession(&aws.Config{Region: aws.String("us-east-1")}))
	sqsClient := sqs.New(mySession)
	if _, err := sqsClient.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(body),
		QueueUrl:    aws.String(os.Getenv("QUEUE_NAME")),
	}); err != nil {
		return err
	}
	return nil
}
