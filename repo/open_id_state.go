package repo

import (
	"errors"
	"fmt"
	"time"

	"github.com/atomicjolt/atomiclti/config"
	"github.com/atomicjolt/atomiclti/lib"
	"github.com/atomicjolt/atomiclti/model"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
)

type OpenIdStateRepo struct {
	*BaseRepo
}

func (r *OpenIdStateRepo) All() ([]model.OpenIdState, error) {
	var jwks []model.OpenIdState
	err := r.DB.Model(&jwks).Select()

	return jwks, err
}

func (r *OpenIdStateRepo) Find(id int64) (*model.OpenIdState, error) {
	jwk := &model.OpenIdState{ID: id}
	err := r.DB.Model(jwk).WherePK().Select()

	return jwk, err
}

func (r *OpenIdStateRepo) NewState() (*model.OpenIdState, error) {
	nonce, err := lib.RandomHex(64)

	if err != nil {
		return nil, err
	}

	openIdState := &model.OpenIdState{
		Nonce: nonce,
	}

	err = r.Insert(openIdState)

	if err != nil {
		return nil, err
	}

	return openIdState, nil
}

func (r *OpenIdStateRepo) IssueToken() (string, error) {
	openIdState, err := r.NewState()

	if err != nil {
		return "", err
	}

	serverConfig := config.GetServerConfig()
	authClientId := serverConfig.AuthClientId
	authClientSecret := serverConfig.AuthClientSecret
	authKey := jwk.NewSymmetricKey()

	err = authKey.FromRaw(authClientSecret)

	if err != nil {
		return "", err
	}

	authKey.Set(jwk.AlgorithmKey, jwa.HS512)
	authKey.Set(jwk.KeyIDKey, "atomiclti_auth0")

	token := jwt.New()
	token.Set(jwt.AudienceKey, authClientId)
	token.Set(jwt.IssuedAtKey, time.Now().UTC().Unix())
	token.Set(jwt.ExpirationKey, time.Now().Add(time.Hour*24).UTC().Unix())
	token.Set("nonce", openIdState.Nonce)

	signed, err := jwt.Sign(token, jwa.HS512, authKey)

	if err != nil {
		return "", err
	}

	return string(signed), nil
}

func (r *OpenIdStateRepo) ValidateStateOf(nonce string) (bool, error) {
	openIdState := new(model.OpenIdState)

	err := r.DB.Model(openIdState).Where("nonce = ?", nonce).Select()

	if err != nil {
		return false, errors.New("nonce from OpenID response is unknown")
	}

	res, err := r.DB.Model(openIdState).Where("nonce = ?", nonce).Delete()

	if err != nil {
		return false, err
	}

	if res.RowsAffected() != 1 {
		// This may look like a test, but if the nonce isn't deleted we open the user
		// to a replay attack.
		return false, fmt.Errorf("expected to delete one OpenIdState, deleted %d instead", res.RowsAffected())
	}

	return true, nil
}
