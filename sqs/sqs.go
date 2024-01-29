package sqs

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"log"
	"sync"
)

var sqsSvc *sqs.SQS
var queueUrl string
var svcOnce sync.Once

func Init() {
	// Initialize a session
	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String("us-east-1"),
		Credentials:      credentials.NewStaticCredentials("test", "test", ""),
		S3ForcePathStyle: aws.Bool(true),
		Endpoint:         aws.String("http://localhost:4566"),
	})

	if err != nil {
		log.Fatal("Failed to initialize sqs session!")
	}

	svcOnce.Do(func() {
		sqsSvc = sqs.New(sess)
		CreateQueue("testQueue")
	})

	result, err := sqsSvc.ListQueues(nil)
	for i, url := range result.QueueUrls {
		fmt.Printf("%d: %s\n", i, *url)
	}
}

func CreateQueue(name string) {
	result, err := sqsSvc.CreateQueue(&sqs.CreateQueueInput{
		QueueName: aws.String(name),
	})

	if err != nil {
		log.Fatal("Queue not created")
	}

	queueUrl = aws.StringValue(result.QueueUrl)
}

func SendMessageToSqs(msg []byte) (err error) {
	_, err = sqsSvc.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(string(msg)),
		QueueUrl:    &queueUrl,
	})

	if err != nil {
		log.Println("Could not send message")
	}
	return
}

func RecieveMessageFromSqs() (msgs []interface{}, msgErr error) {
	for {
		msgResult, err := sqsSvc.ReceiveMessage(&sqs.ReceiveMessageInput{
			QueueUrl:            &queueUrl,
			MaxNumberOfMessages: aws.Int64(10),
		})

		if err != nil {
			return msgs, err
		}

		if len(msgResult.Messages) == 0 {
			return msgs, err
		}

		for _, val := range msgResult.Messages {
			var body interface{}
			err := json.Unmarshal([]byte(*val.Body), &body)
			if err != nil {
				return msgs, err
			}

			msgs = append(msgs, body)

			_, err = sqsSvc.DeleteMessage(&sqs.DeleteMessageInput{
				QueueUrl:      aws.String(queueUrl),
				ReceiptHandle: val.ReceiptHandle,
			})

			if err != nil {
				return msgs, err
			}
		}
	}
}
