package cli_test

import (
	"fmt"
	"testing"

	"github.com/codeedu/go-hexagonal/adapters/cli"
	mock_application "github.com/codeedu/go-hexagonal/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRun(test *testing.T) {
	controller := gomock.NewController(test)

	defer controller.Finish()

	productName := "Product Test"
	productPrice := 30.0
	productStatus := "enabled"
	productId := "abc"

	productMock := mock_application.NewMockProductInterface(controller)

	productMock.EXPECT().GetId().Return(productId).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(controller)
	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(productMock).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(productMock).Return(productMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Product ID %s with the name %s has been create with the price %f and status %s", productId, productName, productPrice, productStatus)
	result, err := cli.Run(service, "create", "", productName, productPrice)
	require.Nil(test, err)
	require.Equal(test, resultExpected, result)

	resultExpected = fmt.Sprintf("Product %s has been enabled.", productName)
	result, err = cli.Run(service, "enable", productId, productName, productPrice)
	require.Nil(test, err)
	require.Equal(test, resultExpected, result)

	resultExpected = fmt.Sprintf("Product %s has been disabled.", productName)
	result, err = cli.Run(service, "disable", productId, productName, productPrice)
	require.Nil(test, err)
	require.Equal(test, resultExpected, result)

	resultExpected = fmt.Sprintf("Product ID: %s \nName: %s \nPrice: %f \nStatus: %s\n", productId, productName, productPrice, productStatus)
	result, err = cli.Run(service, "", productId, productName, productPrice)
	require.Nil(test, err)
	require.Equal(test, resultExpected, result)
}
