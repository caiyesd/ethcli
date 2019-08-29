package chain

import (
	"github.com/caiyesd/ethcli/utils"
	cli "gopkg.in/urfave/cli.v1"
)

var (
	FlagRpcAddr = cli.StringFlag{
		Name:  "rpc",
		Usage: "geth server rpc address: format: http://..., https://..., ws://...",
		Value: "http://127.0.0.1:8545",
	}
	FlagNoPending = cli.BoolTFlag{
		Name:  "no-pending",
		Usage: "get no pending state",
	}
)

func ParseFlagRpcAddr(c *cli.Context) string {
	return c.Parent().String(FlagRpcAddr.Name)
}

func ParseFlagNoPending(c *cli.Context) bool {
	return c.BoolT(FlagNoPending.Name)
}

var ParseBigInt = utils.ParseBigInt
var ParseUint64 = utils.ParseUint64
