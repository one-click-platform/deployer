package service

import (
	"github.com/one-click-platform/deployer/internal/deploy"
	"github.com/one-click-platform/deployer/resources"
	"gitlab.com/distributed_lab/logan/v3"
)

type Processor struct {
	storage   map[string]resources.EnvConfig
	tasks     chan string
	log       *logan.Entry
	githubKey string
}

func (p *Processor) Run() {
	p.log.Info("Async processor started")
	task := <-p.tasks

	go p.CreateEnv(task)
}

func (p *Processor) CreateEnv(name string) {
	p.log.Infof("Creating env: %s", name)
	envConfig, err := deploy.Deploy(name, p.log, p.githubKey)
	if err != nil {
		p.log.WithError(err).Error("failed to deploy node")
		return
	}

	p.storage[name] = resources.EnvConfig{
		Status:       "created",
		Passphrase:   envConfig.Passphrase,
		SshKey:       envConfig.SSHKey,
		ValidatorKey: string(envConfig.ValidatorKey),
	}
}
