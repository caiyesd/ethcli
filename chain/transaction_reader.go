package chain

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func printReceipt(r *types.Receipt) {
	fmt.Printf("BlockHash:       %s\n", r.BlockHash.String())
	fmt.Printf("BlockNumber:     %s\n", r.BlockNumber)
	fmt.Printf("ContractAddress: %s\n", r.ContractAddress.String())
	fmt.Printf("GasUsed:         %d\n", r.GasUsed)
	fmt.Printf("Status:          %d\n", r.Status)
	fmt.Printf("TxHash:          %s\n", r.TxHash.String())
}

func TransactionByHash(rpcUrl string, hash string) error {
	cli, err := ethclient.DialContext(context.Background(), rpcUrl)
	if err != nil {
		return err
	}
	defer cli.Close()
	tx, isPending, err := cli.TransactionByHash(context.Background(), common.HexToHash(hash))
	if err != nil {
		return err
	}
	fmt.Printf("Pending:  %v\n", isPending)
	printTransaction(tx)
	return nil
}

func TransactionReceipt(rpcUrl string, hash string) error {
	cli, err := ethclient.DialContext(context.Background(), rpcUrl)
	if err != nil {
		return err
	}
	defer cli.Close()
	receipt, err := cli.TransactionReceipt(context.Background(), common.HexToHash(hash))
	if err != nil {
		return err
	}
	printReceipt(receipt)
	return nil
}
