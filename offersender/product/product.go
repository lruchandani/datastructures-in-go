package product

import (
	"bytes"
	"time"

	"github.com/lruchandani/datastructures-in-go/offersender/common"
)

var productTemplate = common.LoadTemplate("product.template")

//Product Type
type Product struct {
	liam      string
	updatedAt string
}

// NewProduct - create new Product Instance
func NewProduct(productID string) *Product {
	m := new(Product)
	m.liam = productID
	m.updatedAt = time.Now().Format("2006-01-02T15:04:05")
	return m
}

//ToJSON -convert offer to offer Json
func (product *Product) ToJSON() string {
	buf := &bytes.Buffer{}
	productTemplate.Execute(buf, map[string]interface{}{"liam": product.liam})
	return buf.String()
}

//LoadProducts - load products
func LoadProducts() ([]string, error) {
	products, err := common.LoadFile("products.txt")
	return products, err
}
