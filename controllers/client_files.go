package controllers

import "net/http"

func NewClientFilesHandler() http.Handler {
	return http.FileServer(http.Dir("client/build"))
}
