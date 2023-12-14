package app_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/marielifm/hexagonal-architecture/app"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := app.Product{}

	product.Name = "Hello"
	product.Status = app.DISABLED
	product.Price = 10

	err := product.Enable()

	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := app.Product{}

	product.Name = "Hello"
	product.Status = app.ENABLED
	product.Price = 0

	err := product.Disable()

	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "the price must be zero to disable the product", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := app.Product{}

	product.ID = uuid.New().String()
	product.Name = "Hello"
	product.Status = app.DISABLED
	product.Price = 10

	valid, err := product.IsValid()

	require.Nil(t, err)
	require.Equal(t, true, valid)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = app.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "the price must be greater or equal to zero", err.Error())
}
