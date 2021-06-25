package middleware

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func NewDebugger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dump, err := httputil.DumpRequest(r, true)

		if err != nil {
			panic(err)
		}

		fmt.Printf("\n%s", dump)

		next.ServeHTTP(w, r)
	})
}
