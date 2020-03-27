package main

import (
	"context"
	"log"
	"os"

	"github.com/qasir-id/qistream/example/database"
	"github.com/qasir-id/qistream/service/pubsub"
)

func main() {
	os.Setenv("GCP_CREDENTIALS", "xxxx")

	os.Setenv("GCP_PROJECT_ID", "xxxx")
	os.Setenv("PUBSUB_TOPIC", "INVENTORY")
	os.Setenv("PUBSUB_SUBSCRIPTION_ID", "INVENTORY_SUB")

	os.Setenv("db_user", "root")
	os.Setenv("db_password", "root")
	os.Setenv("db_host", "127.0.0.1")
	os.Setenv("db_port", "3306")
	os.Setenv("db_database", "qasir_db")
	os.Setenv("driver", "mysql")

	log.Println("starting Pub/Sub Client ")
	// RUN service PubSub
	conn, _ := database.GetGormConnection()
	ctx := context.Background()
	//run publish
	var pubMessage []byte
	pubsub.PublishTopic(ctx, pubMessage, "INVENTORY")
	//run sync
	psService := pubsub.NewPubSubService(conn, pubsub.NewClient())
	if err := psService.AsyncPull(func() {
		//put code function in here
		log.Println("function")
	}); err != nil {
		log.Fatalf("failed to pull pub/sub message : %v", err)
	}
}
