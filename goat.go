package goat

type Goat struct {
	TokenMaker TokenMaker
	TokenStore TokenStore
}

func NewGoat(tokenMaker TokenMaker, tokenStore TokenStore) *Goat {
	return &Goat{
		tokenMaker,
		tokenStore,
	}
}
