package goat

import "errors"

var (
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("token is invalid")
)

type TokenMaker interface {
	Encrypt(payload interface{}) (string, error)
	VerifyToken(token string, payload interface{}) error
}
