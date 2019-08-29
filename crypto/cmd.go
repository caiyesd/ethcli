package crypto

import (
	cli "gopkg.in/urfave/cli.v1"
)

var CryptoCmd = cli.Command{
	Name:  "crypto",
	Usage: "crypto operations",
	Flags: []cli.Flag{FlagKeystore, FlagPassword},
	Subcommands: []cli.Command{
		CalcHashCmd,
		SignHashCmd,
		BuildTxCmd,
		SignTxCmd,
		VerifyCmd,
		EcrecoverCmd,
		PackCmd,
	},
}

var CalcHashCmd = cli.Command{
	Name:      "hash",
	Usage:     "calculate keccak256 hash of data",
	UsageText: "hash <data>",
	Action:    CalcHashAction,
	Flags:     []cli.Flag{},
}

var SignHashCmd = cli.Command{
	Name:      "signhash",
	Usage:     "sign a hash",
	UsageText: "sign <account> <hash>",
	Action:    SignHashAction,
	Flags:     []cli.Flag{},
}

var BuildTxCmd = cli.Command{
	Name:      "tx",
	Usage:     "build a new tx",
	UsageText: "tx <nonce> <to> <value> <gasLimit> <gasPrice> [data]",
	Action:    BuildTxAction,
	Flags:     []cli.Flag{},
}

var SignTxCmd = cli.Command{
	Name:      "signtx",
	Usage:     "sign a transaction",
	UsageText: "sign <account> <tx>",
	Action:    SignTxAction,
	Flags:     []cli.Flag{},
}

var VerifyCmd = cli.Command{
	Name:      "verify",
	Usage:     "verify a signature",
	UsageText: "verify <pubkey> <hash> <sig>",
	Action:    VerifyAction,
	Flags:     []cli.Flag{},
}

var EcrecoverCmd = cli.Command{
	Name:      "ecrecover",
	Usage:     "recover address from a signature",
	UsageText: "ecrecover <hash> <sig>",
	Action:    EcrecoverAction,
	Flags:     []cli.Flag{},
}

var PackCmd = cli.Command{
	Name:      "pack",
	Usage:     "pack arguments",
	UsageText: "pack <type0> <arg0> <type1> <arg1> ...",
	Action:    PackAction,
	Flags:     []cli.Flag{},
}
