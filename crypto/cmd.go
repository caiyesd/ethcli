package crypto

import (
	"errors"

	cli "gopkg.in/urfave/cli.v1"
)

var ErrInvalidArgs = errors.New("invalid arguments")
var ErrInvalidAmount = errors.New("invalid amount")

var CryptoCmd = cli.Command{
	Name:  "crypto",
	Usage: "manage and use keystore",
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
