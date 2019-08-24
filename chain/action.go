package chain

import (
	"math"
	"strconv"

	cli "gopkg.in/urfave/cli.v1"
)

func BlockAction(c *cli.Context) error {
	if c.NArg() == 0 || c.Args()[0] == "latest" || c.Args()[0] == "" {
		return BlockByNumber(c.Parent().String(FlagRpcAddr.Name), math.MaxUint64)
	} else {
		if c.NArg() < 1 {
			return ErrInvalidArgs
		}
		number, err := strconv.ParseUint(c.Args()[0], 10, 64)
		if err != nil {
			return BlockByHash(c.Parent().String(FlagRpcAddr.Name), c.Args()[0])
		} else {
			return BlockByNumber(c.Parent().String(FlagRpcAddr.Name), number)
		}
	}
}

func HeaderAction(c *cli.Context) error {
	if c.NArg() == 0 || c.Args()[0] == "latest" || c.Args()[0] == "" {
		return HeaderByNumber(c.Parent().String(FlagRpcAddr.Name), math.MaxUint64)
	} else {
		if c.NArg() < 1 {
			return ErrInvalidArgs
		}
		number, err := strconv.ParseUint(c.Args()[0], 10, 64)
		if err != nil {
			return HeaderByHash(c.Parent().String(FlagRpcAddr.Name), c.Args()[0])
		} else {
			return HeaderByNumber(c.Parent().String(FlagRpcAddr.Name), number)
		}
	}
}

func SubNewHeaderAction(c *cli.Context) error {
	return SubscribeNewHead(c.Parent().String(FlagRpcAddr.Name))
}

func TransactionAction(c *cli.Context) error {
	if c.NArg() < 1 {
		return ErrInvalidArgs
	}
	return TransactionByHash(c.Parent().String(FlagRpcAddr.Name), c.Args()[0])
}

func ReceiptAction(c *cli.Context) error {
	if c.NArg() < 1 {
		return ErrInvalidArgs
	}
	return TransactionReceipt(c.Parent().String(FlagRpcAddr.Name), c.Args()[0])
}

func BalanceAction(c *cli.Context) error {
	if c.NArg() < 1 {
		return ErrInvalidArgs
	}
	if c.BoolT(FlagNoPending.Name) {
		if c.NArg() < 2 {
			return BalanceAt(c.Parent().String(FlagRpcAddr.Name), c.Args()[0], math.MaxUint64)
		} else {
			number, err := strconv.ParseUint(c.Args()[1], 10, 64)
			if err != nil {
				return err
			}
			return BalanceAt(c.Parent().String(FlagRpcAddr.Name), c.Args()[0], number)
		}
	} else {
		return PendingBalanceAt(c.Parent().String(FlagRpcAddr.Name), c.Args()[0])
	}

	return nil
}

func NonceAction(c *cli.Context) error {
	if c.NArg() < 1 {
		return ErrInvalidArgs
	}
	if c.BoolT(FlagNoPending.Name) {
		if c.NArg() < 2 {
			BalanceAt(c.Parent().String(FlagRpcAddr.Name), c.Args()[0], math.MaxUint64)
		} else {
			number, err := strconv.ParseUint(c.Args()[1], 10, 64)
			if err != nil {
				return err
			}
			NonceAt(c.Parent().String(FlagRpcAddr.Name), c.Args()[0], number)
		}
	} else {
		return PendingNonceAt(c.Parent().String(FlagRpcAddr.Name), c.Args()[0])
	}

	return nil
}

func CodeAction(c *cli.Context) error {
	if c.NArg() < 1 {
		return ErrInvalidArgs
	}
	if c.BoolT(FlagNoPending.Name) {
		if c.NArg() < 2 {
			CodeAt(c.Parent().String(FlagRpcAddr.Name), c.Args()[0], math.MaxUint64)
		} else {
			number, err := strconv.ParseUint(c.Args()[1], 10, 64)
			if err != nil {
				return err
			}
			CodeAt(c.Parent().String(FlagRpcAddr.Name), c.Args()[0], number)
		}
	} else {
		return PendingCodeAt(c.Parent().String(FlagRpcAddr.Name), c.Args()[0])
	}

	return nil
}

func StorageAction(c *cli.Context) error {
	if c.NArg() < 2 {
		return ErrInvalidArgs
	}
	if c.BoolT(FlagNoPending.Name) {
		if c.NArg() < 3 {
			StorageAt(c.Parent().String(FlagRpcAddr.Name), c.Args()[0], c.Args()[1], math.MaxUint64)
		} else {
			number, err := strconv.ParseUint(c.Args()[2], 10, 64)
			if err != nil {
				return err
			}
			StorageAt(c.Parent().String(FlagRpcAddr.Name), c.Args()[0], c.Args()[1], number)
		}
	} else {
		return PendingStorageAt(c.Parent().String(FlagRpcAddr.Name), c.Args()[0], c.Args()[1])
	}
	return nil
}

func RawTxAction(c *cli.Context) error {
	if c.NArg() < 6 {
		return ErrInvalidArgs
	}
	nonce, err := strconv.ParseUint(c.Args()[0], 10, 64)
	if err != nil {
		return err
	}
	to := c.Args()[1]
	amount := c.Args()[2]
	gasLimit, err := strconv.ParseUint(c.Args()[3], 10, 64)
	if err != nil {
		return err
	}
	gasPrice := c.Args()[4]
	data := c.Args()[5]
	return BuildTransaction(nonce, to, amount, gasLimit, gasPrice, data)
}

func SendAction(c *cli.Context) error {
	if c.NArg() < 1 {
		return ErrInvalidArgs
	}
	return SendTransaction(c.Parent().String(FlagRpcAddr.Name), c.Args()[0])
}
