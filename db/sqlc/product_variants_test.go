package db

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"

	"Product/util"
)

func CreateRandomProductVariants(t *testing.T, productID string) ProductVariant {
	var jsondata *gofakeit.JSONOptions
	json, err := gofakeit.JSON(jsondata)
	require.NoError(t, err)

	data := CreateProductVariantsParams{
		ProductID:  productID,
		Sku:        util.GenerateSKU(),
		Price:      util.GenerateNumeric(),
		Stock:      util.GenerateInt32(),
		Attributes: json,
		Status:     util.GeneratePVStatus(),
	}

	pv, err := testStore.CreateProductVariants(context.Background(), data)
	require.NoError(t, err)
	require.NotEmpty(t, pv)
	require.NotEmpty(t, pv.PvID)
	require.Equal(t, data.ProductID, pv.ProductID)
	require.Equal(t, data.Sku, pv.Sku)
	require.Equal(t, data.Price, pv.Price)
	require.Equal(t, data.Stock, pv.Stock)
	require.Equal(t, data.Attributes, pv.Attributes)
	require.Equal(t, data.Status, pv.Status)
	require.NotEmpty(t, pv.CreatedAt)

	return pv
}

func TestCreateProductVariants(t *testing.T) {
	vendor := CreateRandomVendors(t)
	product := CreateRandomProduct(t, vendor.VendorID)
	CreateRandomProductVariants(t, product.ProductID)
}

func TestDeleteProductVariantsByPvid(t *testing.T) {
	vendor := CreateRandomVendors(t)
	product := CreateRandomProduct(t, vendor.VendorID)
	pv := CreateRandomProductVariants(t, product.ProductID)

	err := testStore.DeleteProductVariantsByPvid(context.Background(), pv.PvID)
	require.NoError(t, err)
}

func TestGetProductVariantsByProductId(t *testing.T) {
	vendor := CreateRandomVendors(t)
	product := CreateRandomProduct(t, vendor.VendorID)
	for i := 0; i < 5; i++ {
		CreateRandomProductVariants(t, product.ProductID)
	}

	pvList, err := testStore.GetProductVariantsByProductId(context.Background(), product.ProductID)
	require.NoError(t, err)
	require.NotEmpty(t, pvList)
	require.GreaterOrEqual(t, len(pvList), 3)
}
func TestGetProductVariantsByPvId(t *testing.T) {
	vendor := CreateRandomVendors(t)
	product := CreateRandomProduct(t, vendor.VendorID)
	pv := CreateRandomProductVariants(t, product.ProductID)

	data, err := testStore.GetProductVariantsByPvid(context.Background(), pv.PvID)
	require.NoError(t, err)
	require.NotEmpty(t, data)
}

func TestGetProductVariantsByStatus(t *testing.T) {
	for i := 0; i < 30; i++ {
		vendor := CreateRandomVendors(t)
		product := CreateRandomProduct(t, vendor.VendorID)
		CreateRandomProductVariants(t, product.ProductID)
	}
	data, err := testStore.GetProductVariantsByStatus(context.Background(), "ACTIVE")
	require.NoError(t, err)
	require.NotEmpty(t, data)
	require.GreaterOrEqual(t, len(data), 3)
}

func TestUpdateProductVariant(t *testing.T) {
	vendor := CreateRandomVendors(t)
	product := CreateRandomProduct(t, vendor.VendorID)
	pv := CreateRandomProductVariants(t, product.ProductID)

	var jsondata *gofakeit.JSONOptions
	json, err := gofakeit.JSON(jsondata)
	require.NoError(t, err)

	newData := UpdateProductVariantParams{
		PvID:       pv.PvID,
		Price:      util.GenerateNumeric(),
		Stock:      util.GenerateInt32(),
		Attributes: json,
		Status:     util.GeneratePVStatus(),
	}

	newPV, err := testStore.UpdateProductVariant(context.Background(), newData)
	require.NoError(t, err)
	require.NotEmpty(t, newPV)
	require.Equal(t, pv.PvID, newPV.PvID)
	require.NotEqual(t, pv.Price, newPV.Price)
	require.NotEqual(t, pv.Stock, newPV.Stock)
	require.NotEqual(t, pv.Attributes, newPV.Attributes)
	require.NotEqual(t, pv.Status, newPV.Status)
	require.NotEmpty(t, newPV.UpdatedAt)
}
