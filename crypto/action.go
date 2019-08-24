package crypto

import (
	cli "gopkg.in/urfave/cli.v1"
)

func NewAccountAction(c *cli.Context) error {
	return NewAccount(c.Parent().String(FlagKeystore.Name))
}

func ListAccountAction(c *cli.Context) error {
	return ListAccount(c.Parent().String(FlagKeystore.Name))
}

func UpdateAccountAction(c *cli.Context) error {
	if c.NArg() < 1 {
		return ErrInvalidArgs
	}
	return UpdateAccount(c.Parent().String(FlagKeystore.Name), c.Args()[0])
}

func DeleteAccountAction(c *cli.Context) error {
	if c.NArg() < 1 {
		return ErrInvalidArgs
	}
	return DeleteAccount(c.Parent().String(FlagKeystore.Name), c.Args()[0])
}
