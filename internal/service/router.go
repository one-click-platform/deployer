package service

import (
    "github.com/go-chi/chi"
    "gitlab.com/distributed_lab/ape"
    "github.com/one-click-platform/deployer/internal/service/handlers"
)

func (s *service) router() chi.Router {
    r := chi.NewRouter()

    r.Use(
        ape.RecoverMiddleware(s.log),
        ape.LoganMiddleware(s.log),
        ape.CtxMiddleware(
            handlers.CtxLog(s.log),
        ),
    )

    r.Route("/envs", func(r chi.Router) {
        r.Post("/", handlers.CreateNode)
    })

    return r
}
