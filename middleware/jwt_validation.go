package middleware

import (
	"context"
	"github.com/lestrrat-go/jwx/jwt"
	"net/http"
)

func newJwtValidator(next http.Handler, key contextKey, options ...jwt.ParseOption) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := jwt.ParseRequest(r,
			options...,
		)

		if err != nil {
			panic(err)
		}

		//Totally undocumented why a context is needed here
		claims, err := token.AsMap(context.Background())

		if err != nil {
			panic(err)
		}

		newCtx := context.WithValue(r.Context(), key, claims)
		newReq := r.WithContext(newCtx)
		next.ServeHTTP(w, newReq)
	})
}
