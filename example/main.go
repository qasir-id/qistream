package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	psg "cloud.google.com/go/pubsub"
	"github.com/qasir-id/qistream/example/database"
	"github.com/qasir-id/qistream/service/pubsub"
)

func main() {
	os.Setenv("GCP_CREDENTIALS", "xxx")

	os.Setenv("GCP_PROJECT_ID", "qasir-pos-dev")
	os.Setenv("PUBSUB_TOPIC", "category")
	os.Setenv("PUBSUB_SUBSCRIPTION_ID", "category_sub")

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
	pubsub.PublishTopic(ctx, pubMessage, "category")
	//run sync
	psService := pubsub.NewPubSubService(conn, pubsub.NewClient())
	if err := psService.AsyncPull(func(ctx context.Context, msg *psg.Message) {
		var mu sync.Mutex
		func() {
			mu.Lock()
			var errs []string
			log.Println(string(msg.Data))
			defer psService.SaveLog(msg, errs, mu)
			//put your code in here
			msg.Ack()
			fmt.Println("Finish")
		}()
	}); err != nil {
		log.Fatalf("failed to pull pub/sub message : %v", err)
	}
}
