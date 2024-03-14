package middleware

import (
	"filmlibr/internal/sending"
	"filmlibr/internal/session"
	"net/http"
)

var MethodsAdminOnly = map[string]bool{
	http.MethodPost:   true,
	http.MethodPatch:  true,
	http.MethodPut:    true,
	http.MethodDelete: true,
}

func AccessControl(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess, err := session.SessionFromContext(r.Context())
		if err != nil {
			// internal error cause server must set session in auth middleware
			sending.JSONError(w, ErrInternal, http.StatusInternalServerError)
			return
		}
		if !sess.User.IsAdmin && MethodsAdminOnly[r.Method] {
			sending.JSONError(w, ErrNoAccess, http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
