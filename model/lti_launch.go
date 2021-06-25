package model

import "github.com/atomicjolt/atomiclti/config"

type LtiLaunch struct {
	ID                       int64
	DeploymentID             string                    `pg:",notnull"`
	Config                   *config.ApplicationConfig `pg:"type:jsonb,notnull"`
	ContextID                string                    `pg:",notnull"`
	ResourceLinkID           string                    `pg:",notnull"`
	Token                    string                    `pg:",notnull"`
	ToolConsumerInstanceGuid string                    `pg:",notnull"`

	ApplicationInstanceID int64 `pg:"on_delete:CASCADE"`

	Timestamps
}
