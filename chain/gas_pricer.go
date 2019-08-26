package chain

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

func SuggestGasPrice(rpcUrl string) error {
	cli, err := ethclient.DialContext(context.Background(), rpcUrl)
	if err != nil {
		return err
	}

	price, err := cli.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}
	fmt.Println(price)
	return nil
}
