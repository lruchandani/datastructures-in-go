package pubsubtopic

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
	"github.com/lruchandani/datastructures-in-go/offersender/offer"
	"github.com/lruchandani/datastructures-in-go/offersender/product"
)

//PubSubTopic Type
type PubSubTopic struct {
	client  pubsub.Client
	topic   pubsub.Topic
	context context.Context
}

//NewPubSubTopic - create new pubsub topic instance
func NewPubSubTopic(topicName string, project string) (*PubSubTopic, error) {
	pubsubTopic := new(PubSubTopic)
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, project)
	if err != nil {
		return nil, err
	}
	pubsubTopic.client = *client
	pubsubTopic.topic = *client.Topic(topicName)
	pubsubTopic.context = ctx
	return pubsubTopic, nil
}

func (pubsubClient *PubSubTopic) publish(message string) error {
	var results []*pubsub.PublishResult
	r := pubsubClient.topic.Publish(pubsubClient.context, &pubsub.Message{
		Data: []byte(message),
	})
	results = append(results, r)
	// Do other work ...
	for _, r := range results {
		id, err := r.Get(pubsubClient.context)
		if err != nil {
			fmt.Println("Error :", err)
			panic(err)
		}
		fmt.Printf("Published a message with a message ID: %s\n", id)
	}
	defer pubsubClient.topic.Stop()
	return nil
}

//PublishOffers - publish offers to pubsub
func (pubsubClient *PubSubTopic) PublishOffers(products []string, stores []string) {
	for _, p := range products {
		for _, s := range stores {
			json := offer.NewOffer(s, p).ToJSON()
			// fmt.Println(json)
			pubsubClient.publish(json)
		}
	}
}

//PublishProducts - Publish products to pubsub
func (pubsubClient *PubSubTopic) PublishProducts(products []string) {
	for _, p := range products {
		json := product.NewProduct(p).ToJSON()
		// fmt.Println(json)
		pubsubClient.publish(json)
	}
}
