package crypto

import (
	"fmt"
	"strconv"

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
		return fmt.Errorf("insufficient arguments")
	}
	return UpdateAccount(c.Parent().String(FlagKeystore.Name), c.Args()[0])
}

func DeleteAccountAction(c *cli.Context) error {
	if c.NArg() < 1 {
		return fmt.Errorf("insufficient arguments")
	}
	return DeleteAccount(c.Parent().String(FlagKeystore.Name), c.Args()[0])
}

func NewHashAction(c *cli.Context) error {
	if c.NArg() < 1 {
		return fmt.Errorf("insufficient arguments")
	}
	return NewHash(c.Args()[0])
}

func SignHashAction(c *cli.Context) error {
	if c.NArg() < 2 {
		return fmt.Errorf("insufficient arguments")
	}
	return SignHash(c.Parent().String(FlagKeystore.Name), c.Args()[0], c.Args()[1])
}

func NewTxAction(c *cli.Context) error {
	if c.NArg() < 5 {
		return fmt.Errorf("insufficient arguments")
	}
	nonce, err := strconv.ParseUint(c.Args()[0], 10, 64)
	if err != nil {
		return err
	}
	gasLimit, err := strconv.ParseUint(c.Args()[3], 10, 64)
	if err != nil {
		return err
	}
	if c.NArg() < 6 {
		return NewTx(nonce, c.Args()[1], c.Args()[2], gasLimit, c.Args()[4], "")
	} else {
		return NewTx(nonce, c.Args()[1], c.Args()[2], gasLimit, c.Args()[4], c.Args()[5])
	}
}

func SignTxAction(c *cli.Context) error {
	if c.NArg() < 2 {
		return fmt.Errorf("insufficient arguments")
	}
	if c.NArg() < 3 {
		return SignTx(c.Parent().String(FlagKeystore.Name), c.Args()[0], c.Args()[1], "")
	} else {
		return SignTx(c.Parent().String(FlagKeystore.Name), c.Args()[0], c.Args()[1], c.Args()[2])
	}
}

func VerifyAction(c *cli.Context) error {
	if c.NArg() < 3 {
		return fmt.Errorf("insufficient arguments")
	}
	return Verify(c.Args()[0], c.Args()[1], c.Args()[2])
}

func EcrecoverAction(c *cli.Context) error {
	if c.NArg() < 2 {
		return fmt.Errorf("insufficient arguments")
	}
	return Ecrecover(c.Args()[0], c.Args()[1])
}
