package util

import (
	"fmt"
	"math/big"
	"math/rand/v2"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/jackc/pgx/v5/pgtype"
)

func GenerateRandomEmail() pgtype.Text {
	return pgtype.Text{
		String: gofakeit.Email(),
		Valid:  true,
	}
}

func GenerateRandomPhone() pgtype.Text {
	return pgtype.Text{
		String: gofakeit.Phone(),
		Valid:  true,
	}
}

func GenerateRandomDate() pgtype.Timestamp {
	daysOffset := rand.IntN(365) - 180
	return pgtype.Timestamp{
		Time:  time.Now().Add(time.Duration(daysOffset) * 24 * time.Hour),
		Valid: true,
	}
}

func GenerateRandomNumeric() pgtype.Numeric {
	intPart := rand.IntN(100000)
	fracPart := rand.IntN(100)
	value := int64(intPart) + int64(fracPart)/100.0

	return pgtype.Numeric{
		Int:   big.NewInt(value),
		Exp:   -2,
		Valid: true,
	}
}

func GenerateRandomStatus() string {
	statuses := []string{
		"ACTIVE",
		"INACTIVE",
		"PENDING",
		"SUSPENDED",
		"BLOCKED",
		"ARCHIVED",
	}
	return statuses[rand.IntN(len(statuses))]
}

func GenerateRandomVendorCategory() []string {
	categories := []string{
		"ELECTRONICS",
		"APPAREL & ACCESSORIES",
		"HOME GOODS",
		"BEAUTY & HEALTH",
		"FOOD & BEVERAGE",
		"TOYS & GAMES",
		"SPORTS & OUTDOORS",
		"BOOKS & MUSIC",
		"AUTOMOTIVE",
		"DIGITAL GOODS",
	}

	i := rand.IntN(3) + 1
	if i > 1 {
		arr := []string{}
		for k := 0; k <= i; k++ {
			arr = append(arr, categories[rand.IntN(len(categories))])
		}
		return arr
	} else {
		return []string{categories[rand.IntN(len(categories))]}
	}
}

func GenerateRandomInt32() int32 {
	return rand.Int32N(1000) + 1
}

func GenerateRandomProductStatus() string {
	statuses := []string{
		"ACTIVE",
		"INACTIVE",
		"OUT OF STOCK",
		"DISCONTINUED",
		"PENDING",
		"DRAFT",
		"PRE-ORDER",
	}

	return statuses[rand.IntN(len(statuses))]
}

func GenerateRandomSKU() pgtype.Text {
	colors := []string{"RED", "BLUE", "GREEN", "BLACK", "WHITE"}
	sizes := []string{"S", "M", "L", "XL", "XXL"}
	materials := []string{"COTTON", "POLYESTER", "LEATHER"}
	styles := []string{"T_SHIRT", "HOODIE", "JACKET", "SHIRT"}

	color := colors[rand.IntN(len(colors))]
	size := sizes[rand.IntN(len(sizes))]
	material := materials[rand.IntN(len(materials))]
	style := styles[rand.IntN(len(styles))]

	return pgtype.Text{
		String: fmt.Sprintf("%s_%s_%s_%s", style, color, material, size),
		Valid:  true,
	}
}

func GenerateRandomPVStatus() string {
	statuses := []string{"ACTIVE", "INACTIVE", "OUT OF STOCK", "PENDING", "DISCONTINUED"}
	return statuses[rand.IntN(len(statuses))]
}
