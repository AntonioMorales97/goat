package goat

import "errors"

var (
	NoTokenFound = errors.New("no token found for the given token id")
)

type TokenStore interface {
	// Stores the given token with the token id
	StoreToken(tokenId string, token string)

	// Gets the token for the given token id or returns an error if it doesn't exists
	GetToken(tokenId string) (string, error)

	// Deletes token without moving it to the blocked tokens
	DeleteToken(tokenId string)

	// Deletes token and moves it to the blocked tokens
	DeleteAndBlockToken(tokenId string)

	// Returns true if token with the given token id is blocked
	IsTokenBlocked(tokenId string) bool
}
