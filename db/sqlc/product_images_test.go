package db

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"

	"Product/util"
)

func CreateRandomProductImage(t *testing.T, productID, pvID string) ProductImage {

	data := CreateProductImageParams{
		ProductID: productID,
		PvID:      pvID,
		ImageUrl:  "https://example.com/image.jpg",
		Position:  util.GenerateInt32(),
		IsPrimary: util.GenerateBool(),
	}

	pi, err := testStore.CreateProductImage(context.Background(), data)
	require.NoError(t, err)
	require.NotEmpty(t, pi.PiID)
	require.Equal(t, productID, pi.ProductID)
	require.Equal(t, pvID, pi.PvID)
	require.Equal(t, data.ImageUrl, pi.ImageUrl)
	require.Equal(t, data.Position, pi.Position)
	require.Equal(t, data.IsPrimary, pi.IsPrimary)
	require.NotZero(t, pi.CreatedAt)

	return pi
}

func TestCreateRandomProductImage(t *testing.T) {
	vendor := CreateRandomVendors(t)
	product := CreateRandomProduct(t, vendor.VendorID)
	pv := CreateRandomProductVariants(t, product.ProductID)
	CreateRandomProductImage(t, product.ProductID, pv.PvID)
}

func TestDeleteProductImage(t *testing.T) {
	vendor := CreateRandomVendors(t)
	product := CreateRandomProduct(t, vendor.VendorID)
	pv := CreateRandomProductVariants(t, product.ProductID)
	pi := CreateRandomProductImage(t, product.ProductID, pv.PvID)

	err := testStore.DeleteProductImage(context.Background(), pi.PiID)
	require.NoError(t, err)
}

func TestGetProductImageByPiid(t *testing.T) {
	vendor := CreateRandomVendors(t)
	product := CreateRandomProduct(t, vendor.VendorID)
	pv := CreateRandomProductVariants(t, product.ProductID)
	pi := CreateRandomProductImage(t, product.ProductID, pv.PvID)

	images, err := testStore.GetProductImageByPiid(context.Background(), pi.PiID)
	require.NoError(t, err)
	require.Len(t, images, 1)
}

func TestGetProductImageByProductid(t *testing.T) {
	vendor := CreateRandomVendors(t)
	product := CreateRandomProduct(t, vendor.VendorID)
	for i := 0; i < 10; i++ {
		pv := CreateRandomProductVariants(t, product.ProductID)
		CreateRandomProductImage(t, product.ProductID, pv.PvID)
	}

	images, err := testStore.GetProductImageByProductid(context.Background(), product.ProductID)
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(images), 10)
}

func TestGetProductImageByPvid(t *testing.T) {
	vendor := CreateRandomVendors(t)
	product := CreateRandomProduct(t, vendor.VendorID)
	pv := CreateRandomProductVariants(t, product.ProductID)
	for i := 0; i < 10; i++ {
		CreateRandomProductImage(t, product.ProductID, pv.PvID)
	}

	data := GetProductImageByPvidParams{
		ProductID: product.ProductID,
		PvID:      pv.PvID,
	}

	images, err := testStore.GetProductImageByPvid(context.Background(), data)
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(images), 10)
}

func TestGetProductImagePrimary(t *testing.T) {
	for i := 0; i < 10; i++ {
		vendor := CreateRandomVendors(t)
		product := CreateRandomProduct(t, vendor.VendorID)
		pv := CreateRandomProductVariants(t, product.ProductID)
		CreateRandomProductImage(t, product.ProductID, pv.PvID)
	}

	images, err := testStore.GetProductImagePrimary(context.Background(), pgtype.Bool{Bool: util.GenerateBool().Bool, Valid: true})
	require.NoError(t, err)
	require.NotEmpty(t, images)
	require.GreaterOrEqual(t, len(images), 2)
}

func TestUpdateProductImage(t *testing.T) {
	vendor := CreateRandomVendors(t)
	product := CreateRandomProduct(t, vendor.VendorID)
	pv := CreateRandomProductVariants(t, product.ProductID)
	pi := CreateRandomProductImage(t, product.ProductID, pv.PvID)

	newData := UpdateProductImageParams{
		PiID:      pi.PiID,
		ImageUrl:  "https://example.com/new_image2.jpg",
		Position:  util.GenerateInt32(),
		IsPrimary: util.GenerateBool(),
	}

	newImage, err := testStore.UpdateProductImage(context.Background(), newData)
	require.NoError(t, err)
	require.NotEmpty(t, newImage)
	require.Equal(t, newImage.PiID, pi.PiID)
}
