package inmem

import (
	"testing"

	"github.com/AntonioMorales97/goat"
	"github.com/stretchr/testify/require"
)

func TestGetStoredToken(t *testing.T) {
	store := NewInmemStore()
	tokenId := "token_id"
	token := "something_encrypted"
	storeToken(t, store, tokenId, token)
}

func storeToken(t *testing.T, store goat.TokenStore, tokenId string, token string) {
	store.StoreToken(tokenId, token)
	fetchedToken, err := store.GetToken(tokenId)

	require.NoError(t, err)
	require.NotEmpty(t, fetchedToken)
	require.Equal(t, token, fetchedToken)
}

func TestGetUnexistingToken(t *testing.T) {
	store := NewInmemStore()
	tokenShouldNotExist(t, store, "does_not_exist")
}

func tokenShouldNotExist(t *testing.T, store goat.TokenStore, tokenId string) {
	fetchedToken, err := store.GetToken(tokenId)

	require.Error(t, err)
	require.Equal(t, err, goat.NoTokenFound)
	require.Empty(t, fetchedToken)
}

func TestDeleteToken(t *testing.T) {
	store := NewInmemStore()
	tokenId := "token_id"
	token := "something_encrypted"
	storeToken(t, store, tokenId, token)
	store.DeleteToken(tokenId)
	tokenShouldNotExist(t, store, tokenId)
}

func TestDeleteAndBlockToken(t *testing.T) {
	store := NewInmemStore()
	tokenId := "token_id"
	token := "something_encrypted"
	storeToken(t, store, tokenId, token)
	store.DeleteAndBlockToken(tokenId)
	tokenShouldNotExist(t, store, tokenId)
	isBlocked := store.IsTokenBlocked(tokenId)
	require.True(t, isBlocked)
}
