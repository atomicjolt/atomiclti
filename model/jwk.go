package model

import (
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/spacemonkeygo/openssl"
	"strconv"
	"time"
)

type Jwk struct {
	ID  int64
	Kid string `pg:",notnull"`
	Pem string `pg:",notnull"`

	ApplicationID int64 `pg:"on_delete:CASCADE"`

	Timestamps
}

func NewJwk() *Jwk {
	privateKey, err := openssl.GenerateRSAKey(4096)

	if err != nil {
		panic(err)
	}

	pem, err := privateKey.MarshalPKCS1PrivateKeyPEM()

	if err != nil {
		panic(err)
	}

	return &Jwk{
		Kid: strconv.FormatInt(time.Now().UTC().Unix(), 10),
		Pem: string(pem),
	}
}

func (j *Jwk) ToKey() (jwk.Key, error) {
	key, err := jwk.ParseKey([]byte(j.Pem), jwk.WithPEM(true))

	if err != nil {
		return nil, err
	}

	err = key.Set(jwk.KeyIDKey, j.Kid)
	err = key.Set(jwk.KeyUsageKey, "sig")
	err = key.Set(jwk.AlgorithmKey, "RS256")

	return key, nil
}
