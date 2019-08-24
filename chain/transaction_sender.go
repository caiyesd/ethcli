package chain

import (
	"bytes"
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

func SendTransaction(rpcUrl string, rawTxStr string) error {
	cli, err := ethclient.DialContext(context.Background(), rpcUrl)
	if err != nil {
		return err
	}
	tx := new(types.Transaction)
	err = tx.DecodeRLP(rlp.NewStream(bytes.NewBuffer(common.FromHex(rawTxStr)), 0))
	if err != nil {
		return err
	}
	// printTransaction(tx)
	err = cli.SendTransaction(context.Background(), tx)
	if err != nil {
		return err
	}
	return nil
}

func BuildTransaction(nonce uint64, to string, amount string, gasLimit uint64, gasPrice string, data string) error {
	_amount, ok := new(big.Int).SetString(amount, 10)
	if !ok {
		return ErrInvalidAmount
	}
	_gasPrice, ok := new(big.Int).SetString(gasPrice, 10)
	if !ok {
		return ErrInvalidAmount
	}
	tx := types.NewTransaction(nonce, common.HexToAddress(to), _amount, gasLimit, _gasPrice, common.Hex2Bytes(data))
	// printTransaction(tx)
	buffer := bytes.NewBuffer(nil)
	err := tx.EncodeRLP(buffer)
	if err != nil {
		return err
	}
	fmt.Println(common.ToHex(buffer.Bytes()))
	return nil
}
