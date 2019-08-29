package crypto

import (
	"bytes"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func BuildTx(nonce, to, amount, gasLimit, gasPrice, data string) error {
	_nonce, err := ParseUint64(nonce)
	if err != nil {
		return err
	}
	_amount, err := ParseBigInt(amount)
	if err != nil {
		return err
	}
	_gasLimit, err := ParseUint64(gasLimit)
	if err != nil {
		return err
	}
	_gasPrice, err := ParseBigInt(gasPrice)
	if err != nil {
		return err
	}
	var _data []byte = nil
	if data != "" {
		_data = common.FromHex(data)
	}
	tx := types.NewTransaction(_nonce, common.HexToAddress(to), _amount, _gasLimit, _gasPrice, _data)
	buffer := bytes.NewBuffer(nil)
	err = tx.EncodeRLP(buffer)
	if err != nil {
		return err
	}
	fmt.Println(common.ToHex(buffer.Bytes()))
	return nil
}
