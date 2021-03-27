package contracts

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/one-click-platform/deployer/contracts/deployer"
	"github.com/one-click-platform/system-contracts/generated"
)

func Tasks() []deployer.DeployFunc {
	return []deployer.DeployFunc{
		func(dep *deployer.Deployer) (common.Address, *types.Transaction, error) {
			addr, tx, _, err := generated.DeployAuction(dep.TransactOpts(), dep.Client)
			return addr, tx, err
		},
		func(dep *deployer.Deployer) (common.Address, *types.Transaction, error) {
			addr, tx, _, err := generated.DeployWETH(dep.TransactOpts(), dep.Client, "Wrapped ETH", "WETH")
			return addr, tx, err
		},
	}
}
