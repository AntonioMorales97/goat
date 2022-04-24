package tokenmaker

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var symmetricKey = "01234567890123456789012345678912"

type TestPayload struct {
	FirstName string
	LastName  string
	Age       int
}

func newTestPayload() *TestPayload {
	return &TestPayload{
		FirstName: "John",
		LastName:  "Doe",
		Age:       40,
	}
}

func TestPasetoMaker(t *testing.T) {
	duration := time.Minute
	maker, err := NewPasetoMaker(symmetricKey, duration)
	require.NoError(t, err)

	testPayload := newTestPayload()

	token, err := maker.Encrypt(testPayload)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload := &TestPayload{}
	err = maker.VerifyToken(token, payload)

	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.Equal(t, testPayload.FirstName, payload.FirstName)
	require.Equal(t, testPayload.LastName, payload.LastName)
	require.Equal(t, testPayload.Age, payload.Age)
}

func TestExpiredPasetoToken(t *testing.T) {
	duration := -time.Minute
	maker, err := NewPasetoMaker(symmetricKey, duration)
	require.NoError(t, err)

	testPayload := newTestPayload()
	require.NoError(t, err)

	token, err := maker.Encrypt(testPayload)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload := &TestPayload{}
	err = maker.VerifyToken(token, payload)

	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
}

func TestPasetoTokenWithDifferentMaker(t *testing.T) {
	duration := time.Minute
	maker, err := NewPasetoMaker(symmetricKey, duration)
	require.NoError(t, err)

	testPayload := newTestPayload()
	require.NoError(t, err)

	token, err := maker.Encrypt(testPayload)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	maker, err = NewPasetoMaker("09876543210987654321098765432109", duration)
	require.NoError(t, err)

	payload := &TestPayload{}
	err = maker.VerifyToken(token, payload)

	require.Error(t, err)
	require.EqualError(t, err, ErrInvalidToken.Error())
}
