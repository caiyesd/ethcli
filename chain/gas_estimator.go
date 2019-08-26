package chain

import (
	"context"
	"fmt"
	"math/big"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func EstimateGas(rpcUrl string, from, to string, gas uint64, gasPrice, value, data string) error {
	cli, err := ethclient.DialContext(context.Background(), rpcUrl)
	if err != nil {
		return err
	}
	defer cli.Close()

	_from := common.HexToAddress(from)
	var _to *common.Address = nil
	if to != "" {
		tmp := common.HexToAddress(to)
		_to = &tmp
	}
	_gasPrice, ok := new(big.Int).SetString(gasPrice, 10)
	if !ok {
		return fmt.Errorf("failed to parse %s", gasPrice)
	}
	_value, ok := new(big.Int).SetString(value, 10)
	if !ok {
		return fmt.Errorf("failed to parse %s", value)
	}
	_data := common.FromHex(data)

	price, err := cli.EstimateGas(context.Background(), ethereum.CallMsg{
		From:     _from,
		To:       _to,
		Gas:      gas,
		GasPrice: _gasPrice,
		Value:    _value,
		Data:     _data,
	})
	if err != nil {
		return err
	}
	fmt.Println(price)
	return nil
}
