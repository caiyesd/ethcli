package account

import (
	"fmt"

	cli "gopkg.in/urfave/cli.v1"
)

var (
	ErrInsufficientArguments = fmt.Errorf("insufficient arguments")
)

func NewAccountAction(c *cli.Context) error {
	return NewAccount(ParseFlagKeystore(c))
}

func ListAccountAction(c *cli.Context) error {
	if c.NArg() < 1 {
		return ListAccount(ParseFlagKeystore(c), "")
	} else {
		return ListAccount(ParseFlagKeystore(c), c.Args()[0])
	}
}

func UpdateAccountAction(c *cli.Context) error {
	if c.NArg() < 1 {
		return ErrInsufficientArguments
	}
	return UpdateAccount(ParseFlagKeystore(c), c.Args()[0])

}

func DeleteAccountAction(c *cli.Context) error {
	if c.NArg() < 1 {
		return ErrInsufficientArguments
	}
	return DeleteAccount(ParseFlagKeystore(c), c.Args()[0])
}
