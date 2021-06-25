package config

import (
	"encoding/json"
	"github.com/atomicjolt/atomiclti/lib"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwk"
	"log"
)

type LtiPlacement struct {
	Placement       string `json:"placement"`
	TargetLinkUri   string `json:"target_link_uri"`
	Text            string `json:"text"`
	Enabled         bool   `json:"enabled"`
	IconUrl         string `json:"icon_url"`
	MessageType     string `json:"message_type"`
	CanvasIconClass string `json:"canvas_icon_class"`
}

type LtiAdvantageExtensionSetting struct {
	Text            string         `json:"text"`
	IconUrl         string         `json:"icon_url"`
	SelectionWidth  int            `json:"selection_width"`
	SelectionHeight int            `json:"selection_height"`
	Placements      []LtiPlacement `json:"placements"`
}

type LtiAdvantageExtension struct {
	Platform     string                       `json:"platform"`
	Domain       string                       `json:"domain"`
	ToolId       string                       `json:"tool_id"`
	PrivacyLevel string                       `json:"privacy_level"`
	Settings     LtiAdvantageExtensionSetting `json:"settings"`
}

type LtiAdvantageConfig struct {
	Title             string                  `json:"title"`
	Scopes            []string                `json:"scopes"`
	Extensions        []LtiAdvantageExtension `json:"extensions"`
	TargetLinkUri     string                  `json:"target_link_uri"`
	OidcInitiationUrl string                  `json:"oidc_initiation_url"`
	PublicJwk         string                  `json:"public_jwk"`
	Description       string                  `json:"description"`
	CustomFields      []map[string]string
}

func GetLtiAdvantageConfig() LtiAdvantageConfig {
	var ltiAdvantageConfig LtiAdvantageConfig

	err := json.Unmarshal(lib.LoadJsonFrom("./lti_advantage_config.json"), &ltiAdvantageConfig)

	if err != nil {
		log.Fatal("LTI config file is not valid json: " + err.Error())
	}

	return ltiAdvantageConfig
}

func LtiKeySet() jwk.Set {
	authClientSecret := GetServerConfig().AuthClientSecret
	authKey := jwk.NewSymmetricKey()

	err := authKey.FromRaw(authClientSecret)
	if err != nil {
		log.Fatal(err)
	}

	if err := authKey.Set(jwk.AlgorithmKey, jwa.HS512); err != nil {
		log.Fatal(err)
	}
	if err := authKey.Set(jwk.KeyIDKey, "atomiclti_auth0"); err != nil {
		log.Fatal(err)
	}

	keyset := jwk.NewSet()
	keyset.Add(authKey)

	return keyset
}
