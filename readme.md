# QiStream

simple package stream in qasir tech

## Instalation

```
# go get github.com/qasir-id/qistream
# Go modules
$> go mod tidy
```

## Basic Usage
 
Env
```
convert file credential.json to base64 --> https://www.base64decode.org/
GCP_CREDENTIALS=someCodeBas64
GCP_PROJECT_ID=ProjectId
PUBSUB_TOPIC=TopicName
PUBSUB_SUBSCRIPTION_ID=SubcriptionName
```

Sample
```
import (
	"log"
	"github.com/jinzhu/gorm"
	"github.com/qasir-id/qistream/service/pubsub"
    psg "cloud.google.com/go/pubsub"
)
    //for async
	log.Println("starting Pub/Sub Client ")
	// RUN service PubSub
	var db *gorm.DB
		psService := pubsub.NewPubSubService(conn, pubsub.NewClient())
    	if err := psService.AsyncPull(func(ctx context.Context, msg *psg.Message) {
    		var mu sync.Mutex
    		func() {
    			mu.Lock()
    			var errs []string
    			log.Println(string(msg.Data))
    			defer psService.SaveLog(msg, errs, mu)
    			//put your code in here
    
                switch 1 {
                case 1 :
                    err := salesInventory.Handle(ctx, req)
                      //save error to logs pubsub
                      if err != nil {
                      errs = append(errs, err.Error())
                      }
                  break:    
                }
            
                //acknowlege message
    			msg.Ack()
    			fmt.Println("Finish")
    		}()
    	}); err != nil {
    		log.Fatalf("failed to pull pub/sub message : %v", err)
    	}

    //for publish
    ctx := context.Background()
	var pubMessage []byte
	pubsub.PublishTopic(ctx, pubMessage, "TOPIC_NAME")
    
    //you can also see in example folder
```