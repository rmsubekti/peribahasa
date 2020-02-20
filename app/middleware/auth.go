package middleware

import (
	"net/http"
	"peribahasa/app/models"

	"github.com/gorilla/context"
)

// Xclaim context
var Xclaim = &models.Token{}

// JwtAuthentication middleware
var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		noAuth := []string{"/api/register", "/api/login", "/"}
		requestPath := r.URL.Path

		//skip authentication for request path that in whitelist
		for _, value := range noAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		bearerToken := r.Header.Get("Authorization")
		claim := &models.Token{}
		if err := claim.Parse(bearerToken); err != nil && requestPath == "/api" {
			next.ServeHTTP(w, r)
			return
		}

		context.Set(r, Xclaim, claim)
		next.ServeHTTP(w, r)
	})
}

// AllowAccess verification
func AllowAccess(r *http.Request, AllowedRoles models.RoleTypes) error {
	claim := context.Get(r, Xclaim).(*models.Token)
	var roles models.Roles = claim.Roles
	if err := roles.IsAllowed(AllowedRoles); err != nil {
		return err
	}
	return nil
}
