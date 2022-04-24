package inmem

import (
	"sync"

	"github.com/AntonioMorales97/goat"
)

type InmemStorage struct {
	blockedTokens map[string]bool
	blockMutex    *sync.RWMutex
	cachedTokens  map[string]string
	cacheMutex    *sync.RWMutex
}

func NewInmemStore() goat.TokenStore {
	inmemStorage := &InmemStorage{
		blockedTokens: make(map[string]bool),
		cachedTokens:  make(map[string]string),
		blockMutex:    &sync.RWMutex{},
		cacheMutex:    &sync.RWMutex{},
	}

	return inmemStorage
}

func (inmemStorage *InmemStorage) StoreToken(tokenId string, token string) {
	//TODO: Check that it doesn't already exists - small chance tho...
	inmemStorage.cacheMutex.Lock()
	inmemStorage.cachedTokens[tokenId] = token
	inmemStorage.cacheMutex.Unlock()
}

func (inmemStorage *InmemStorage) GetToken(tokenId string) (string, error) {
	accessToken, ok := inmemStorage.cachedTokens[tokenId]
	if !ok {
		return "", goat.NoTokenFound
	}

	return accessToken, nil
}

func (inmemStorage *InmemStorage) DeleteToken(tokenId string) {
	inmemStorage.cacheMutex.Lock()
	delete(inmemStorage.cachedTokens, tokenId)
	inmemStorage.cacheMutex.Unlock()
}

func (inmemStorage *InmemStorage) DeleteAndBlockToken(tokenId string) {
	inmemStorage.DeleteToken(tokenId)
	inmemStorage.blockMutex.Lock()
	inmemStorage.blockedTokens[tokenId] = true
	inmemStorage.blockMutex.Unlock()
}

func (inmemStorage *InmemStorage) IsTokenBlocked(tokenId string) bool {
	return inmemStorage.blockedTokens[tokenId]
}
