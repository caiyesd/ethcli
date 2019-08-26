package chain

import (
	"context"
	"fmt"
	"math"
	"math/big"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func CallContract(rpcUrl string, from, to string, gas uint64, gasPrice, value, data string, number uint64) error {
	cli, err := ethclient.DialContext(context.Background(), rpcUrl)
	if err != nil {
		return err
	}
	defer cli.Close()
	var num *big.Int
	if number != math.MaxUint64 {
		num = new(big.Int).SetUint64(number)
	}

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

	result, err := cli.CallContract(context.Background(), ethereum.CallMsg{
		From:     _from,
		To:       _to,
		Gas:      gas,
		GasPrice: _gasPrice,
		Value:    _value,
		Data:     _data,
	}, num)
	if err != nil {
		return err
	}
	fmt.Println(common.ToHex(result))
	return nil
}
