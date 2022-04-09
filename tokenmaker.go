package goat

type TokenMaker interface {
	Encrypt(payload interface{}) (string, error)
	VerifyToken(token string, payload interface{}) error
}
