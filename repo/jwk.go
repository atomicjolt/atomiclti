package repo

import (
	"github.com/atomicjolt/atomiclti/model"
	"github.com/lestrrat-go/jwx/jwk"
)

type JwkRepo struct {
	*BaseRepo
}

func (r *JwkRepo) First() (model.Jwk, error) {
	var jwk model.Jwk
	_, err := r.DB.QueryOne(&jwk, "SELECT * FROM jwks ORDER BY created_at DESC LIMIT 1")

	return jwk, err
}

func (r *JwkRepo) All() ([]model.Jwk, error) {
	var jwks []model.Jwk
	err := r.DB.Model(&jwks).Select()

	return jwks, err
}

func (r *JwkRepo) Find(id int64) (*model.Jwk, error) {
	jwk := &model.Jwk{ID: id}
	err := r.DB.Model(jwk).WherePK().Select()

	return jwk, err
}

func (r *JwkRepo) PrivateKey() (jwk.Key, error) {
	jwk, err := r.First()

	if err != nil {
		return nil, err
	}

	key, err := jwk.ToKey()

	if err != nil {
		return nil, err
	}

	return key, nil
}

func (r *JwkRepo) PublicJwkSet() (jwk.Set, error) {
	jwks, err := r.All()

	if err != nil {
		return nil, err
	}

	set := jwk.NewSet()

	for _, j := range jwks {
		key, err := j.ToKey()

		if err != nil {
			panic(err)
		}

		set.Add(key)
	}

	return jwk.PublicSetOf(set)
}
