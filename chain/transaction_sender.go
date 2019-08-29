package chain

import (
	"bytes"
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

func SendTransaction(rpcUrl, rawTxStr string) error {
	cli, err := ethclient.DialContext(context.Background(), rpcUrl)
	if err != nil {
		return err
	}
	defer cli.Close()
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
