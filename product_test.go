package shopware

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleProductService_Get() {
	app := New(token, endpoint)

	product, err := app.Products.Get(exampleProduct)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(product)
}

func TestSpacesServiceGet(t *testing.T) {
	var err error
	assert := assert.New(t)

	app := New(token, endpoint)

	app.Debug = true

	product, err := app.Products.Get(exampleProduct)
	assert.Nil(err)

	t.Log(product)
}
