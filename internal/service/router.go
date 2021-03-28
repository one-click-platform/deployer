package service

import (
	"github.com/go-chi/chi"
	"github.com/one-click-platform/deployer/internal/service/handlers"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxGithubKey(s.cfg.GithubKey()),
			handlers.CtxStorage(s.storage),
			handlers.CtxTasks(s.tasks),
		),
	)

	r.Route("/envs", func(r chi.Router) {
		r.Post("/", handlers.CreateNode)
	})

	return r
}
