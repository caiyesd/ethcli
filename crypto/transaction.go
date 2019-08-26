package crypto

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func NewTx(nonce uint64, to string, amount string, gasLimit uint64, gasPrice string, data string) error {
	_amount, ok := new(big.Int).SetString(amount, 10)
	if !ok {
		return fmt.Errorf("failed to parse %s", amount)
	}
	_gasPrice, ok := new(big.Int).SetString(gasPrice, 10)
	if !ok {
		return fmt.Errorf("failed to parse %s", gasPrice)
	}
	var _data []byte = nil
	if data != "" {
		_data = common.FromHex(data)
	}
	tx := types.NewTransaction(nonce, common.HexToAddress(to), _amount, gasLimit, _gasPrice, _data)
	buffer := bytes.NewBuffer(nil)
	err := tx.EncodeRLP(buffer)
	if err != nil {
		return err
	}
	fmt.Println(common.ToHex(buffer.Bytes()))
	return nil
}
