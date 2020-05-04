package pubsub

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/qasir-id/qistream/model"

	"github.com/qasir-id/qistream/repo"

	"cloud.google.com/go/pubsub"
	"github.com/jinzhu/gorm"
)

// Service struct
type PubSub struct {
	c   *pubsub.Client
	db  *gorm.DB
	log repo.PubSubMessageLogRepo
}

func NewPubSubService(db *gorm.DB, client *pubsub.Client) *PubSub {
	var ps PubSub
	ps.c = client
	ps.db = db
	ps.log = repo.NewPubSubMessageLogGorm()
	return &ps
}

var (
	Source = ""
)

func (ps *PubSub) SaveLog(msg *pubsub.Message, errs []string, mu sync.Mutex) {
	// log pub/sub message
	log.Println("[POS] [pos-inventory-stream] pubsub insert log")
	bytesAttribute, _ := json.Marshal(msg.Attributes)
	_, _ = ps.log.Create(ps.db, model.PubSubMessageLog{
		PubsubID:       msg.ID,
		Data:           string(msg.Data),
		Attribute:      string(bytesAttribute),
		PublishTime:    &msg.PublishTime,
		ReceiveTime:    &msg.PublishTime,
		CreatedAt:      time.Now(),
		Source:         Source,
		Topic:          os.Getenv("PUBSUB_TOPIC"),
		SubscriptionID: os.Getenv("PUBSUB_SUBSCRIPTION_ID"),
		ErrorProcess:   strings.Join(errs, "|"),
	})
	mu.Unlock()
}

func (ps *PubSub) AsyncPull(function func(ctx context.Context, msg *pubsub.Message)) error {
	topicName := os.Getenv("PUBSUB_TOPIC")
	subscriptionID := os.Getenv("PUBSUB_SUBSCRIPTION_ID")

	sub, err := ps.createSubscriptionIfNotExists(topicName, subscriptionID)
	if err != nil {
		log.Println("subscription error : ", err)
		return err
	}

	ctx := context.Background()
	if err := sub.Receive(ctx, function); err != nil {
		log.Println("[POS] [pos-inventory-stream] receive error : %s", err.Error())
		return err
	}
	return nil
}

func (ps *PubSub) createSubscriptionIfNotExists(topicName, subscription string) (*pubsub.Subscription, error) {
	ctx := context.Background()

	s := ps.c.Subscription(subscription)
	ok, err := s.Exists(ctx)
	if err != nil {
		return s, err
	}
	if ok {
		return s, nil
	}

	s, err = ps.c.CreateSubscription(ctx, subscription, pubsub.SubscriptionConfig{
		Topic:       ps.c.Topic(topicName),
		AckDeadline: 30 * time.Second,
	})
	if err != nil {
		return s, err
	}
	return s, nil
}
