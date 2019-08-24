package main

import (
	"github.com/caiyesd/ethcli/chain"
	"github.com/caiyesd/ethcli/crypto"
	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Name = "geth-cli"
	app.Version = "v1.0.0"
	app.Usage = "an implementation of geth client"
	app.Description = "the client of geth"
	app.Flags = []cli.Flag{}
	app.Commands = []cli.Command{
		chain.ChainCmd,
		crypto.CryptoCmd,
	}
	app.RunAndExitOnError()
}
