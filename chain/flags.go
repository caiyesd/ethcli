package chain

import (
	cli "gopkg.in/urfave/cli.v1"
)

var (
	FlagRpcAddr = cli.StringFlag{
		Name:  "rpc-addr",
		Usage: "geth server rpc address: format: http://..., https://..., ws://...",
		Value: "ws://127.0.0.1:18546",
	}
	FlagNoPending = cli.BoolTFlag{
		Name:  "no-pending",
		Usage: "get no pending state",
	}
)
