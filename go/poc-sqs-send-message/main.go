package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

var REGION = "ap-northeast-1"

func sqsClient(ctx context.Context) *sqs.Client {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(REGION))
	if err != nil {
		log.Fatalln(err)
	}
	client := sqs.NewFromConfig(cfg)
	return client
}

func main() {
	ctx := context.Background()
	client := sqsClient(ctx)
	params := sqs.SendMessageInput{
		MessageBody:       aws.String("test message"),
		QueueUrl:          aws.String("queue url"),
		MessageAttributes: map[string]types.MessageAttributeValue{},
	}
	out, err := client.SendMessage(ctx, &params)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("messageId", *out.MessageId)
}
