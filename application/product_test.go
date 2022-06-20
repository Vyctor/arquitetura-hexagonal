package application_test

import (
	"testing"

	"github.com/codeedu/go-hexagonal/application"
	uuid "github.com/satori/go.uuid"
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

func TestProduct_Disable(test *testing.T) {
	product := application.Product{}
	product.Name = "Product 1"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(test, err)

	product.Price = 1000
	err = product.Disable()
	require.Equal(test, "price must be zero in order to have the product disabled", err.Error())

	product.Price = 0
	err = product.Disable()
	require.Equal(test, "product is already disabled", err.Error())
}

func TestProduct_IsValid(test *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Product 1"
	product.Status = application.DISABLED
	product.Price = 10

	_, err := product.IsValid()

	require.Nil(test, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(test, "status is required", err.Error())

	product.Status = application.ENABLED
	product.ID = ""
	_, err = product.IsValid()
	require.Equal(test, "id is required", err.Error())

	product.ID = uuid.NewV4().String()

	product.Name = ""
	_, err = product.IsValid()
	require.Equal(test, "name is required", err.Error())

	product.Name = "Product 1"
	product.Price = -10
	_, err = product.IsValid()
	require.Equal(test, "price must be greater than zero", err.Error())

	product.Price = 0
	_, err = product.IsValid()
	require.Equal(test, "price is required", err.Error())
}
