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
	option := flag.String("publish-to", "products", "valid values : products / offers")
	project := flag.String("project-id", "playground-es", "projet id where pubsub is created")
	batchNumber := flag.Int("batch-number", 0, "batch number ")
	batchSize := flag.Int("batch-size", 1, "no of product records to pick from processing by this container")
	noOfRoutines := flag.Int("concurrent-requests", 5000, "no of product records to pick from processing by this container")
	sourceFilePath := flag.String("source-file-path", ".", "source file path containing  products.txt and stores.txt")
	pickFrom := *batchNumber * *batchSize
	var pubsubClient *pubsubtopic.PubSubTopic
	flag.Parse()
	if *option == "products" {
		pubsubClient, _ = pubsubtopic.NewPubSubTopic("catalog-products-local", *project, *noOfRoutines)
	} else {
		pubsubClient, _ = pubsubtopic.NewPubSubTopic("catalog-offers-local", *project, *noOfRoutines)
	}
	products, _ := product.LoadProducts(*sourceFilePath)
	if len(products) < pickFrom {
		err := fmt.Errorf("Total of records is %d and instructed to pick from %d", len(products), pickFrom)
		panic(err)
	}
	pickTo := pickFrom + *batchSize
	if pickTo > len(products) {
		pickTo = len(products)
	}
	products = products[pickFrom:pickTo]
	rand.Seed(time.Now().UTC().UnixNano())
	if *option == "offers" {
		stores, _ := offer.LoadStores(*sourceFilePath)
		pubsubClient.PublishOffers(products, stores)
		return
	}
	fmt.Println("Processing ", len(products), " products")
	pubsubClient.PublishProducts(products)
}
