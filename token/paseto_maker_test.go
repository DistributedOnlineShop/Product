package token

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"

	"UsersManagement/util"
)

func TestPasetoMaker(t *testing.T) {

	config, err := util.LoadConfig("../")
	require.NoError(t, err)

	maker, err := NewPasetoMaker(config.KeySeed)
	require.NoError(t, err)

	email := gofakeit.Email()
	role := "user"

	issuedAt := time.Now()

	token, payload, err := maker.CreateToken(email, role)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	require.NotZero(t, payload.ID)
	require.Equal(t, email, payload.Email)
	require.Equal(t, role, payload.Role)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
}
