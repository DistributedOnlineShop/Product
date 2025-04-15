package db

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"

	"Product/util"
)

func CreateRandomInventoryAdjustments(t *testing.T, productID, pvID string) InventoryAdjustment {
	data := CreateInventoryAdjustmentsParams{
		AdjustmentID:   util.CreateUUID(),
		ProductID:      productID,
		PvID:           pvID,
		AdjustmentType: util.GenerateAdjustmentTypes(),
		Quantity:       util.GenerateInt32(),
		Reason:         gofakeit.Digit(),
	}

	ia, err := testStore.CreateInventoryAdjustments(context.Background(), data)
	require.NoError(t, err)
	require.NotEmpty(t, ia)
	require.Equal(t, data.AdjustmentID, ia.AdjustmentID)
	require.Equal(t, data.ProductID, ia.ProductID)
	require.Equal(t, data.PvID, ia.PvID)
	require.Equal(t, data.AdjustmentType, ia.AdjustmentType)
	require.Equal(t, data.Quantity, ia.Quantity)
	require.Equal(t, data.Reason, ia.Reason)
	require.NotZero(t, ia.CreatedAt)

	return ia
}

func TestCreateInventoryAdjustments(t *testing.T) {
	vendor := CreateRandomVendors(t)
	product := CreateRandomProduct(t, vendor.VendorID)
	pv := CreateRandomProductVariants(t, product.ProductID)
	CreateRandomInventoryAdjustments(t, product.ProductID, pv.PvID)
}

func TestGetInventoryAdjustmentsById(t *testing.T) {
	vendor := CreateRandomVendors(t)
	product := CreateRandomProduct(t, vendor.VendorID)
	pv := CreateRandomProductVariants(t, product.ProductID)
	ia := CreateRandomInventoryAdjustments(t, product.ProductID, pv.PvID)

	got, err := testStore.GetInventoryAdjustmentsByAdjustmentById(context.Background(), ia.AdjustmentID)
	require.NoError(t, err)
	require.NotEmpty(t, got)
}

func TestGetInventoryAdjustmentsByPvid(t *testing.T) {
	vendor := CreateRandomVendors(t)
	product := CreateRandomProduct(t, vendor.VendorID)
	pv := CreateRandomProductVariants(t, product.ProductID)
	CreateRandomInventoryAdjustments(t, product.ProductID, pv.PvID)

	data := GetInventoryAdjustmentsByPvidParams{
		ProductID: product.ProductID,
		PvID:      pv.PvID,
	}

	got, err := testStore.GetInventoryAdjustmentsByPvid(context.Background(), data)
	require.NoError(t, err)
	require.NotEmpty(t, got)
}

func TestGetInventoryAdjustmentsByType(t *testing.T) {
	for i := 0; i < 20; i++ {
		vendor := CreateRandomVendors(t)
		product := CreateRandomProduct(t, vendor.VendorID)
		pv := CreateRandomProductVariants(t, product.ProductID)
		CreateRandomInventoryAdjustments(t, product.ProductID, pv.PvID)
	}

	got, err := testStore.GetInventoryAdjustmentsByType(context.Background(), util.GenerateAdjustmentTypes())
	require.NoError(t, err)
	require.NotEmpty(t, got)
	require.GreaterOrEqual(t, len(got), 1)
}
