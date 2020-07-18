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
	productTemplate.Execute(buf, map[string]interface{}{"liam": product.liam, "updated_at": product.updatedAt})
	return buf.String()
}

//LoadProducts - load products
func LoadProducts(folder string) ([]string, error) {
	file := folder + "/" + "products.txt"
	products, err := common.LoadFile(file)
	if err != nil {
		panic(err)
	}
	return products, err
}
