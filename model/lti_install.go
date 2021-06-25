package model

type LtiInstall struct {
	ID       int64
	ClientID string `pg:",notnull"`
	Iss      string `pg:",notnull"`
	JwksUrl  string `pg:",notnull"`
	OidcUrl  string `pg:",notnull"`
	TokenUrl string `pg:",notnull"`

	LtiDeployments []*LtiDeployment `pg:"rel:has-many"`

	ApplicationID int64 `pg:"on_delete:CASCADE"`

	Timestamps
}
