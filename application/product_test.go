package application_test

import (
	"testing"

	"github.com/codeedu/go-hexagonal/application"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(test *testing.T) {
	product := application.Product{}

	product.Name = "Product 1"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()

	require.Nil(test, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(test, "price must be greater than zero to enable the product", err.Error())

	product.Price = 100
	product.Status = application.ENABLED
	err = product.Enable()
	require.Equal(test, "product is already enabled", err.Error())
}
