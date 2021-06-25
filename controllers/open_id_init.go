package controllers

import (
	"github.com/atomicjolt/atomiclti/lib"
	"github.com/atomicjolt/atomiclti/middleware"
	"net/http"
	"net/url"
)

func NewOpenIDInitHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		controllerResources := middleware.GetResources(r.Context())
		var iss, clientId, targetLinkUri, loginHint, ltiMessageHint string

		switch r.Method {
		case "GET":
			query := r.URL.Query()

			iss = query.Get("iss")
			clientId = query.Get("client_id")
			targetLinkUri = query.Get("target_link_uri")
			loginHint = query.Get("login_hint")
			ltiMessageHint = query.Get("lti_message_hint")
		case "POST":
			if err := r.ParseForm(); err != nil {
				panic(err)
			}

			iss = r.FormValue("iss")
			clientId = r.FormValue("client_id")
			targetLinkUri = r.FormValue("target_link_uri")
			loginHint = r.FormValue("login_hint")
			ltiMessageHint = r.FormValue("lti_message_hint")
		default:
			panic("Open ID Connect Controller cannot handle this type of request.")
		}

		ltiInstall, err := controllerResources.Repo.LtiInstall.From(iss, clientId)

		if err != nil {
			panic(err)
		}

		authNonce, err := lib.RandomHex(64)

		if err != nil {
			panic(err)
		}

		oidcUrl := ltiInstall.OidcUrl
		oidcQuery := url.Values{}

		state, err := controllerResources.Repo.OpenIdState.IssueToken()

		if err != nil {
			panic(err)
		}

		oidcQuery.Add("response_type", "id_token")
		oidcQuery.Add("redirect_uri", targetLinkUri)
		oidcQuery.Add("response_mode", "form_post")
		oidcQuery.Add("client_id", clientId)
		oidcQuery.Add("scope", "openid")
		oidcQuery.Add("state", state)
		oidcQuery.Set("login_hint", loginHint)
		oidcQuery.Set("prompt", "none")
		oidcQuery.Set("lti_message_hint", ltiMessageHint)
		oidcQuery.Set("nonce", authNonce)

		http.Redirect(w, r, oidcUrl+"?"+oidcQuery.Encode(), http.StatusSeeOther)
	})
}
