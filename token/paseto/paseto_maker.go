package paseto

import (
	"fmt"
	"time"

	"github.com/AntonioMorales97/goat"
	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
	duration     time.Duration
}

func NewPasetoMaker(symmetricKey string, duration time.Duration) (goat.Maker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
		duration:     duration,
	}

	return maker, nil
}

func (maker *PasetoMaker) Encrypt(payload interface{}) (string, error) {
	goatPayload := &goatPayload{
		ExpiredAt: time.Now().Add(maker.duration),
		Data:      payload,
	}
	return maker.paseto.Encrypt(maker.symmetricKey, goatPayload, nil)
}

func (maker *PasetoMaker) VerifyToken(token string, payload interface{}) error {
	goatPayload := &goatPayload{
		Data: payload,
	}

	err := maker.paseto.Decrypt(token, maker.symmetricKey, goatPayload, nil)
	if err != nil {
		return ErrInvalidToken
	}

	err = goatPayload.Valid()
	if err != nil {
		return err
	}

	return nil
}
