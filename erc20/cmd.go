package erc20

import (
	cli "gopkg.in/urfave/cli.v1"
)

var Erc20Cmd = cli.Command{
	Name:  "erc20",
	Usage: "interact with erc20 token",
	Flags: []cli.Flag{FlagKeystore, FlagRpcAddr, FlagErc20Addr},
	Subcommands: []cli.Command{
		TotalSupplyCmd,
		BalanceOfCmd,
		AllowanceCmd,
		TransferCmd,
		TransferFromCmd,
		ApproveCmd,
		IsPauserCmd,
		AddPauserCmd,
		RenouncePauserCmd,
		PausedCmd,
		PauseCmd,
		UnpauseCmd,
		IsMinterCmd,
		AddMinterCmd,
		RenouncePauserCmd,
		MintCmd,
		BurnCmd,
		BurnFromCmd,
	},
}

var TotalSupplyCmd = cli.Command{
	Name:      "supply",
	Usage:     "get total supply",
	UsageText: "supply",
	Action:    TotalSupplyAction,
	Flags:     []cli.Flag{},
}

var BalanceOfCmd = cli.Command{
	Name:      "balanceof",
	Usage:     "get balance of OWNER account",
	UsageText: "balanceof <owner>",
	Action:    BalanceOfAction,
	Flags:     []cli.Flag{},
}

var AllowanceCmd = cli.Command{
	Name:      "allowance",
	Usage:     "get SPENDER's allowance of OWNER account",
	UsageText: "allowance <owner> <spender>",
	Action:    AllowanceAction,
	Flags:     []cli.Flag{},
}

var TransferCmd = cli.Command{
	Name:      "transfer",
	Usage:     "transfer VALUE tokens from this account to TO account",
	UsageText: "transfer <account> <to> <value>",
	Action:    TransferAction,
	Flags:     []cli.Flag{},
}

var TransferFromCmd = cli.Command{
	Name:      "transfer-from",
	Usage:     "transfer VALUE tokens from FROM account to TO account",
	UsageText: "transfer-from <account> <from> <to> <value>",
	Action:    TransferFromAction,
	Flags:     []cli.Flag{},
}

var ApproveCmd = cli.Command{
	Name:      "approve",
	Usage:     "approve SPENDER spends VALUE tokens",
	UsageText: "approve <account> <spender> <value>",
	Action:    ApproveAction,
	Flags:     []cli.Flag{},
}

var IsPauserCmd = cli.Command{
	Name:      "is-pauser",
	Usage:     "check whether ACCOUNT is a pauser",
	UsageText: "is-pauser <pauser>",
	Action:    IsPauserAction,
	Flags:     []cli.Flag{},
}

var AddPauserCmd = cli.Command{
	Name:      "add-pauser",
	Usage:     "add a new pauser",
	UsageText: "add-pauser <account> <pauser>",
	Action:    AddPauserAction,
	Flags:     []cli.Flag{},
}

var RenouncePauserCmd = cli.Command{
	Name:      "renounce-pauser",
	Usage:     "renounce pauser",
	UsageText: "renounce-pauser <account>",
	Action:    RenouncePauserAction,
	Flags:     []cli.Flag{},
}

var PausedCmd = cli.Command{
	Name:      "paused",
	Usage:     "get paused state",
	UsageText: "paused",
	Action:    PausedAction,
	Flags:     []cli.Flag{},
}

var PauseCmd = cli.Command{
	Name:      "pause",
	Usage:     "pause contract",
	UsageText: "pause <account>",
	Action:    PauseAction,
	Flags:     []cli.Flag{},
}

var UnpauseCmd = cli.Command{
	Name:      "unpause",
	Usage:     "unpause contract",
	UsageText: "unpause <account>",
	Action:    UnpauseAction,
	Flags:     []cli.Flag{},
}

var IsMinterCmd = cli.Command{
	Name:      "is-minter",
	Usage:     "check whether ACCOUNT is a minter",
	UsageText: "is-minter <minter>",
	Action:    IsMinterAction,
	Flags:     []cli.Flag{},
}

var AddMinterCmd = cli.Command{
	Name:      "add-minter",
	Usage:     "add a new minter",
	UsageText: "add-minter <account> <minter>",
	Action:    AddMinterAction,
	Flags:     []cli.Flag{},
}

var RenounceMinterCmd = cli.Command{
	Name:      "renounce-minter",
	Usage:     "renounce minter",
	UsageText: "renounce-minter <account>",
	Action:    RenounceMinterAction,
	Flags:     []cli.Flag{},
}

var MintCmd = cli.Command{
	Name:      "mint",
	Usage:     "mint tokens to TO address",
	UsageText: "mint <account> <to> <value>",
	Action:    MintAction,
	Flags:     []cli.Flag{},
}

var BurnCmd = cli.Command{
	Name:      "burn",
	Usage:     "burn tokens",
	UsageText: "mint <account> <value>",
	Action:    BurnAction,
	Flags:     []cli.Flag{},
}

var BurnFromCmd = cli.Command{
	Name:      "burn-from",
	Usage:     "burn tokens for ACCOUNT",
	UsageText: "burn-from <account> <from> <value>",
	Action:    BurnFromAction,
	Flags:     []cli.Flag{},
}
