package pubsub

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/pubsub"
)

func PublishTopic(ctx context.Context, pubMessage []byte, topicName string) error {
	client := NewClient()
	topic := client.Topic(topicName)

	defer topic.Stop()

	msg, _ := json.Marshal(pubMessage)
	pubRes := topic.Publish(ctx, &pubsub.Message{Data: msg})
	fmt.Println("Pubsub message : " + string(msg))
	if _, err := pubRes.Get(ctx); err != nil {
		return err
	}

	return nil
}
