package middlewares

import (
	"github.com/one-click-platform/deployer/internal/service/auth"
	"github.com/one-click-platform/deployer/internal/service/handlers"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
	"strings"
)

func AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := auth.VerifyToken(ParseHeader(r))

		if err != nil {
			handlers.Log(r).WithError(err).Info("Failed to authenticate")
			// TODO change to unauthorized
			ape.RenderErr(w, problems.Conflict())
			return
		}

		h.ServeHTTP(w, r)
	})
}

func ParseHeader(r *http.Request) string {
	authorizationHeader := r.Header.Get("Authorization")
	header := strings.Split(authorizationHeader, " ")
	if len(header) == 2 {
		return header[1]
	}
	return ""
}
