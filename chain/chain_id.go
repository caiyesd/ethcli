package chain

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

func ChainId(rpcUrl string) error {
	cli, err := ethclient.DialContext(context.Background(), rpcUrl)
	if err != nil {
		return err
	}

	id, err := cli.ChainID(context.Background())
	if err != nil {
		return err
	}
	fmt.Println(id)
	return nil
}
