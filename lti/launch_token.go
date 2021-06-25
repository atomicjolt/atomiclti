package lti

import (
	"context"
	"github.com/lestrrat-go/jwx/jwt"
	"net/http"
	"time"
)

type LaunchToken struct {
	jwt.Token
}

type InheritedClaim struct {
	key    string
	derive func(IdToken) interface{}
}

/**
 * Claims declared here will be inherited from
 * the id_token and deserialized on API calls
 */
func inheritedClaims() []InheritedClaim {
	return []InheritedClaim{
		{
			key: "contextId",
			derive: func(t IdToken) interface{} {
				return t.ContextId()
			},
		},
	}
}

func NewLaunchToken(idToken IdToken) *LaunchToken {
	token := jwt.New()

	token.Set(jwt.SubjectKey, idToken.Subject())
	token.Set(jwt.AudienceKey, idToken.Audience())
	token.Set(jwt.IssuerKey, "https://atomiclti.atomicjolt.xyz")
	token.Set(jwt.IssuedAtKey, time.Now().UTC().Unix())

	for _, inherited := range inheritedClaims() {
		token.Set(inherited.key, inherited.derive(idToken))
	}

	return &LaunchToken{
		Token: token,
	}
}

func ParseLaunchToken(r *http.Request, options ...jwt.ParseOption) (*LaunchToken, error) {
	rawToken, err := jwt.ParseRequest(r, options...)

	if err != nil {
		return nil, err
	}

	claims, err := rawToken.AsMap(context.Background())

	token := jwt.New()

	token.Set(jwt.SubjectKey, claims[jwt.SubjectKey])
	token.Set(jwt.AudienceKey, claims[jwt.AudienceKey])
	token.Set(jwt.IssuerKey, claims[jwt.IssuerKey])
	token.Set(jwt.IssuedAtKey, claims[jwt.IssuedAtKey])

	for _, inherited := range inheritedClaims() {
		token.Set(inherited.key, claims[inherited.key])
	}

	return &LaunchToken{
		Token: token,
	}, nil
}

func (t *LaunchToken) ContextId() string {
	contextId, ok := t.Get("contextId")

	if !ok {
		panic("Unable to infer context ID from launch token.")
	}

	return contextId.(string)
}
