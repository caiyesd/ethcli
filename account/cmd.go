package account

import (
	cli "gopkg.in/urfave/cli.v1"
)

var AccountCmd = cli.Command{
	Name:  "account",
	Usage: "manage accounts",
	Flags: []cli.Flag{FlagKeystore},
	Subcommands: []cli.Command{
		NewAccountCmd,
		ListAccountCmd,
		UpdateAccountCmd,
		DeleteAccountCmd,
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
	Usage:     "list account",
	UsageText: "list [index]",
	Action:    ListAccountAction,
	Flags:     []cli.Flag{},
}

var UpdateAccountCmd = cli.Command{
	Name:      "update",
	Usage:     "update an account",
	UsageText: "update <address|index>",
	Action:    UpdateAccountAction,
	Flags:     []cli.Flag{},
}

var DeleteAccountCmd = cli.Command{
	Name:      "delete",
	Usage:     "delete an account",
	UsageText: "delete <account|index>",
	Action:    DeleteAccountAction,
	Flags:     []cli.Flag{},
}
