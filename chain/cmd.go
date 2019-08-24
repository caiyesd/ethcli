package chain

import (
	"errors"

	cli "gopkg.in/urfave/cli.v1"
)

var ErrInvalidArgs = errors.New("invalid arguments")
var ErrInvalidAmount = errors.New("invalid amount")

var ChainCmd = cli.Command{
	Name:  "chain",
	Usage: "communicate with chain",
	Flags: []cli.Flag{FlagRpcAddr},
	Subcommands: []cli.Command{
		HeaderCmd,
		BlockCmd,
		TransactionCmd,
		ReceiptCmd,
		SubNewHeaderCmd,
		BalanceCmd,
		NonceCmd,
		CodeCmd,
		StorageCmd,
		RawTxCmd,
		SendCmd,
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

var SubNewHeaderCmd = cli.Command{
	Name:      "sub",
	Usage:     "sub new header",
	UsageText: "sub",
	Action:    SubNewHeaderAction,
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
	Usage:     "get code of an contract",
	UsageText: "code <address> [number]",
	Action:    CodeAction,
	Flags:     []cli.Flag{FlagNoPending},
}

var StorageCmd = cli.Command{
	Name:      "storage",
	Usage:     "get value by key",
	UsageText: "storage <address> <hash> [number]",
	Action:    StorageAction,
	Flags:     []cli.Flag{FlagNoPending},
}

var RawTxCmd = cli.Command{
	Name:      "rawtx",
	Usage:     "build rawtx",
	UsageText: "rawtx <nonce> <to> <amount> <gasLimit> <gasPrice> <data>",
	Action:    RawTxAction,
	Flags:     []cli.Flag{},
}

var SendCmd = cli.Command{
	Name:      "send",
	Usage:     "send rawtx",
	UsageText: "send <rawtx>",
	Action:    SendAction,
	Flags:     []cli.Flag{},
}
