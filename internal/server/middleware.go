package server

import (
	"net/http"
)

func SecureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Note: This is split across multiple lines for readability. You don't // need to do this in your own code.
		w.Header().Set(
			"Content-Security-Policy",
			"default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com")
		w.Header().Set(
			"Referrer-Policy",
			"origin-when-cross-origin")
		w.Header().Set(
			"X-Content-Type-Options",
			"nosniff")
		w.Header().Set(
			"X-Frame-Options",
			"deny")
		w.Header().Set(
			"X-XSS-Protection",
			"0")
		next.ServeHTTP(w, r)
	})
}

// ReportPanic is middleware for catching panics and reporting them.
func ReportPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				// TODO pham-k: report panic
			}
		}()

		next.ServeHTTP(w, r)
	})
}
