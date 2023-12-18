package app_test

import (
	"testing"

	"github.com/marielifm/hexagonal-architecture/app"
	mock_app "github.com/marielifm/hexagonal-architecture/app/mocks"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_app.NewMockProductInterface(ctrl)
	persistence := mock_app.NewMockProductPersistenceInterface(ctrl)

	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := app.ProductService{
		Persistence: persistence,
	}

	result, err := service.Get("abc")
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_app.NewMockProductInterface(ctrl)
	persistence := mock_app.NewMockProductPersistenceInterface(ctrl)

	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := app.ProductService{
		Persistence: persistence,
	}

	result, err := service.Create("abc", 10.0)
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_Enable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_app.NewMockProductInterface(ctrl)
	product.EXPECT().Enable().Return(nil)
	persistence := mock_app.NewMockProductPersistenceInterface(ctrl)

	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := app.ProductService{
		Persistence: persistence,
	}

	result, err := service.Enable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_Disable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_app.NewMockProductInterface(ctrl)
	product.EXPECT().Disable().Return(nil)
	persistence := mock_app.NewMockProductPersistenceInterface(ctrl)

	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := app.ProductService{
		Persistence: persistence,
	}

	result, err := service.Disable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)
}
