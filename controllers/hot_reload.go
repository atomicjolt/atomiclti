package controllers

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func NewHotReloadProxy(rawUrl string) http.Handler {
	filesUrl, err := url.Parse(rawUrl)

	if err != nil {
		log.Fatal(err)
	}

	return httputil.NewSingleHostReverseProxy(filesUrl)
}
