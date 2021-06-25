package controllers

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/atomicjolt/atomiclti/config"
	"github.com/atomicjolt/atomiclti/middleware"
	"github.com/atomicjolt/atomiclti/resources"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(controllerResources resources.Resources) http.Handler {
	router := mux.NewRouter()

	graphqlRouter := router.Methods("GET", "POST").Subrouter()
	graphqlRouter.Handle("/graphql", NewGraphqlHandler())
	graphqlRouter.Handle("/graphql/playground", playground.Handler("Playground", "/graphql"))
	graphqlRouter.Use(middleware.LaunchTokenFromAuth)

	router.Handle("/oidc_init", NewOpenIDInitHandler()).Methods("GET", "POST")

	ltiRouter := router.Methods("GET", "POST").Subrouter()
	ltiRouter.Handle("/lti_launches", NewLtiLaunchHandler())
	ltiRouter.Use(middleware.LaunchTokenFromIdToken, middleware.OidcStateValidator)

	router.Handle("/jwks", NewJwksController())

	if config.DetermineEnv() == "development" {
		router.Handle("/{path:.*}", NewHotReloadProxy("http://127.0.0.1:3000"))
	} else {
		router.Handle("/{path:.*}", NewClientFilesHandler())
	}

	router.Use(
		handlers.RecoveryHandler(),
		middleware.Logger,
		middleware.WithResources(controllerResources),
	)
	return router
}
