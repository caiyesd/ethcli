package chain

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"os"
	"os/signal"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func printBlock(b *types.Block) {
	h := b.Header()
	txs := b.Transactions()
	fmt.Printf("Number:       %s\n", h.Number)
	fmt.Printf("Hash:         %s\n", h.Hash().String())
	fmt.Printf("Parent:       %s\n", h.ParentHash.String())
	fmt.Printf("Transactions: %d\n", len(txs))
	for i := 0; i < len(txs); i++ {
		fmt.Printf("  %s\n", txs[i].Hash().String())
	}
}

func printHeader(h *types.Header) {
	fmt.Printf("Number: %s\n", h.Number)
	fmt.Printf("Hash:   %s\n", h.Hash().String())
	fmt.Printf("Parent: %s\n", h.ParentHash.String())
}

func printTransaction(tx *types.Transaction) {
	fmt.Printf("Hash:      %s\n", tx.Hash().String())
	fmt.Printf("Nonce:     %d\n", tx.Nonce())
	if tx.To() == nil {
		fmt.Printf("To:        %s\n", tx.To())
	} else {
		fmt.Printf("To:        %s\n", tx.To().String())
	}

	fmt.Printf("Value:     %s\n", tx.Value())
	fmt.Printf("Gas:       %d\n", tx.Gas())
	fmt.Printf("GasPrice:  %s\n", tx.GasPrice())
	fmt.Printf("Data:      %s\n", common.ToHex(tx.Data()))
	fmt.Printf("ChainId:   %s\n", tx.ChainId())
	v, r, s := tx.RawSignatureValues()
	fmt.Printf("V:         %s\n", common.ToHex(v.Bytes()))
	fmt.Printf("R:         %s\n", common.ToHex(r.Bytes()))
	fmt.Printf("S:         %s\n", common.ToHex(s.Bytes()))
	sig := [65]byte{}
	copy(sig[0:32], r.Bytes())
	copy(sig[32:64], s.Bytes())
	copy(sig[64:65], v.Bytes())
	fmt.Printf("Signature: %s\n", common.ToHex(sig[:]))
}

func HeaderByHash(rpcUrl string, hash string) error {
	cli, err := ethclient.DialContext(context.Background(), rpcUrl)
	if err != nil {
		return err
	}
	defer cli.Close()
	header, err := cli.HeaderByHash(context.Background(), common.HexToHash(hash))
	if err != nil {
		return err
	}
	printHeader(header)
	return nil
}

func BlockByHash(rpcUrl string, hash string) error {
	cli, err := ethclient.DialContext(context.Background(), rpcUrl)
	if err != nil {
		return err
	}
	defer cli.Close()
	block, err := cli.BlockByHash(context.Background(), common.HexToHash(hash))
	if err != nil {
		return err
	}
	printBlock(block)
	return nil
}

func TransactionCount(rpcUrl string, hash string) error {
	cli, err := ethclient.DialContext(context.Background(), rpcUrl)
	if err != nil {
		return err
	}
	defer cli.Close()
	count, err := cli.TransactionCount(context.Background(), common.HexToHash(hash))
	fmt.Println(count)
	return nil
}

func HeaderByNumber(rpcUrl string, number uint64) error {
	cli, err := ethclient.DialContext(context.Background(), rpcUrl)
	if err != nil {
		return err
	}
	defer cli.Close()
	var num *big.Int
	if number != math.MaxUint64 {
		num = new(big.Int).SetUint64(number)
	}
	header, err := cli.HeaderByNumber(context.Background(), num)
	if err != nil {
		return err
	}
	printHeader(header)
	return nil
}

func BlockByNumber(rpcUrl string, number uint64) error {
	cli, err := ethclient.DialContext(context.Background(), rpcUrl)
	if err != nil {
		return err
	}
	defer cli.Close()
	var num *big.Int
	if number != math.MaxUint64 {
		num = new(big.Int).SetUint64(number)
	}
	block, err := cli.BlockByNumber(context.Background(), num)
	if err != nil {
		return err
	}
	printBlock(block)
	return nil
}

func TransactionInBlock(rpcUrl string, hash string, index uint) error {
	cli, err := ethclient.DialContext(context.Background(), rpcUrl)
	if err != nil {
		return err
	}
	defer cli.Close()
	tx, err := cli.TransactionInBlock(context.Background(), common.HexToHash(hash), index)
	if err != nil {
		return err
	}
	printTransaction(tx)
	return nil
}

func SubscribeNewHead(rpcUrl string) error {
	cli, err := ethclient.DialContext(context.Background(), rpcUrl)
	if err != nil {
		return err
	}
	defer cli.Close()
	c := make(chan *types.Header)
	sub, err := cli.SubscribeNewHead(context.Background(), c)
	if err != nil {
		return err
	}

	go func() {
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, os.Interrupt)
		<-signalChan
		// fmt.Println("Received an interrupt, stopping subscription...\n")
		sub.Unsubscribe()
	}()

	for {
		select {
		case header, ok := <-c:
			if !ok {
				break
			}
			printHeader(header)
			break
		case err := <-sub.Err():
			sub.Unsubscribe()
			return err
		}
	}
	return nil
}
