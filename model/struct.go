package model

import "time"

type PubSubMessageLog struct {
	ID             int64      `json:"id" gorm:"id"`
	PubsubID       string     `json:"pubsub_id" gorm:"pubsub_id"`
	Data           string     `json:"data" gorm:"data"`
	Attribute      string     `json:"attribute" gorm:"attribute"`
	PublishTime    *time.Time `json:"publish_time" gorm:"publish_time"`
	ReceiveTime    *time.Time `json:"receive_time" gorm:"receive_time"`
	CreatedAt      time.Time  `json:"created_at" gorm:"created_at"`
	Source         string     `json:"source" gorm:"source"`
	Topic          string     `json:"topic" gorm:"topic"`
	SubscriptionID string     `json:"subscription_id" gorm:"subscription_id"`
	ErrorProcess   string     `json:"error_process" gorm:"error_process"`
}

func (m PubSubMessageLog) TableName() string {
	return "pubsub_message_log"
}
