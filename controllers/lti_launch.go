package controllers

import (
	"fmt"
	"github.com/atomicjolt/atomiclti/config"
	"github.com/atomicjolt/atomiclti/middleware"
	"github.com/atomicjolt/atomiclti/webpack"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
	"net/http"
	"path"
	"text/template"
)

type ViewState struct {
	Manifest          *webpack.Manifest
	LaunchTokenSigned string
}

func index(w http.ResponseWriter, r *http.Request) error {
	controllerResources := middleware.GetResources(r.Context())
	view, err := template.ParseFiles(path.Join("views", "index.html"))

	if err != nil {
		return err
	}

	var manifest *webpack.Manifest

	if config.DetermineEnv() == "development" {
		manifest, err = webpack.NewFromDevServer("http://127.0.0.1:3000/asset-manifest.json")
	} else {
		manifest, err = webpack.NewFromBuildPath("client/build")
	}

	if err != nil {
		return err
	}

	/**
	 * Use of the private key is restricted to this scope
	 */
	var launchTokenSigned []byte
	{
		pKey, err := controllerResources.Repo.Jwk.PrivateKey()

		if err != nil {
			return err
		}

		launchToken := middleware.GetLaunchToken(r.Context())
		launchTokenSigned, err = jwt.Sign(launchToken.Token, jwa.RS256, pKey)

		if err != nil {
			return err
		}
	}

	state := &ViewState{
		Manifest:          manifest,
		LaunchTokenSigned: string(launchTokenSigned),
	}

	return view.Execute(w, state)
}

func NewLtiLaunchHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		controllerResources := middleware.GetResources(r.Context())
		payload := middleware.GetOidcState(r.Context())
		nonce, ok := payload["nonce"]

		if !ok {
			panic(fmt.Errorf("expected nonce in OIDC claims"))
		}

		valid, err := controllerResources.Repo.OpenIdState.ValidateStateOf(nonce.(string))

		if err != nil {
			panic(err)
		} else if !valid {
			panic(fmt.Errorf("OIDC token could not be validated"))
		}

		if err := index(w, r); err != nil {
			panic(err)
		}
	})
}
