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
		noAuth := []string{"/v1/user/new", "/v1/user/login", "/v1/product/all"}
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
		claim.Parse(bearerToken)
		//all passed
		//ctx := context.WithValue(r.Context(),"claim",claim)
		//r.WithContext(ctx)
		context.Set(r, Xclaim, claim)
		next.ServeHTTP(w, r)
	})
}

// AllowAccess verification
func AllowAccess(r *http.Request, AllowedRoles []models.RoleType) error {
	claim := context.Get(r, Xclaim).(*models.Token)
	roles := claim.Roles
	if err := roles.IsAllowed(AllowedRoles); err != nil {
		return err
	}
}
