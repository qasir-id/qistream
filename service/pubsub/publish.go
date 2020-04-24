package pubsub

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/pubsub"
)

//publish topic with Message Data Only
func PublishTopic(ctx context.Context, pubMessage []byte, topicName string) error {
	client := NewClient()
	topic := client.Topic(topicName)

	defer topic.Stop()

	pubRes := topic.Publish(ctx, &pubsub.Message{Data: pubMessage})
	fmt.Println("Pubsub message : " + string(pubMessage))
	if _, err := pubRes.Get(ctx); err != nil {
		return err
	}

	return nil
}

//publish topic with custom message
func Publish(ctx context.Context,message pubsub.Message,topicName string) error{
	client := NewClient()
	topic := client.Topic(topicName)

	defer topic.Stop()

	pubRes := topic.Publish(ctx, &message)
	bytes, _ := json.Marshal(message)
	fmt.Println("Pubsub message : " + string(bytes))
	if _, err := pubRes.Get(ctx); err != nil {
		return err
	}

	return nil
}

