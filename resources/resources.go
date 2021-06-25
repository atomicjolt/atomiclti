package resources

import (
	"context"
	"github.com/atomicjolt/atomiclti/model"
	"github.com/atomicjolt/atomiclti/repo"
	"github.com/lestrrat-go/jwx/jwk"
	"log"
)

type ToolConsumerJwks struct {
	ToolConsumerAutoRefresh *jwk.AutoRefresh
}

func NewToolConsumerJwks(ar *jwk.AutoRefresh) ToolConsumerJwks {
	return ToolConsumerJwks{
		ToolConsumerAutoRefresh: ar,
	}
}

func (toolConsumerJwks *ToolConsumerJwks) ForInstall(ctx context.Context, inst *model.LtiInstall) (jwk.Set, error) {
	keySet, err := toolConsumerJwks.ToolConsumerAutoRefresh.Fetch(ctx, inst.JwksUrl)
	if err != nil {
		return nil, err
	}

	return keySet, nil
}

type Resources struct {
	Repo             *repo.Repo
	ToolConsumerJwks ToolConsumerJwks
}

func NewResources() (Resources, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	rootRepo := repo.NewRepo()

	toolConsumerAutoRefresh, err := rootRepo.LtiInstall.NewAutoRefresh(ctx)

	if err != nil {
		log.Fatal("Could not initialize JWK auto-refresh.")
	}

	return Resources{
		Repo:             rootRepo,
		ToolConsumerJwks: NewToolConsumerJwks(toolConsumerAutoRefresh),
	}, cancel
}
