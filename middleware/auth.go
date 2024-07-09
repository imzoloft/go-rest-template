package middleware

import (
	"net/http"

	"github.com/imzoloft/go-rest-api/httputil"
	"github.com/imzoloft/go-rest-api/response"
)

func Auth(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Real Authentication here
		token := r.Header.Get("Authorization")
		if token == "" {
			httputil.WriteError(w, http.StatusUnauthorized, response.ErrNotAuthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}
