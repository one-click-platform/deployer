package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/copus"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type Config interface {
	comfig.Logger
	types.Copuser
	comfig.Listenerer
	pgdb.Databaser
	Aws
}

type config struct {
	comfig.Logger
	types.Copuser
	comfig.Listenerer
	pgdb.Databaser
	Aws
	getter kv.Getter
}

func New(getter kv.Getter) Config {
	return &config{
		getter:     getter,
		Copuser:    copus.NewCopuser(getter),
		Listenerer: comfig.NewListenerer(getter),
		Logger:     comfig.NewLogger(getter, comfig.LoggerOpts{}),
		Databaser:  pgdb.NewDatabaser(getter),
		Aws:        NewAws(getter),
	}
}
