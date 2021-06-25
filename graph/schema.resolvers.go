package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/atomicjolt/atomiclti/graph/generated"
	"github.com/atomicjolt/atomiclti/middleware"
)

func (r *queryResolver) HelloWorld(ctx context.Context) (string, error) {
	ltiToken := middleware.GetLaunchToken(ctx)
	contextId := ltiToken.ContextId()

	return fmt.Sprintf("Yo! Hello %s from the LTI GraphQL backend.", contextId), nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
