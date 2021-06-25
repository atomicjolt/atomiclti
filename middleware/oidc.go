package middleware

import (
	"context"
	"github.com/atomicjolt/atomiclti/config"
	"github.com/lestrrat-go/jwx/jwt"
	"net/http"
)

func OidcStateValidator(next http.Handler) http.Handler {
	return newJwtValidator(next,
		oidcStateKey,
		jwt.WithFormKey("state"),
		jwt.WithValidate(true),
		jwt.WithKeySet(config.LtiKeySet()),
	)
}

func GetOidcState(ctx context.Context) map[string]interface{} {
	return ctx.Value(oidcStateKey).(map[string]interface{})
}
