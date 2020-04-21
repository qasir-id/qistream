package pubsub

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
)

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
