package erc20

import (
	"github.com/caiyesd/ethcli/chain"
	"github.com/caiyesd/ethcli/crypto"
	"github.com/caiyesd/ethcli/utils"
	cli "gopkg.in/urfave/cli.v1"
)

var (
	FlagKeystore = crypto.FlagKeystore
	FlagPassword = crypto.FlagPassword
	FlagRpcAddr  = chain.FlagRpcAddr

	FlagErc20Addr = cli.StringFlag{
		Name:  "erc20",
		Usage: "erc20 token address",
		Value: "",
	}
)

var ParseFlagKeystore = crypto.ParseFlagKeystore
var ParseFlagPassword = crypto.ParseFlagPassword
var ReadPassphrase = crypto.ReadPassphrase
var ParseBigInt = utils.ParseBigInt
var ParseUint64 = utils.ParseUint64
