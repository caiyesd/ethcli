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
	for id, vlog := range r.Logs {
		fmt.Printf("Logs[%d]:\n", id)
		fmt.Printf("  Address:       %s\n", vlog.Address.String())
		fmt.Printf("  Topic:         %d\n", len(vlog.Topics))
		for _, topic := range vlog.Topics {
			fmt.Printf("                 %s\n", topic.String())
		}
		fmt.Printf("  Data:          %s\n", common.ToHex(vlog.Data))
		fmt.Printf("  Removed:       %v\n", vlog.Removed)
	}
}

func TransactionByHash(rpcUrl, hash string) error {
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

func TransactionReceipt(rpcUrl, hash string) error {
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
