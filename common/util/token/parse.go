package tokenutil

import (
	"context"

	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

func Parse(ctx context.Context, tokenRaw string, jwkUrl string) (jwt.Token, error) {
	key, err := jwk.Fetch(ctx, jwkUrl)
	if err != nil {
		return nil, err
	}
	return jwt.Parse([]byte(tokenRaw), jwt.WithKeySet(key))
}
