package erc20

import (
	"github.com/caiyesd/ethcli/chain"
	"github.com/caiyesd/ethcli/crypto"
	cli "gopkg.in/urfave/cli.v1"
)

var (
	FlagKeystore    = crypto.FlagKeystore
	FlagRpcAddr     = chain.FlagRpcAddr
	DefaultKeystore = crypto.DefaultKeystore

	FlagErc20Addr = cli.StringFlag{
		Name:  "erc20",
		Usage: "erc20 token address",
		Value: "",
	}
)
