package db

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"

	"Product/util"
)

func CreateRandomVendors(t *testing.T) Vendor {
	data := CreateVendorsParams{
		VendorID:    util.CreateUUID(),
		VendorName:  gofakeit.Company(),
		ContactName: gofakeit.Name(),
		ProductType: util.GenerateVendorCategory(),
		Email:       util.GenerateEmail(),
		Phone:       util.GeneratePhone(),
		Status:      util.GenerateStatus(),
	}

	vendor, err := testStore.CreateVendors(context.Background(), data)
	require.NoError(t, err)
	require.NotEmpty(t, vendor)
	require.Equal(t, vendor.VendorID, data.VendorID)
	require.Equal(t, vendor.VendorName, data.VendorName)
	require.Equal(t, vendor.ContactName, data.ContactName)
	require.Equal(t, vendor.ProductType, data.ProductType)
	require.Equal(t, vendor.Email, data.Email)
	require.Equal(t, vendor.Phone, data.Phone)
	require.Equal(t, vendor.Status, data.Status)
	require.NotZero(t, vendor.CreatedAt)

	return vendor
}

func TestCreateVendors(t *testing.T) {
	CreateRandomVendors(t)
}

func TestGetVendorsList(t *testing.T) {
	for i := 0; i < 20; i++ {
		CreateRandomVendors(t)
	}

	vendors, err := testStore.GetVendorsList(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, vendors)
}

func TestUpdateVendor(t *testing.T) {
	vendor := CreateRandomVendors(t)

	newData := UpdateVendorParams{
		VendorID:    vendor.VendorID,
		VendorName:  gofakeit.Company(),
		ContactName: gofakeit.Name(),
		Email:       util.GenerateEmail(),
		Phone:       util.GeneratePhone(),
		Status:      util.GenerateStatus(),
	}

	updatedVendor, err := testStore.UpdateVendor(context.Background(), newData)
	require.NoError(t, err)
	require.NotEmpty(t, updatedVendor)
	require.Equal(t, updatedVendor.VendorID, vendor.VendorID)
	require.Equal(t, updatedVendor.VendorName, newData.VendorName)
	require.Equal(t, updatedVendor.ContactName, newData.ContactName)
	require.Equal(t, updatedVendor.Email, newData.Email)
	require.Equal(t, updatedVendor.Phone, newData.Phone)
	require.Equal(t, updatedVendor.Status, newData.Status)
	require.NotZero(t, updatedVendor.UpdatedAt)
}
