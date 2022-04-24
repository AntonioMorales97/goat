package tokenmaker

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

type pasetoPayload struct {
	ExpiredAt time.Time
	Data      interface{}
}

func (payload *pasetoPayload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return goat.ErrExpiredToken
	}

	return nil
}

func NewPasetoMaker(symmetricKey string, duration time.Duration) (goat.TokenMaker, error) {
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
	pasetoPayload := &pasetoPayload{
		ExpiredAt: time.Now().Add(maker.duration),
		Data:      payload,
	}
	return maker.paseto.Encrypt(maker.symmetricKey, pasetoPayload, nil)
}

func (maker *PasetoMaker) VerifyToken(token string, payload interface{}) error {
	pasetoPayload := &pasetoPayload{
		Data: payload,
	}

	err := maker.paseto.Decrypt(token, maker.symmetricKey, pasetoPayload, nil)
	if err != nil {
		return goat.ErrInvalidToken
	}

	err = pasetoPayload.Valid()
	if err != nil {
		return err
	}

	return nil
}
