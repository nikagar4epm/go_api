package middleware

import (
	c "context"
	"errors"
	"net/http"
	"strings"

	"github.com/nikagar4epm/go_api/api"
	"github.com/nikagar4epm/go_api/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("Invalid username or token.")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")

		if username == "" {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		var database *tools.DatabaseInterface
		var err error
		database, err = tools.NewDatabase()

		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)
		var token string = r.Header.Get("Authorization")

		if !isValidToken(loginDetails, token) {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		ctx := c.WithValue(r.Context(), "loginDetails", loginDetails)
		r = r.WithContext(ctx)
		// TO READ THIS DO:
		// r.Context().Value("loginDetails").(*tools.LoginDetails)

		next.ServeHTTP(w, r)
	})
}

func isValidToken(loginDetails *tools.LoginDetails, token string) bool {
	if loginDetails == nil {
		return false
	}

	if token == "" {
		return false
	}

	return strings.TrimPrefix(token, "Bearer ") == loginDetails.AuthToken
}
