package deploy

import (
	"context"
	"fmt"
	"io/ioutil"
	"os/exec"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"

	"github.com/one-click-platform/deployer/contracts"

	"github.com/one-click-platform/deployer/contracts/deployer"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/logan/v3"
)

func Deploy(name string, log *logan.Entry) (EnvConfig, error) {
	config, err := DeployEC2(name, log)
	if err != nil {
		return EnvConfig{}, errors.Wrap(err, "failed to create ec2 instance")
	}
	if err := DeploySmartcontracts(config, log); err != nil {
		return EnvConfig{}, errors.Wrap(err, "failed to deploy smartcontracts")
	}

	return EnvConfig{}, nil
}

func DeployEC2(name string, log *logan.Entry) (NodeConfig, error) {
	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("cd /scripts && sh aws/create_instance.sh 1 %s Etherium_vpc Etherium_sub", name))
	log.Info(cmd.String())
	if err := cmd.Run(); err != nil {
		return NodeConfig{}, errors.Wrap(err, "failed to execute create instance script")
	}
	stdout, _ := cmd.StdoutPipe()
	b, _ := ioutil.ReadAll(stdout)
	log.Info(string(b))

	return NodeConfig{}, nil
}

func DeploySmartcontracts(config NodeConfig, log *logan.Entry) error {
	client, err := ethclient.Dial(config.Endpoint)
	if err != nil {
		return errors.Wrap(err, "failed to create connection to node")
	}

	ks := keystore.NewKeyStore(config.KeyStoreDir, keystore.StandardScryptN, keystore.StandardScryptP)
	acc := accounts.Account{Address: common.HexToAddress(config.Address)}
	if err := ks.Unlock(acc, config.Password); err != nil {
		return errors.Wrap(err, "failed to unlock contracts")
	}

	contractsDeployer, err := deployer.New(context.TODO(), client, ks, acc, log)
	if err != nil {
		return errors.Wrap(err, "failed to create contracts deployer")
	}
	if err := contractsDeployer.Run(context.TODO(), contracts.Tasks()); err != nil {
		return errors.Wrap(err, "failed to deploy contracts")
	}

	return nil
}
