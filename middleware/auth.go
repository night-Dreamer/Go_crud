package middleware

import (
	"net/http"
)

type AuthMiddleware struct {
	Next http.Handler
}

//
func (m *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if m.Next == nil {
		m.Next = http.DefaultServeMux
	}
	m.Next.ServeHTTP(w, r)
	// auth := r.Header.Get("identity")
	// if auth == "admin" {
	// 	m.Next.ServeHTTP(w, r)
	// } else {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// }
}
