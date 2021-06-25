package main

import (
	"time"

	"github.com/go-pg/pg/v10/orm"
	migrations "github.com/robinjoseph08/go-pg-migrations/v3"
)

type ApplicationConfig struct{}
type LtiAdvantageConfig struct{}

type Timestamps struct {
	CreatedAt time.Time `pg:",notnull,default:now()"`
	UpdatedAt time.Time `pg:",notnull,default:now()"`
}

type LtiDeployment struct {
	ID           int64
	DeploymentID string `pg:",notnull"`

	LtiInstallID          int64 `pg:"on_delete:CASCADE"`
	ApplicationInstanceID int64 `pg:"on_delete:CASCADE"`

	Timestamps
}

type User struct {
	ID          int64
	Name        string `pg:",notnull"`
	Email       string `pg:",notnull"`
	LmsUserId   string `pg:",notnull"`
	LtiUserId   string `pg:",notnull"`
	LtiProvider string `pg:",notnull"`

	ApplicationInstanceID int64 `pg:"on_delete:CASCADE"`

	Timestamps
}

type LtiLaunch struct {
	ID                       int64
	DeploymentID             string             `pg:",notnull"`
	Config                   *ApplicationConfig `pg:"type:jsonb,notnull"`
	ContextID                string             `pg:",notnull"`
	ResourceLinkID           string             `pg:",notnull"`
	Token                    string             `pg:",notnull"`
	ToolConsumerInstanceGuid string             `pg:",notnull"`

	ApplicationInstanceID int64 `pg:"on_delete:CASCADE"`

	Timestamps
}

type ApplicationInstance struct {
	ID                    int64
	ClientApplicationName string             `pg:",notnull"`
	Config                *ApplicationConfig `pg:"type:jsonb,notnull"`
	Description           string             `pg:",notnull"`
	Key                   string             `pg:",notnull"`
	LtiDeployments        []*LtiDeployment   `pg:"rel:has-many"`
	LtiLaunches           []*LtiLaunch       `pg:"rel:has-many"`
	Users                 []*User            `pg:"rel:has-many"`

	ApplicationID int64 `pg:"on_delete:CASCADE"`

	Timestamps
}

type Jwk struct {
	ID  int64
	Kid string `pg:",notnull"`
	Pem string `pg:",notnull"`

	ApplicationID int64 `pg:"on_delete:CASCADE"`

	Timestamps
}

type OpenIdState struct {
	ID int64

	Nonce string

	Timestamps
}

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

type Application struct {
	ID                    int64
	Name                  string                 `pg:",notnull"`
	ClientApplicationName string                 `pg:",notnull"`
	DefaultConfig         *ApplicationConfig     `pg:"type:jsonb,notnull"`
	Description           string                 `pg:",notnull"`
	Key                   string                 `pg:",notnull"`
	LtiAdvantageConfig    *LtiAdvantageConfig    `pg:",notnull"`
	ApplicationInstances  []*ApplicationInstance `pg:"rel:has-many"`
	Jwks                  []*Jwk                 `pg:"rel:has-many"`
	LtiInstalls           []*LtiInstall          `pg:"rel:has-many"`

	Timestamps
}

var models []interface{} = []interface{}{
	(*LtiDeployment)(nil),
	(*User)(nil),
	(*LtiLaunch)(nil),
	(*ApplicationInstance)(nil),
	(*Jwk)(nil),
	(*OpenIdState)(nil),
	(*LtiInstall)(nil),
	(*Application)(nil),
}

func init() {
	up := func(db orm.DB) error {
		for _, model := range models {
			err := db.Model(model).CreateTable(&orm.CreateTableOptions{FKConstraints: true})
			if err != nil {
				return err
			}
		}
		return nil
	}

	down := func(db orm.DB) error {
		for _, model := range models {
			err := db.Model(model).DropTable(&orm.DropTableOptions{})
			if err != nil {
				return err
			}
		}

		return nil
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20210412211919_create_lti_schema", up, down, opts)
}
