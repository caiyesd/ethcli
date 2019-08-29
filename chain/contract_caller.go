package chain

import (
	"context"
	"fmt"
	"math/big"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func CallContract(rpcUrl, from, to, gas, gasPrice, value, data, number string) error {
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

	_from := common.HexToAddress(from)

	var _to *common.Address = nil
	if to != "" {
		tmp := common.HexToAddress(to)
		_to = &tmp
	}

	_gas, err := ParseUint64(gas)
	if err != nil {
		return err
	}

	_gasPrice, err := ParseBigInt(gasPrice)
	if err != nil {
		return err
	}

	_value, err := ParseBigInt(value)
	if err != nil {
		return err
	}

	_data := common.FromHex(data)

	result, err := cli.CallContract(context.Background(), ethereum.CallMsg{
		From:     _from,
		To:       _to,
		Gas:      _gas,
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
