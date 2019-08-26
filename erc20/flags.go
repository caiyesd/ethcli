package erc20

import (
	"os"
	"path/filepath"

	cli "gopkg.in/urfave/cli.v1"
)

var (
	FlagKeystore = cli.StringFlag{
		Name:  "keystore",
		Usage: "keystore directory",
		Value: DefaultKeystore(),
	}
	FlagRpcAddr = cli.StringFlag{
		Name:  "rpc-addr",
		Usage: "geth server rpc address: format: http://..., https://..., ws://...",
		Value: "http://127.0.0.1:8545",
	}
	FlagErc20Addr = cli.StringFlag{
		Name:  "erc20-addr",
		Usage: "erc20 token address",
		Value: "",
	}
)

func DefaultHome() string {
	home, err := os.UserHomeDir()
	if err != nil {
		home = ""
	}
	return home
}

func DefaultKeystore() string {
	return filepath.Join(DefaultHome(), ".ethereum", "keystore")
}
