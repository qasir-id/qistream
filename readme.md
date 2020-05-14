# QiStream

simple package stream in qasir tech

## Instalation

```
# Go modules
$> go mod tidy or go get github.com/qasir-id/qistream
```

## Basic Usage
 
Env
```
convert file credential.json to base64 --> https://www.base64decode.org/
GCP_CREDENTIALS=someCodeBas64
GCP_PROJECT_ID=ProjectId
PUBSUB_TOPIC=TopicName
PUBSUB_SUBSCRIPTION_ID=SubcriptionName

PUBSUB_TOPIC_CATEGORY=TopicName
```

### Sample

## Pull Message

```go
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

        pubsub.Source = "source service" // set source log service 

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
                    err := createCategory.Handle(ctx, req)
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
```

## Publish Message

- struct recommendation
```
type PubSubMessage struct {
	Action string       // action message ex: category:create or category:delete etc
	Data   interface{}  // data for processing action
}

```
- example

```
    msg := PubSubMessage{
    		Action: "create:category",
    		Data:   {"id":0,"name":"Umum", ...},
    	}

    ctx := context.Background()
	byteFirebasePublish, err := json.Marshal(msg)
    		if err != nil {
    			return nil, err
    		}
    
	pubsub.PublishTopic(ctx, byteFirebasePublish, "PUBSUB_TOPIC_CATEGORY")
```

*you can also see in example folder*
