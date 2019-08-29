package chain

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func BalanceAt(rpcUrl, account, number string) error {
	cli, err := ethclient.DialContext(context.Background(), rpcUrl)
	if err != nil {
		return err
	}
	defer cli.Close()
	var num *big.Int = nil
	if number != "" {
		num, err = ParseBigInt(number)
		if err != nil {
			return err
		}
	}
	balance, err := cli.BalanceAt(context.Background(), common.HexToAddress(account), num)
	if err != nil {
		return err
	}
	fmt.Println(balance)
	return nil
}

func NonceAt(rpcUrl, account, number string) error {
	cli, err := ethclient.DialContext(context.Background(), rpcUrl)
	if err != nil {
		return err
	}
	defer cli.Close()
	var num *big.Int = nil
	if number != "" {
		num, err = ParseBigInt(number)
		if err != nil {
			return err
		}
	}
	nonce, err := cli.NonceAt(context.Background(), common.HexToAddress(account), num)
	if err != nil {
		return err
	}
	fmt.Println(nonce)
	return nil
}

func CodeAt(rpcUrl, account, number string) error {
	cli, err := ethclient.DialContext(context.Background(), rpcUrl)
	if err != nil {
		return err
	}
	defer cli.Close()
	var num *big.Int = nil
	if number != "" {
		num, err = ParseBigInt(number)
		if err != nil {
			return err
		}
	}
	code, err := cli.CodeAt(context.Background(), common.HexToAddress(account), num)
	if err != nil {
		return err
	}
	fmt.Println(common.ToHex(code))
	return nil
}

func StorageAt(rpcUrl, account, key, number string) error {
	cli, err := ethclient.DialContext(context.Background(), rpcUrl)
	if err != nil {
		return err
	}
	defer cli.Close()
	var num *big.Int = nil
	if number != "" {
		num, err = ParseBigInt(number)
		if err != nil {
			return err
		}
	}
	data, err := cli.StorageAt(context.Background(), common.HexToAddress(account), common.HexToHash(key), num)
	if err != nil {
		return err
	}
	fmt.Println(common.ToHex(data))
	return nil
}
