package chain

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func PendingBalanceAt(rpcUrl string, account string) error {
	cli, err := ethclient.DialContext(context.Background(), rpcUrl)
	if err != nil {
		return err
	}
	defer cli.Close()
	balance, err := cli.PendingBalanceAt(context.Background(), common.HexToAddress(account))
	if err != nil {
		return err
	}
	fmt.Println(balance)
	return nil
}

func PendingNonceAt(rpcUrl string, account string) error {
	cli, err := ethclient.DialContext(context.Background(), rpcUrl)
	if err != nil {
		return err
	}
	defer cli.Close()
	nonce, err := cli.PendingNonceAt(context.Background(), common.HexToAddress(account))
	if err != nil {
		return err
	}
	fmt.Println(nonce)
	return nil
}

func PendingCodeAt(rpcUrl string, account string) error {
	cli, err := ethclient.DialContext(context.Background(), rpcUrl)
	if err != nil {
		return err
	}
	defer cli.Close()
	code, err := cli.PendingCodeAt(context.Background(), common.HexToAddress(account))
	if err != nil {
		return err
	}
	fmt.Println(common.ToHex(code))
	return nil
}

func PendingStorageAt(rpcUrl string, account string, key string) error {
	cli, err := ethclient.DialContext(context.Background(), rpcUrl)
	if err != nil {
		return err
	}
	defer cli.Close()
	data, err := cli.PendingStorageAt(context.Background(), common.HexToAddress(account), common.HexToHash(key))
	if err != nil {
		return err
	}
	fmt.Println(common.ToHex(data))
	return nil
}
