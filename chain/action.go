package chain

import (
	"fmt"
	"math"
	"strconv"

	cli "gopkg.in/urfave/cli.v1"
)

func BlockAction(c *cli.Context) error {
	if c.NArg() == 0 || c.Args()[0] == "latest" || c.Args()[0] == "" {
		return BlockByNumber(c.Parent().String(FlagRpcAddr.Name), math.MaxUint64)
	} else {
		if c.NArg() < 1 {
			return fmt.Errorf("insufficient arguments")
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
			return fmt.Errorf("insufficient arguments")
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
		return fmt.Errorf("insufficient arguments")
	}
	return TransactionByHash(c.Parent().String(FlagRpcAddr.Name), c.Args()[0])
}

func ReceiptAction(c *cli.Context) error {
	if c.NArg() < 1 {
		return fmt.Errorf("insufficient arguments")
	}
	return TransactionReceipt(c.Parent().String(FlagRpcAddr.Name), c.Args()[0])
}

func BalanceAction(c *cli.Context) error {
	if c.NArg() < 1 {
		return fmt.Errorf("insufficient arguments")
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
}

func NonceAction(c *cli.Context) error {
	if c.NArg() < 1 {
		return fmt.Errorf("insufficient arguments")
	}
	if c.BoolT(FlagNoPending.Name) {
		if c.NArg() < 2 {
			return NonceAt(c.Parent().String(FlagRpcAddr.Name), c.Args()[0], math.MaxUint64)
		} else {
			number, err := strconv.ParseUint(c.Args()[1], 10, 64)
			if err != nil {
				return err
			}
			return NonceAt(c.Parent().String(FlagRpcAddr.Name), c.Args()[0], number)
		}
	} else {
		return PendingNonceAt(c.Parent().String(FlagRpcAddr.Name), c.Args()[0])
	}
}

func CodeAction(c *cli.Context) error {
	if c.NArg() < 1 {
		return fmt.Errorf("insufficient arguments")
	}
	if c.BoolT(FlagNoPending.Name) {
		if c.NArg() < 2 {
			return CodeAt(c.Parent().String(FlagRpcAddr.Name), c.Args()[0], math.MaxUint64)
		} else {
			number, err := strconv.ParseUint(c.Args()[1], 10, 64)
			if err != nil {
				return err
			}
			return CodeAt(c.Parent().String(FlagRpcAddr.Name), c.Args()[0], number)
		}
	} else {
		return PendingCodeAt(c.Parent().String(FlagRpcAddr.Name), c.Args()[0])
	}
}

func StorageAction(c *cli.Context) error {
	if c.NArg() < 2 {
		return fmt.Errorf("insufficient arguments")
	}
	if c.BoolT(FlagNoPending.Name) {
		if c.NArg() < 3 {
			return StorageAt(c.Parent().String(FlagRpcAddr.Name), c.Args()[0], c.Args()[1], math.MaxUint64)
		} else {
			number, err := strconv.ParseUint(c.Args()[2], 10, 64)
			if err != nil {
				return err
			}
			return StorageAt(c.Parent().String(FlagRpcAddr.Name), c.Args()[0], c.Args()[1], number)
		}
	} else {
		return PendingStorageAt(c.Parent().String(FlagRpcAddr.Name), c.Args()[0], c.Args()[1])
	}
}

func SendAction(c *cli.Context) error {
	if c.NArg() < 1 {
		return fmt.Errorf("insufficient arguments")
	}
	return SendTransaction(c.Parent().String(FlagRpcAddr.Name), c.Args()[0])
}

func CallAction(c *cli.Context) error {
	if c.NArg() < 6 {
		return fmt.Errorf("insufficient arguments")
	}
	gas, err := strconv.ParseUint(c.Args()[2], 10, 64)
	if err != nil {
		return err
	}
	if c.BoolT(FlagNoPending.Name) {
		if c.NArg() < 7 {
			return CallContract(c.Parent().String(FlagRpcAddr.Name),
				c.Args()[0],
				c.Args()[1],
				gas,
				c.Args()[3],
				c.Args()[4],
				c.Args()[5],
				math.MaxUint64)
		} else {
			number, err := strconv.ParseUint(c.Args()[6], 10, 64)
			if err != nil {
				return err
			}
			return CallContract(c.Parent().String(FlagRpcAddr.Name),
				c.Args()[0],
				c.Args()[1],
				gas,
				c.Args()[3],
				c.Args()[4],
				c.Args()[5],
				number)
		}
	} else {
		return PendingCodeAt(c.Parent().String(FlagRpcAddr.Name), c.Args()[0])
	}
}

func EstimateGasAction(c *cli.Context) error {
	if c.NArg() < 6 {
		return fmt.Errorf("insufficient arguments")
	}
	gas, err := strconv.ParseUint(c.Args()[2], 10, 64)
	if err != nil {
		return err
	}

	return EstimateGas(c.Parent().String(FlagRpcAddr.Name), c.Args()[0],
		c.Args()[1],
		gas,
		c.Args()[3],
		c.Args()[4],
		c.Args()[5])
}

func SuggestGasPriceAction(c *cli.Context) error {
	return SuggestGasPrice(c.Parent().String(FlagRpcAddr.Name))
}

func ChainIdAction(c *cli.Context) error {
	return ChainId(c.Parent().String(FlagRpcAddr.Name))
}
