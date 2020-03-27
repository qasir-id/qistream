package pubsub

import (
	"context"
	"encoding/base64"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

func NewClient() *pubsub.Client {
	// open config
	b, err := base64.StdEncoding.WithPadding(base64.NoPadding).DecodeString(os.Getenv("GCP_CREDENTIALS"))
	if err != nil {
		log.Fatalf("failed load credential config : %v", err)
	}
	// new client
	ctx := context.Background()
	psClient, err := pubsub.NewClient(ctx, os.Getenv("GCP_PROJECT_ID"), option.WithCredentialsJSON(b))
	if err != nil {
		log.Fatalf("failed to connect Pubsub Client : %v", err)
	}
	log.Println("service Pub/Sub started !!!")
	return psClient
}
