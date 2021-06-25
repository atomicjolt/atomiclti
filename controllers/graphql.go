package controllers

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/atomicjolt/atomiclti/graph"
	"github.com/atomicjolt/atomiclti/graph/generated"
	"net/http"
)

func NewGraphqlHandler() http.Handler {
	return handler.NewDefaultServer(
		generated.NewExecutableSchema(generated.Config{
			Resolvers: &graph.Resolver{},
		}),
	)
}
