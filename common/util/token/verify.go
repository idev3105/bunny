package tokenutil

import (
	"context"

	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jws"
)

// Verify token return true if token is valid end error or false if token is invalid
func Verify(ctx context.Context, tokenRaw string, jwkUrl string) (bool, error) {
	key, err := jwk.Fetch(ctx, jwkUrl)
	if err != nil {
		return false, err
	}
	_, err = jws.Verify([]byte(tokenRaw), jws.WithKeySet(key))
	if err != nil {
		return false, err
	}

	return true, nil
}
