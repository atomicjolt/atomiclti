package controllers

import (
	"encoding/json"
	"github.com/atomicjolt/atomiclti/middleware"
	"net/http"
)

func NewJwksController() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		controllerResources := middleware.GetResources(r.Context())
		publicKeys, err := controllerResources.Repo.Jwk.PublicJwkSet()

		if err != nil {
			panic(err)
		}

		enc := json.NewEncoder(w)
		enc.Encode(publicKeys)
	})
}
