package config

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

type Aws interface {
	SetupAWS()
	GithubKey() string
}

func NewAws(getter kv.Getter) Aws {
	return &aws{
		getter: getter,
	}
}

type aws struct {
	getter kv.Getter
	once   comfig.Once

	githubKey string
}

func (d *aws) SetupAWS() {
	d.loadConfig()
}

func (d *aws) GithubKey() string {
	d.loadConfig()
	return d.githubKey
}

func (d *aws) loadConfig() {
	d.once.Do(func() interface{} {
		config := struct {
			KeyID     string `figure:"key_id"`
			SecretKey string `figure:"secret_key"`
			GithubKey string `figure:"github_key"`
		}{}

		raw, err := d.getter.GetStringMap("aws")
		if err != nil {
			raw = make(map[string]interface{})
		}
		err = figure.Out(&config).From(raw).Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out"))
		}

		file, err := os.Create("/root/.aws/credentials")
		if err != nil {
			panic(err)
		}
		_, err = file.WriteString(fmt.Sprintf("[hackaton]\naws_access_key_id = %s\naws_secret_access_key = %s",
			config.KeyID, config.SecretKey))
		if err != nil {
			panic(err)
		}

		d.githubKey = config.GithubKey

		return nil
	})
}
