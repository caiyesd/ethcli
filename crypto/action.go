package crypto

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	cli "gopkg.in/urfave/cli.v1"
)

var (
	ErrInsufficientArguments = fmt.Errorf("insufficient arguments")
)

func CalcHashAction(c *cli.Context) error {
	if c.NArg() < 1 {
		return ErrInsufficientArguments
	}
	data := common.FromHex(c.Args()[0])
	h := crypto.Keccak256Hash(data)
	fmt.Println(h.String())
	return nil
}

func SignHashAction(c *cli.Context) error {
	if c.NArg() < 2 {
		return ErrInsufficientArguments
	}
	passphrase, err := ReadPassphrase(c)
	if err != nil {
		return err
	}
	return SignHash(c.Parent().String(FlagKeystore.Name), c.Args()[0], passphrase, c.Args()[1])
}

func BuildTxAction(c *cli.Context) error {
	if c.NArg() < 5 {
		return ErrInsufficientArguments
	}

	if c.NArg() < 6 {
		return BuildTx(c.Args()[0], c.Args()[1], c.Args()[2], c.Args()[3], c.Args()[4], "")
	} else {
		return BuildTx(c.Args()[0], c.Args()[1], c.Args()[2], c.Args()[3], c.Args()[4], c.Args()[5])
	}
}

func SignTxAction(c *cli.Context) error {
	if c.NArg() < 2 {
		return ErrInsufficientArguments
	}
	passphrase, err := ReadPassphrase(c)
	if err != nil {
		return err
	}
	if c.NArg() < 3 {
		return SignTx(c.Parent().String(FlagKeystore.Name), c.Args()[0], passphrase, c.Args()[1], "")
	} else {
		return SignTx(c.Parent().String(FlagKeystore.Name), c.Args()[0], passphrase, c.Args()[1], c.Args()[2])
	}
}

func VerifyAction(c *cli.Context) error {
	if c.NArg() < 3 {
		return ErrInsufficientArguments
	}
	return Verify(c.Args()[0], c.Args()[1], c.Args()[2])
}

func EcrecoverAction(c *cli.Context) error {
	if c.NArg() < 2 {
		return ErrInsufficientArguments
	}
	return Ecrecover(c.Args()[0], c.Args()[1])
}

func PackAction(c *cli.Context) error {
	if c.NArg()%2 != 0 {
		return ErrInsufficientArguments
	}
	return PackArguments([]string(c.Args()))
}
