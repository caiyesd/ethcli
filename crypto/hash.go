package crypto

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func NewHash(data string) error {
	_data := common.FromHex(data)
	h := crypto.Keccak256Hash(_data)
	fmt.Println(h.String())
	return nil
}
