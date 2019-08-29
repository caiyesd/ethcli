package main

import (
	"github.com/caiyesd/ethcli/account"
	"github.com/caiyesd/ethcli/chain"
	"github.com/caiyesd/ethcli/crypto"
	"github.com/caiyesd/ethcli/erc20"
	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Name = "ethcli"
	app.Version = "v0.0.0"
	app.Usage = "an implementation of geth client"
	app.Description = "the client of geth"
	app.Flags = []cli.Flag{}
	app.Commands = []cli.Command{
		account.AccountCmd,
		chain.ChainCmd,
		crypto.CryptoCmd,
		erc20.Erc20Cmd,
	}
	app.RunAndExitOnError()
}
