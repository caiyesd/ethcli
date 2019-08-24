package chain

import (
	"context"
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func BalanceAt(rpcUrl string, account string, number uint64) error {
	cli, err := ethclient.DialContext(context.Background(), rpcUrl)
	if err != nil {
		return err
	}
	defer cli.Close()
	var num *big.Int
	if number != math.MaxUint64 {
		num = new(big.Int).SetUint64(number)
	}
	balance, err := cli.BalanceAt(context.Background(), common.HexToAddress(account), num)
	if err != nil {
		return err
	}
	fmt.Println(balance)
	return nil
}

func NonceAt(rpcUrl string, account string, number uint64) error {
	cli, err := ethclient.DialContext(context.Background(), rpcUrl)
	if err != nil {
		return err
	}
	defer cli.Close()
	var num *big.Int
	if number != math.MaxUint64 {
		num = new(big.Int).SetUint64(number)
	}
	nonce, err := cli.NonceAt(context.Background(), common.HexToAddress(account), num)
	if err != nil {
		return err
	}
	fmt.Println(nonce)
	return nil
}

func CodeAt(rpcUrl string, account string, number uint64) error {
	cli, err := ethclient.DialContext(context.Background(), rpcUrl)
	if err != nil {
		return err
	}
	defer cli.Close()
	var num *big.Int
	if number != math.MaxUint64 {
		num = new(big.Int).SetUint64(number)
	}
	code, err := cli.CodeAt(context.Background(), common.HexToAddress(account), num)
	if err != nil {
		return err
	}
	fmt.Println(common.ToHex(code))
	return nil
}

func StorageAt(rpcUrl string, account string, key string, number uint64) error {
	cli, err := ethclient.DialContext(context.Background(), rpcUrl)
	if err != nil {
		return err
	}
	defer cli.Close()
	var num *big.Int
	if number != math.MaxUint64 {
		num = new(big.Int).SetUint64(number)
	}
	data, err := cli.StorageAt(context.Background(), common.HexToAddress(account), common.HexToHash(key), num)
	if err != nil {
		return err
	}
	fmt.Println(common.ToHex(data))
	return nil
}
