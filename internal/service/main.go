package service

import (
	"gitlab.com/distributed_lab/kit/pgdb"
	"net"
	"net/http"

	"github.com/one-click-platform/deployer/resources"

	"github.com/one-click-platform/deployer/internal/config"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type service struct {
	log      *logan.Entry
	copus    types.Copus
	listener net.Listener

	storage map[string]resources.EnvConfig
	tasks   chan string
	cfg     config.Config
	db      *pgdb.DB
}

func (s *service) run() error {
	proc := Processor{
		log:       s.log,
		tasks:     s.tasks,
		storage:   s.storage,
		githubKey: s.cfg.GithubKey(),
	}
	go proc.Run()

	r := s.router()

	if err := s.copus.RegisterChi(r); err != nil {
		return errors.Wrap(err, "cop failed")
	}

	return http.Serve(s.listener, r)
}

func newService(cfg config.Config) *service {
	return &service{
		log:      cfg.Log(),
		copus:    cfg.Copus(),
		listener: cfg.Listener(),
		db:       cfg.DB(),
		storage:  make(map[string]resources.EnvConfig),
		tasks:    make(chan string),
		cfg:      cfg,
	}
}

func Run(cfg config.Config) {
	cfg.SetupAWS()
	if err := newService(cfg).run(); err != nil {
		panic(err)
	}
}
