package middleware

import (
	"context"
	"github.com/vadym-98/playground/slide13/domain"
	"net/http"
)

type Func func(handler http.Handler) http.Handler

func RequireAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u := domain.GetAuthUser()
		if u == nil {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}

		c := context.WithValue(r.Context(), domain.AuthKey, u)
		req := r.WithContext(c)

		//Authentication was successful, send the request to the next handler
		next.ServeHTTP(w, req)
	})
}
