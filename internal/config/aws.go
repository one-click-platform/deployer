package config

import (
	"os"

	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

type Aws interface {
	SetupAWS()
}

func NewAws(getter kv.Getter) Aws {
	return &aws{
		getter: getter,
	}
}

type aws struct {
	getter kv.Getter
	once   comfig.Once
}

func (d *aws) SetupAWS() {
	d.once.Do(func() interface{} {
		config := struct {
			KeyID     string `figure:"key_id"`
			SecretKey string `figure:"secret_key"`
		}{}

		raw, err := d.getter.GetStringMap("aws")
		if err != nil {
			raw = make(map[string]interface{})
		}
		err = figure.Out(&config).From(raw).Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out"))
		}

		if err := os.Setenv("AWS_ACCESS_KEY_ID", config.KeyID); err != nil {
			panic(err)
		}
		if err := os.Setenv("AWS_SECRET_ACCESS_KEY", config.SecretKey); err != nil {
			panic(err)
		}

		return nil
	})
}
