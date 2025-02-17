package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"Product/util"
)

func CreateRandomInventoryLevels(t *testing.T, productID, pvID string) InventoryLevel {
	data := CreateInventoryLevelsParams{
		InventoryID: util.CreateUUID(),
		ProductID:   productID,
		PvID:        pvID,
		Stock:       util.GenerateRandomInt32(),
	}

	inventoryLevel, err := testStore.CreateInventoryLevels(context.Background(), data)
	require.NoError(t, err)
	require.NotEmpty(t, inventoryLevel)
	require.Equal(t, inventoryLevel.InventoryID, data.InventoryID)
	require.Equal(t, inventoryLevel.ProductID, data.ProductID)
	require.Equal(t, inventoryLevel.PvID, data.PvID)
	require.Equal(t, inventoryLevel.Stock, data.Stock)
	require.NotZero(t, inventoryLevel.UpdatedAt)

	return inventoryLevel
}

func TestCreateInventoryLevels(t *testing.T) {
	vendor := CreateRandomVendors(t)
	product := CreateRandomProduct(t, vendor.VendorID)
	pv := CreateRandomProductVariants(t, product.ProductID)
	CreateRandomInventoryLevels(t, product.ProductID, pv.PvID)
}

func TestGetInventoryLevelByInventoryId(t *testing.T) {
	vendor := CreateRandomVendors(t)
	product := CreateRandomProduct(t, vendor.VendorID)
	pv := CreateRandomProductVariants(t, product.ProductID)
	iv := CreateRandomInventoryLevels(t, product.ProductID, pv.PvID)

	data, err := testStore.GetInventoryLevelByInventoryId(context.Background(), iv.InventoryID)
	require.NoError(t, err)
	require.NotEmpty(t, data)
}

func TestGetInventoryLevelByPvid(t *testing.T) {
	vendor := CreateRandomVendors(t)
	product := CreateRandomProduct(t, vendor.VendorID)
	pv := CreateRandomProductVariants(t, product.ProductID)
	CreateRandomInventoryLevels(t, product.ProductID, pv.PvID)

	data := GetInventoryLevelByPvidParams{
		ProductID: product.ProductID,
		PvID:      pv.PvID,
	}

	iv, err := testStore.GetInventoryLevelByPvid(context.Background(), data)
	require.NoError(t, err)
	require.NotEmpty(t, iv)
}

func TestUpdateInventoryLevel(t *testing.T) {
	vendor := CreateRandomVendors(t)
	product := CreateRandomProduct(t, vendor.VendorID)
	pv := CreateRandomProductVariants(t, product.ProductID)
	iv := CreateRandomInventoryLevels(t, product.ProductID, pv.PvID)

	newData := UpdateInventoryLevelParams{
		InventoryID: iv.InventoryID,
		Stock:       util.GenerateRandomInt32(),
	}

	updatedIV, err := testStore.UpdateInventoryLevel(context.Background(), newData)
	require.NoError(t, err)
	require.NotEmpty(t, updatedIV)
	require.Equal(t, updatedIV.InventoryID, iv.InventoryID)
	require.NotEqual(t, updatedIV.Stock, iv.Stock)
}
