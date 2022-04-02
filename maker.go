package goat

type Maker interface {
	Encrypt(payload interface{}) (string, error)
	VerifyToken(token string, payload interface{}) error
}
