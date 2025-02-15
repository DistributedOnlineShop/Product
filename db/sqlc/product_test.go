package db

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"Product/util"
)

func CreateRandomProduct(t *testing.T, vendorID uuid.UUID) Product {
	data := CreateProductParams{
		VendorID:    vendorID,
		Name:        gofakeit.ProductName(),
		Description: gofakeit.ProductDescription(),
		Price:       util.GenerateRandomNumeric(),
		Discount:    util.GenerateRandomNumeric(),
		Stock:       util.GenerateRandomInt32(),
		Status:      util.GenerateRandomProductStatus(),
	}
	product, err := testStore.CreateProduct(context.Background(), data)
	require.NoError(t, err)
	require.NotEmpty(t, product)
	require.Equal(t, product.VendorID, vendorID)
	require.Equal(t, product.Name, data.Name)
	require.Equal(t, product.Description, data.Description)
	require.Equal(t, product.Price, data.Price)
	require.Equal(t, product.Discount, data.Discount)
	require.Equal(t, product.Stock, data.Stock)
	require.Equal(t, product.Status, data.Status)
	require.NotEmpty(t, product.ProductID)
	require.NotEmpty(t, product.CreatedAt)

	return product
}

func TestCreateProduct(t *testing.T) {
	vendor := CreateRandomVendors(t)
	CreateRandomProduct(t, vendor.VendorID)
}

func TestDeletProduct(t *testing.T) {
	vendor := CreateRandomVendors(t)
	product := CreateRandomProduct(t, vendor.VendorID)

	err := testStore.DeletProduct(context.Background(), product.ProductID)
	require.NoError(t, err)
}

func TestGetProductByProductId(t *testing.T) {
	vendor := CreateRandomVendors(t)
	product := CreateRandomProduct(t, vendor.VendorID)

	p, err := testStore.GetProductByProductId(context.Background(), product.ProductID)
	require.NoError(t, err)
	require.NotEmpty(t, p)
	require.Equal(t, p.ProductID, product.ProductID)
	require.Equal(t, p.VendorID, product.VendorID)
	require.Equal(t, p.Name, product.Name)
	require.Equal(t, p.Description, product.Description)
	require.Equal(t, p.Price, product.Price)
	require.Equal(t, p.Discount, product.Discount)
	require.Equal(t, p.Stock, product.Stock)
	require.Equal(t, p.Status, product.Status)
	require.NotEmpty(t, p.CreatedAt)
	require.NotEmpty(t, p.UpdatedAt)
}

func TestGetProductByVendorID(t *testing.T) {
	vendor := CreateRandomVendors(t)
	for i := 0; i < 10; i++ {
		CreateRandomProduct(t, vendor.VendorID)
	}
	products, err := testStore.GetProductByVendorID(context.Background(), vendor.VendorID)
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(products), 10)
}

func TestUpdateProduct(t *testing.T) {
	vendor := CreateRandomVendors(t)
	product := CreateRandomProduct(t, vendor.VendorID)

	newData := UpdateProductParams{
		ProductID:   product.ProductID,
		Name:        gofakeit.ProductName(),
		Description: gofakeit.ProductDescription(),
		Price:       util.GenerateRandomNumeric(),
		Discount:    util.GenerateRandomNumeric(),
		Stock:       util.GenerateRandomInt32(),
		Status:      util.GenerateRandomProductStatus(),
	}

	newP, err := testStore.UpdateProduct(context.Background(), newData)
	require.NoError(t, err)
	require.NotEmpty(t, newP)
	require.Equal(t, newP.ProductID, product.ProductID)
	require.Equal(t, newP.VendorID, product.VendorID)
	require.NotEqual(t, newP.Name, product.Name)
	require.NotEqual(t, newP.Description, product.Description)
	require.NotEqual(t, newP.Price, product.Price)
	require.NotEqual(t, newP.Discount, product.Discount)
	require.NotEqual(t, newP.Stock, product.Stock)
	require.NotEqual(t, newP.Status, product.Status)
	require.NotEmpty(t, newP.UpdatedAt)
}
