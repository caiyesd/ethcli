package chain

import (
	"fmt"
	"strconv"

	cli "gopkg.in/urfave/cli.v1"
)

var (
	ErrInsufficientArguments = fmt.Errorf("insufficient arguments")
)

func BlockAction(c *cli.Context) error {
	if c.NArg() == 0 || c.Args()[0] == "latest" || c.Args()[0] == "" {
		return BlockByNumber(ParseFlagRpcAddr(c), "")
	} else {
		if c.NArg() < 1 {
			return ErrInsufficientArguments
		}
		_, err := strconv.ParseUint(c.Args()[0], 10, 64)
		if err != nil {
			return BlockByHash(ParseFlagRpcAddr(c), c.Args()[0])
		} else {
			return BlockByNumber(ParseFlagRpcAddr(c), c.Args()[0])
		}
	}
}

func HeaderAction(c *cli.Context) error {
	if c.NArg() == 0 || c.Args()[0] == "latest" || c.Args()[0] == "" {
		return HeaderByNumber(ParseFlagRpcAddr(c), "")
	} else {
		if c.NArg() < 1 {
			return ErrInsufficientArguments
		}
		_, err := strconv.ParseUint(c.Args()[0], 10, 64)
		if err != nil {
			return HeaderByHash(ParseFlagRpcAddr(c), c.Args()[0])
		} else {
			return HeaderByNumber(ParseFlagRpcAddr(c), c.Args()[0])
		}
	}
}

func SubNewHeaderAction(c *cli.Context) error {
	return SubscribeNewHead(ParseFlagRpcAddr(c))
}

func TransactionAction(c *cli.Context) error {
	if c.NArg() < 1 {
		return ErrInsufficientArguments
	}
	return TransactionByHash(ParseFlagRpcAddr(c), c.Args()[0])
}

func ReceiptAction(c *cli.Context) error {
	if c.NArg() < 1 {
		return ErrInsufficientArguments
	}
	return TransactionReceipt(ParseFlagRpcAddr(c), c.Args()[0])
}

func BalanceAction(c *cli.Context) error {
	if c.NArg() < 1 {
		return ErrInsufficientArguments
	}
	if ParseFlagNoPending(c) {
		if c.NArg() < 2 {
			return BalanceAt(ParseFlagRpcAddr(c), c.Args()[0], "")
		} else {
			return BalanceAt(ParseFlagRpcAddr(c), c.Args()[0], c.Args()[1])
		}
	} else {
		return PendingBalanceAt(ParseFlagRpcAddr(c), c.Args()[0])
	}
}

func NonceAction(c *cli.Context) error {
	if c.NArg() < 1 {
		return ErrInsufficientArguments
	}
	if ParseFlagNoPending(c) {
		if c.NArg() < 2 {
			return NonceAt(ParseFlagRpcAddr(c), c.Args()[0], "")
		} else {
			return NonceAt(ParseFlagRpcAddr(c), c.Args()[0], c.Args()[1])
		}
	} else {
		return PendingNonceAt(ParseFlagRpcAddr(c), c.Args()[0])
	}
}

func CodeAction(c *cli.Context) error {
	if c.NArg() < 1 {
		return ErrInsufficientArguments
	}
	if ParseFlagNoPending(c) {
		if c.NArg() < 2 {
			return CodeAt(ParseFlagRpcAddr(c), c.Args()[0], "")
		} else {
			return CodeAt(ParseFlagRpcAddr(c), c.Args()[0], c.Args()[1])
		}
	} else {
		return PendingCodeAt(ParseFlagRpcAddr(c), c.Args()[0])
	}
}

func StorageAction(c *cli.Context) error {
	if c.NArg() < 2 {
		return ErrInsufficientArguments
	}
	if ParseFlagNoPending(c) {
		if c.NArg() < 3 {
			return StorageAt(ParseFlagRpcAddr(c), c.Args()[0], c.Args()[1], "")
		} else {
			return StorageAt(ParseFlagRpcAddr(c), c.Args()[0], c.Args()[1], c.Args()[2])
		}
	} else {
		return PendingStorageAt(ParseFlagRpcAddr(c), c.Args()[0], c.Args()[1])
	}
}

func SendAction(c *cli.Context) error {
	if c.NArg() < 1 {
		return ErrInsufficientArguments
	}
	return SendTransaction(ParseFlagRpcAddr(c), c.Args()[0])
}

func CallAction(c *cli.Context) error {
	if c.NArg() < 6 {
		return ErrInsufficientArguments
	}
	if ParseFlagNoPending(c) {
		if c.NArg() < 7 {
			return CallContract(ParseFlagRpcAddr(c),
				c.Args()[0],
				c.Args()[1],
				c.Args()[2],
				c.Args()[3],
				c.Args()[4],
				c.Args()[5],
				"")
		} else {
			return CallContract(ParseFlagRpcAddr(c),
				c.Args()[0],
				c.Args()[1],
				c.Args()[2],
				c.Args()[3],
				c.Args()[4],
				c.Args()[5],
				c.Args()[6])
		}
	} else {
		return PendingCodeAt(ParseFlagRpcAddr(c), c.Args()[0])
	}
}

func EstimateGasAction(c *cli.Context) error {
	if c.NArg() < 6 {
		return ErrInsufficientArguments
	}

	return EstimateGas(ParseFlagRpcAddr(c), c.Args()[0],
		c.Args()[1],
		c.Args()[2],
		c.Args()[3],
		c.Args()[4],
		c.Args()[5])
}

func SuggestGasPriceAction(c *cli.Context) error {
	return SuggestGasPrice(ParseFlagRpcAddr(c))
}

func ChainIdAction(c *cli.Context) error {
	return ChainId(ParseFlagRpcAddr(c))
}
