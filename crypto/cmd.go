package crypto

import (
	cli "gopkg.in/urfave/cli.v1"
)

var CryptoCmd = cli.Command{
	Name:  "crypto",
	Usage: "manage and use keystore",
	Flags: []cli.Flag{FlagKeystore},
	Subcommands: []cli.Command{
		NewAccountCmd,
		ListAccountCmd,
		UpdateAccountCmd,
		DeleteAccountCmd,
		NewHashCmd,
		SignHashCmd,
		NewTxCmd,
		SignTxCmd,
		VerifyCmd,
		EcrecoverCmd,
	},
}

var NewAccountCmd = cli.Command{
	Name:      "new",
	Usage:     "generate a new account",
	UsageText: "new",
	Action:    NewAccountAction,
	Flags:     []cli.Flag{},
}

var ListAccountCmd = cli.Command{
	Name:      "list",
	Usage:     "list accounts",
	UsageText: "list",
	Action:    ListAccountAction,
	Flags:     []cli.Flag{},
}

var UpdateAccountCmd = cli.Command{
	Name:      "update",
	Usage:     "update account",
	UsageText: "update <address>",
	Action:    UpdateAccountAction,
	Flags:     []cli.Flag{},
}

var DeleteAccountCmd = cli.Command{
	Name:      "delete",
	Usage:     "delete account",
	UsageText: "update <address>",
	Action:    DeleteAccountAction,
	Flags:     []cli.Flag{},
}

var NewHashCmd = cli.Command{
	Name:      "hash",
	Usage:     "calculate keccak256 hash of the data",
	UsageText: "hash <data>",
	Action:    NewHashAction,
	Flags:     []cli.Flag{},
}

var SignHashCmd = cli.Command{
	Name:      "signhash",
	Usage:     "sign hash",
	UsageText: "sign <address> <hash>",
	Action:    SignHashAction,
	Flags:     []cli.Flag{},
}

var NewTxCmd = cli.Command{
	Name:      "tx",
	Usage:     "create a new tx",
	UsageText: "tx <nonce> <to> <value> <gasLimit> <gasPrice> [data]",
	Action:    NewTxAction,
	Flags:     []cli.Flag{},
}

var SignTxCmd = cli.Command{
	Name:      "signtx",
	Usage:     "sign transaction",
	UsageText: "sign <address> <tx>",
	Action:    SignTxAction,
	Flags:     []cli.Flag{},
}

var VerifyCmd = cli.Command{
	Name:      "verify",
	Usage:     "verify signature",
	UsageText: "verify <pubkey> <hash> <sig>",
	Action:    VerifyAction,
	Flags:     []cli.Flag{},
}

var EcrecoverCmd = cli.Command{
	Name:      "ecrecover",
	Usage:     "recover address from signature",
	UsageText: "ecrecover <hash> <sig>",
	Action:    EcrecoverAction,
	Flags:     []cli.Flag{},
}
