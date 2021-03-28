package deployer

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type DeployFunc func(dep *Deployer) (common.Address, *types.Transaction, error)

// Deployer of native contracts.
type Deployer struct {
	Log    *logan.Entry
	Client *ethclient.Client
	Opts   *bind.TransactOpts
}

// New Deployer.
func New(ctx context.Context, ethClient *ethclient.Client, keyStore *keystore.KeyStore, account accounts.Account,
	log *logan.Entry) (*Deployer, error) {

	client := ethClient
	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get chain id")
	}

	opts, err := bind.NewKeyStoreTransactorWithChainID(keyStore, account, chainID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create transact opts")
	}

	return &Deployer{
		Log:    log,
		Client: client,
		Opts:   opts,
	}, nil
}

// TransactOpts
func (d *Deployer) TransactOpts() *bind.TransactOpts {
	return d.Opts
}

// Run deployment tasks.
func (d *Deployer) Run(ctx context.Context, tasks []DeployFunc) ([]common.Address, error) {
	addresses := make([]common.Address, 0, len(tasks))
	for _, tsk := range tasks {
		addr, tx, err := tsk(d)
		addresses = append(addresses, addr)
		if err != nil {
			return nil, errors.Wrap(err, "failed to send deploy tx")
		}

		d.Log.WithField("address", addr.String()).Info("tx sent")

		if _, err := bind.WaitDeployed(ctx, d.Client, tx); err != nil {
			return nil, errors.Wrap(err, "deploy tx has failed")
		}
	}

	return addresses, nil
}
