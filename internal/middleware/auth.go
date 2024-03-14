package middleware

import (
	"filmlibr/internal/sending"
	"filmlibr/internal/session"
	"net/http"
)

func Auth(sm session.SessionManager, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess, err := sm.Check(r)
		if err != nil {
			// TODO error placement
			sending.JSONError(w, session.ErrNoAuth, http.StatusUnauthorized)
			return
		}
		ctx := session.ContextWithSession(r.Context(), sess)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
