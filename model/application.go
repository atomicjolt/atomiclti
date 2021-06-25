package model

import "github.com/atomicjolt/atomiclti/config"

type Application struct {
	ID                    int64
	Name                  string                     `pg:",notnull"`
	ClientApplicationName string                     `pg:",notnull"`
	DefaultConfig         *config.ApplicationConfig  `pg:"type:jsonb,notnull"`
	Description           string                     `pg:",notnull"`
	Key                   string                     `pg:",notnull"`
	LtiAdvantageConfig    *config.LtiAdvantageConfig `pg:",notnull"`
	ApplicationInstances  []*ApplicationInstance     `pg:"rel:has-many"`
	Jwks                  []*Jwk                     `pg:"rel:has-many"`
	LtiInstalls           []*LtiInstall              `pg:"rel:has-many"`

	Timestamps
}
