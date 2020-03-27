package pubsub

import (
	"context"
	"fmt"
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
	source = ""
)

func (ps *PubSub) AsyncPull(function func()) error {
	topicName := os.Getenv("PUBSUB_TOPIC")
	subscriptionID := os.Getenv("PUBSUB_SUBSCRIPTION_ID")

	sub, err := ps.createSubscriptionIfNotExists(topicName, subscriptionID)
	if err != nil {
		log.Println("subscription error : ", err)
		return err
	}

	var mu sync.Mutex
	ctx := context.Background()
	if err := sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		mu.Lock()
		var errs []string
		defer func() {
			// log pub/sub message
			log.Println("[POS] [pos-inventory-stream] pubsub insert log")
			_, _ = ps.log.Create(ps.db, model.PubSubMessageLog{
				PubsubID:       msg.ID,
				Data:           string(msg.Data),
				PublishTime:    &msg.PublishTime,
				ReceiveTime:    &msg.PublishTime,
				CreatedAt:      time.Now(),
				Source:         source,
				Topic:          topicName,
				SubscriptionID: subscriptionID,
				ErrorProcess:   strings.Join(errs, "|"),
			})
			mu.Unlock()
		}()

		// TODO pull message
		//put logical pull stream
		f := function // Finish
		log.Println(f)
		msg.Ack()
		fmt.Println("Finish")
	}); err != nil {
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
