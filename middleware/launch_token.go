package middleware

import (
	"context"
	"github.com/atomicjolt/atomiclti/lti"
	"github.com/lestrrat-go/jwx/jwt"
	"net/http"
)

func LaunchTokenFromIdToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		controllerResources := GetResources(r.Context())
		var iss, clientId string

		if err := r.ParseForm(); err != nil {
			panic(err)
		}

		/**
		 * Inside this scope, we peek into the id_token before it is validated in order
		 * to get the Client ID and ISS of the tool consumer in order to find their JWKs
		 * URL. If possible, use the token and its claims after it has been verified.
		 */
		{
			token, err := jwt.ParseRequest(r, jwt.WithFormKey("id_token"))

			if err != nil {
				panic(err)
			}

			iss = token.Issuer()
			clientId = token.Audience()[0]
		}

		ltiInstall, err := controllerResources.Repo.LtiInstall.From(iss, clientId)

		if err != nil {
			panic(err)
		}

		keySet, err := controllerResources.ToolConsumerJwks.ForInstall(r.Context(), ltiInstall)

		if err != nil {
			panic("Unable to get public JWK set for tool consumer.")
		}

		idToken, err := jwt.ParseRequest(r,
			jwt.WithFormKey("id_token"),
			jwt.WithValidate(true),
			jwt.WithKeySet(keySet),
		)

		if err != nil {
			panic(err)
		}

		launchToken := lti.NewLaunchToken(lti.NewIdToken(idToken))

		newCtx := context.WithValue(r.Context(), launchTokenKey, launchToken)

		next.ServeHTTP(w, r.WithContext(newCtx))
	})
}

func LaunchTokenFromAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		controllerResources := GetResources(r.Context())
		keySet, err := controllerResources.Repo.Jwk.PublicJwkSet()

		if err != nil {
			panic(err)
		}

		launchToken, err := lti.ParseLaunchToken(r, jwt.WithKeySet(keySet))

		if err != nil {
			panic(err)
		}

		newCtx := context.WithValue(r.Context(), launchTokenKey, launchToken)
		next.ServeHTTP(w, r.WithContext(newCtx))
	})
}

func GetLaunchToken(ctx context.Context) *lti.LaunchToken {
	return ctx.Value(launchTokenKey).(*lti.LaunchToken)
}
