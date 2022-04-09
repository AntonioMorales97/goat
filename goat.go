package goat

type Goat struct {
	tokenMaker TokenMaker
	tokenStore TokenStore
}

func NewGoat(tokenMaker TokenMaker, tokenStore TokenStore) *Goat {
	return &Goat{
		tokenMaker,
		tokenStore,
	}
}
