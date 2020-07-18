package pubsubtopic

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/lruchandani/datastructures-in-go/offersender/offer"
	"github.com/lruchandani/datastructures-in-go/offersender/product"
	"github.com/tidwall/limiter"
)

var ops uint64

//PubSubTopic Type
type PubSubTopic struct {
	client       pubsub.Client
	topic        pubsub.Topic
	context      context.Context
	noOfRoutines int
}

//NewPubSubTopic - create new pubsub topic instance
func NewPubSubTopic(topicName string, project string, concurrentRequests int) (*PubSubTopic, error) {
	pubsubTopic := new(PubSubTopic)
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, project)
	if err != nil {
		return nil, err
	}
	pubsubTopic.client = *client
	pubsubTopic.topic = *client.Topic(topicName)
	pubsubTopic.context = ctx
	pubsubTopic.noOfRoutines = concurrentRequests
	return pubsubTopic, nil
}

func (pubsubClient *PubSubTopic) publish(wg *sync.WaitGroup, limiter *limiter.Limiter, message string) error {
	var results []*pubsub.PublishResult
	defer done(wg)
	limiter.Begin()
	defer limiter.End()
	// fmt.Println("------------: ")
	r := pubsubClient.topic.Publish(pubsubClient.context, &pubsub.Message{
		Data: []byte(message),
	})
	results = append(results, r)
	// Do other work ...
	for _, r := range results {
		_, err := r.Get(pubsubClient.context)
		if err != nil {
			fmt.Println("Error :", err)
			panic(err)
		}
		// fmt.Printf("Published a message with a message ID: %s\n", id)
	}
	atomic.AddUint64(&ops, 1)
	return nil
}

func done(wg *sync.WaitGroup) {
	defer wg.Done()
}

func (pubsubClient *PubSubTopic) stop() {
	defer pubsubClient.topic.Stop()
}

//PublishOffers - publish offers to pubsub
func (pubsubClient *PubSubTopic) PublishOffers(products []string, stores []string) {
	t := 0
	var wg sync.WaitGroup
	timer := time.Now()
	limiter := limiter.New(pubsubClient.noOfRoutines)
	for _, p := range products {
		for _, s := range stores {
			json := offer.NewOffer(s, p).ToJSON()
			wg.Add(1)
			go pubsubClient.publish(&wg, limiter, json)
			t++
		}
	}
	wg.Wait()
	fmt.Printf("total time to publish %d records is %s seconds", t, time.Since(timer))
	pubsubClient.stop()
}

//PublishProducts - Publish products to pubsub
func (pubsubClient *PubSubTopic) PublishProducts(products []string) {
	t := 0
	var wg sync.WaitGroup
	timer := time.Now()
	limiter := limiter.New(pubsubClient.noOfRoutines)
	for _, p := range products {
		json := product.NewProduct(p).ToJSON()
		// time.Sleep(500 * time.Millisecond)
		wg.Add(1)
		// print(json)
		go pubsubClient.publish(&wg, limiter, json)
		t++
	}
	wg.Wait()
	fmt.Printf("total time to publish %d records is %s seconds", t, time.Since(timer))
	pubsubClient.stop()
}
func print(json string) {
	fmt.Println(json)
}
