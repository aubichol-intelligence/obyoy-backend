package middleware

import (
	"context"
	"net/http"

	"obyoy-backend/api/routeutils"
	"obyoy-backend/user"

	"github.com/sirupsen/logrus"
)

// Auth holds the middleware that verifies the session
type AdminAuth struct {
	userSessionVerifier user.SessionVerifier
}

// Middleware implments a middleware that checkes the user session for authorization
func (a *AdminAuth) Middleware(h http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		session, err := a.userSessionVerifier.VerifySession(token)
		if err != nil {
			logrus.Error("Unable to verify user token, error :", err)
			routeutils.ServeError(w, err)
			return
		}

		if session == nil {
			routeutils.ServeResponse(
				w,
				http.StatusForbidden,
				map[string]interface{}{
					"message": "Invalid token",
					"ok":      false,
				},
			)
			return
		}

		if session.UserType != "admin" {
			routeutils.ServeResponse(
				w,
				http.StatusForbidden,
				map[string]interface{}{
					"message": "Access deined",
					"ok":      false,
				},
			)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", session.UserID)
		h.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(f)
}

// NewAuthMiddleware returns an AuthMiddleware
func NewAdminAuthMiddleware(userSessionVerifier user.SessionVerifier) *AdminAuth {
	return &AdminAuth{userSessionVerifier}
}

// AuthMiddlewareURL stores an user session verifier
type AdminAuthMiddlewareURL struct {
	userSessionVerifier user.SessionVerifier
}
