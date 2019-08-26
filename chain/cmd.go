package chain

import (
	cli "gopkg.in/urfave/cli.v1"
)

var ChainCmd = cli.Command{
	Name:  "chain",
	Usage: "communicate with chain",
	Flags: []cli.Flag{FlagRpcAddr},
	Subcommands: []cli.Command{
		HeaderCmd,
		BlockCmd,
		TransactionCmd,
		ReceiptCmd,
		BalanceCmd,
		NonceCmd,
		CodeCmd,
		StorageCmd,
		SendCmd,
		CallCmd,
		GasCmd,
		PriceCmd,
		ChainIdCmd,
	},
}

var BlockCmd = cli.Command{
	Name:      "block",
	Usage:     "get block by number or hash",
	UsageText: "block [number|hash]",
	Action:    BlockAction,
	Flags:     []cli.Flag{},
}

var HeaderCmd = cli.Command{
	Name:      "header",
	Usage:     "get header by number or hash",
	UsageText: "header [number|hash]",
	Action:    HeaderAction,
	Flags:     []cli.Flag{},
}

var TransactionCmd = cli.Command{
	Name:   "transaction",
	Usage:  "get transaction by hash",
	Action: TransactionAction,
	Flags:  []cli.Flag{},
}

var ReceiptCmd = cli.Command{
	Name:   "receipt",
	Usage:  "get receipt by hash",
	Action: ReceiptAction,
	Flags:  []cli.Flag{},
}

var BalanceCmd = cli.Command{
	Name:      "balance",
	Usage:     "get balance of an address",
	UsageText: "balance <address> [number]",
	Action:    BalanceAction,
	Flags:     []cli.Flag{FlagNoPending},
}

var NonceCmd = cli.Command{
	Name:      "nonce",
	Usage:     "get nonce of an address",
	UsageText: "nonce <address> [number]",
	Action:    NonceAction,
	Flags:     []cli.Flag{FlagNoPending},
}

var CodeCmd = cli.Command{
	Name:      "code",
	Usage:     "get code of a contract",
	UsageText: "code <address> [number]",
	Action:    CodeAction,
	Flags:     []cli.Flag{FlagNoPending},
}

var StorageCmd = cli.Command{
	Name:      "storage",
	Usage:     "get storage value of a contract by key",
	UsageText: "storage <address> <hash> [number]",
	Action:    StorageAction,
	Flags:     []cli.Flag{FlagNoPending},
}

var SendCmd = cli.Command{
	Name:      "send",
	Usage:     "send a transaction",
	UsageText: "send <tx>",
	Action:    SendAction,
	Flags:     []cli.Flag{},
}

var CallCmd = cli.Command{
	Name:      "call",
	Usage:     "call contract",
	UsageText: "call <from> <to> <gas> <gasPrice> <value> <data> [number]",
	Action:    CallAction,
	Flags:     []cli.Flag{FlagNoPending},
}

var GasCmd = cli.Command{
	Name:      "gas",
	Usage:     "estimate gas",
	UsageText: "gas <from> <to> <gas> <gasPrice> <value> <data>",
	Action:    EstimateGasAction,
	Flags:     []cli.Flag{},
}

var PriceCmd = cli.Command{
	Name:      "price",
	Usage:     "get suggested gas price",
	UsageText: "price",
	Action:    SuggestGasPriceAction,
	Flags:     []cli.Flag{},
}

var ChainIdCmd = cli.Command{
	Name:      "id",
	Usage:     "get chain id",
	UsageText: "id",
	Action:    ChainIdAction,
	Flags:     []cli.Flag{},
}
