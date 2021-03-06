package middlewares

import (
	"github.com/one-click-platform/deployer/internal/service/handlers"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
	"strings"
)

func AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, tokenClaims, err := handlers.JwtHandler(r).VerifyToken(ParseAuthorizationHeader(r))

		if err != nil {
			handlers.Log(r).WithError(err).Info("failed to authenticate")
			// TODO change to unauthorized
			ape.RenderErr(w, problems.Conflict())
			return
		}

		// TODO add struct for payload(tokenClaims) and pass it into context
		ctx := handlers.CtxJWTPayload(tokenClaims["sub"].(string))(r.Context())
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func ParseAuthorizationHeader(r *http.Request) string {
	authorizationHeader := r.Header.Get("Authorization")
	header := strings.Split(authorizationHeader, " ")
	if len(header) == 2 {
		return header[1]
	}
	return ""
}
