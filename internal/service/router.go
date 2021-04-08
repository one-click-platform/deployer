package service

import (
	"github.com/go-chi/chi"
	"github.com/one-click-platform/deployer/internal/data/pg"
	"github.com/one-click-platform/deployer/internal/service/handlers"
	"github.com/one-click-platform/deployer/internal/service/middlewares"
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
			handlers.CtxAccountsQ(pg.NewAccountsQ(s.db)),
			handlers.CtxJwtHandler(s.cfg.ReadConfig()),
		),
	)

	r.Route("/auth", func(r chi.Router) {
		r.Post("/sign-up", handlers.SignUp)
		r.Post("/sign-in", handlers.SignIn)
	})
	r.Route("/envs", func(r chi.Router) {
		r.Use(middlewares.AuthMiddleware)
		r.Post("/", handlers.CreateNode)
		r.Get("/{name}", handlers.GetEnv)
	})

	return r
}
