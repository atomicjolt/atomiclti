package lti

import (
	"github.com/lestrrat-go/jwx/jwt"
)

type IdToken struct {
	jwt.Token
}

func NewIdToken(idToken jwt.Token) IdToken {
	return IdToken{
		Token: idToken,
	}
}

func (t IdToken) ValueOf(key definition) map[string]interface{} {
	value, ok := t.Get(string(key))

	if !ok {
		panic("Value could not be found in ID token")
	}

	return value.(map[string]interface{})
}

func (t IdToken) ContextId() string {
	return t.ValueOf(Definitions.Claims.Context)["id"].(string)
}
