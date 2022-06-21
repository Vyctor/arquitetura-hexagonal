package application_test

import (
	"testing"

	"github.com/codeedu/go-hexagonal/application"
	mock_application "github.com/codeedu/go-hexagonal/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestProductService_Get(test *testing.T) {
	controller := gomock.NewController(test)
	defer controller.Finish()

	product := mock_application.NewMockProductInterface(controller)
	persistence := mock_application.NewMockProductPersistenceInterface(controller)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil)

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Get("abc")
	require.Nil(test, err)
	require.Equal(test, product, result)
}

func TestProductService_Create(test *testing.T) {
	controller := gomock.NewController(test)
	defer controller.Finish()

	product := mock_application.NewMockProductInterface(controller)
	persistence := mock_application.NewMockProductPersistenceInterface(controller)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil)

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Create("Product 1", 100.0)
	require.Nil(test, err)
	require.Equal(test, product, result)
}

func TestProductService_EnableDisable(test *testing.T) {
	controller := gomock.NewController(test)
	defer controller.Finish()
	product := mock_application.NewMockProductInterface(controller)

	product.EXPECT().Enable().Return(nil)
	product.EXPECT().Disable().Return(nil)

	persistence := mock_application.NewMockProductPersistenceInterface(controller)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Enable(product)
	require.Nil(test, err)
	require.Equal(test, product, result)

	result, err = service.Disable(product)
	require.Nil(test, err)
	require.Equal(test, product, result)

}
