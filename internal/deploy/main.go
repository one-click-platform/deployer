package deploy

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"

	"github.com/one-click-platform/deployer/contracts"

	"github.com/one-click-platform/deployer/contracts/deployer"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/logan/v3"
)

func Deploy(name string, log *logan.Entry, githubKey string) (EnvConfig, error) {
	config, err := DeployNode(name, log)
	if err != nil {
		return EnvConfig{}, errors.Wrap(err, "failed to create ec2 instance")
	}

	ks := keystore.NewKeyStore(config.KeyStoreDir, keystore.StandardScryptN, keystore.StandardScryptP)
	acc := accounts.Account{Address: common.HexToAddress(config.Address)}
	if err := ks.Unlock(acc, config.Password); err != nil {
		return EnvConfig{}, errors.Wrap(err, "failed to unlock contracts")
	}
	passphrase := "asdf78sd7fy83h8348sd"
	keyJSON, err := ks.Export(acc, config.Password, passphrase)
	if err != nil {
		return EnvConfig{}, errors.Wrap(err, "failed to export json key")
	}

	addresses, err := DeploySmartcontracts(config, log)
	if err != nil {
		return EnvConfig{}, errors.Wrap(err, "failed to deploy smartcontracts")
	}
	err = DeployEnv(config, addresses, name, log, githubKey)
	if err != nil {
		return EnvConfig{}, errors.Wrap(err, "failed to deploy env")
	}

	log.Info("Deployment finished")

	return EnvConfig{
		SSHKey:       config.SshKey,
		ValidatorKey: keyJSON,
		Passphrase:   passphrase,
	}, nil
}

func DeployNode(name string, log *logan.Entry) (NodeConfig, error) {
	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("cd /scripts && sh aws/create_instance.sh 1 %s Etherium_vpc Etherium_sub", name))
	log.Info(cmd.String())
	if err := cmd.Run(); err != nil {
		return NodeConfig{}, errors.Wrap(err, "failed to execute create instance script")
	}
	log.Info("node deployed")

	config := NodeConfig{}

	config.Password = "qwerty"

	sshKey, err := ioutil.ReadFile(fmt.Sprintf("/scripts/keys/%s.pem", name))
	if err != nil {
		return config, errors.Wrap(err, "failed to read ssh key")
	}
	config.SshKey = string(sshKey)

	address, err := ioutil.ReadFile(fmt.Sprintf("/scripts/keys/%s/adr.txt", name))
	if err != nil {
		return config, errors.Wrap(err, "failed to read ssh key")
	}
	config.Address = strings.ReplaceAll(string(address), "\n", "")
	log.Info(config.Address)

	ip, err := ioutil.ReadFile(fmt.Sprintf("/scripts/keys/%s/dostup.txt", name))
	config.IP = strings.ReplaceAll(string(ip), "\n", "")
	if err != nil {
		return config, errors.Wrap(err, "failed to read ssh key")
	}
	config.Endpoint = fmt.Sprintf("http://%s:8545", config.IP)
	log.Info(config.IP)
	log.Info(config.Endpoint)

	config.KeyStoreDir = fmt.Sprintf("/scripts/keys/%s/keystore", name)

	return config, nil
}

func DeploySmartcontracts(config NodeConfig, log *logan.Entry) ([]common.Address, error) {
	log.Info("Deploying contracts")
	client, err := ethclient.Dial(config.Endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create connection to node")
	}

	ks := keystore.NewKeyStore(config.KeyStoreDir, keystore.StandardScryptN, keystore.StandardScryptP)
	acc := accounts.Account{Address: common.HexToAddress(config.Address)}
	if err := ks.Unlock(acc, config.Password); err != nil {
		return nil, errors.Wrap(err, "failed to unlock contracts")
	}

	contractsDeployer, err := deployer.New(context.TODO(), client, ks, acc, log)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create contracts deployer")
	}
	addresses, err := contractsDeployer.Run(context.TODO(), contracts.Tasks())
	if err != nil {
		return nil, errors.Wrap(err, "failed to deploy contracts")
	}

	return addresses, nil
}

func DeployEnv(config NodeConfig, addresses []common.Address, name string, log *logan.Entry, githubKey string) error {
	log.Info("Creating env.js file")
	envJs := fmt.Sprintf("document.ENV = {\nAUCTION_ADDRESS: '%s',\nTOKEN_ADDRESS: '%s',\nCURRENCY_ADDRESS: '%s'\n}",
		addresses[0].String(), addresses[2].String(), addresses[1].String())
	file, err := os.Create("/scripts/keys/env.js")
	if err != nil {
		return errors.Wrap(err, "failed to create env.js file")
	}
	_, err = file.WriteString(envJs)
	if err != nil {
		return errors.Wrap(err, "failed to write env.js file")
	}

	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("cd /scripts/keys && scp -i %s.pem env.js ubuntu@%s:/home/ubuntu/env.js", name, config.IP))
	log.Info(cmd.String())
	if err := cmd.Run(); err != nil {
		return errors.Wrap(err, "failed to execute upload env.js script")
	}

	cmd = exec.Command("/bin/sh", "-c", fmt.Sprintf("cd /scripts/keys && scp -i %s.pem start_front.sh ubuntu@%s:/home/ubuntu/start.sh", name, config.IP))
	log.Info(cmd.String())
	if err := cmd.Run(); err != nil {
		return errors.Wrap(err, "failed to execute upload start_front.sh script")
	}

	cmd = exec.Command("/bin/sh", "-c", fmt.Sprintf("cd /scripts/keys && ssh -i %s.pem  ubuntu@%s sh start.sh %s", name, config.IP, githubKey))
	log.Info(cmd.String())
	if err := cmd.Run(); err != nil {
		log.WithError(err).Warn("failed to execute start front script")
	}

	return nil
}
