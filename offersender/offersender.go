package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/lruchandani/datastructures-in-go/offersender/offer"
	"github.com/lruchandani/datastructures-in-go/offersender/product"
	"github.com/lruchandani/datastructures-in-go/offersender/pubsubtopic"
)

func main() {
	option := flag.String("publishTo", "products", "valid values : products / offers")
	project := flag.String("project-id", "playground-es", "projet id where pubsub is created")
	var pubsubClient *pubsubtopic.PubSubTopic
	flag.Parse()
	if *option == "products" {
		pubsubClient, _ = pubsubtopic.NewPubSubTopic("catalog-products-local", *project)
	} else {
		pubsubClient, _ = pubsubtopic.NewPubSubTopic("catalog-offers-local", *project)
	}
	rand.Seed(time.Now().UTC().UnixNano())
	if *option == "offers" {
		products, _ := product.LoadProducts()
		stores, _ := offer.LoadStores()
		pubsubClient.PublishOffers(products, stores)
		return
	}
	products, _ := product.LoadProducts()
	fmt.Println(products, len(products))
	pubsubClient.PublishProducts(products)
}
