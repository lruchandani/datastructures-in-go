package offer

import (
	"bytes"
	"math/rand"
	"strconv"
	"time"

	"github.com/lruchandani/datastructures-in-go/offersender/common"
)

var offerTemplate = common.LoadTemplate("offer.template")

//Offer type
type Offer struct {
	offerID   int
	storeID   int
	version   int
	updatedAt string
	liam      string
	wasPrice  int
	price     []int
}

//NewOffer - create new offer instance
func NewOffer(storeID string, productID string) *Offer {
	m := new(Offer)
	m.storeID, _ = strconv.Atoi(storeID)
	m.liam = productID
	m.offerID = 1000 + rand.Intn(8999)
	m.version = 100000 + rand.Intn(99999)
	m.updatedAt = time.Now().Format("2006-01-02T15:04:05")
	m.wasPrice = rand.Intn(10)
	m.price = []int{rand.Intn(10), rand.Intn(10)}
	return m
}

//ToJSON -convert offer to offer Json
func (offer *Offer) ToJSON() string {
	buf := &bytes.Buffer{}
	offerTemplate.Execute(buf, map[string]interface{}{
		"offerId":    offer.offerID,
		"liam":       offer.liam,
		"updated_at": offer.updatedAt,
		"storeId":    offer.storeID,
		"version":    offer.version,
		"price":      offer.price,
		"wasPrice":   offer.wasPrice})
	return buf.String()
}

//LoadStores - load stores
func LoadStores(folder string) ([]string, error) {
	stores, err := common.LoadFile(folder + "/" + "stores.txt")
	if err != nil {
		panic(err)
	}
	return stores, err
}
