package model

type LtiDeployment struct {
	ID           int64
	DeploymentID string `pg:",notnull"`

	LtiInstallID          int64 `pg:"on_delete:CASCADE"`
	ApplicationInstanceID int64 `pg:"on_delete:CASCADE"`

	Timestamps
}
