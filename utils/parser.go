package utils

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

func ParseBigInt(str string) (*big.Int, error) {
	if strings.HasPrefix(str, "0x") || strings.HasPrefix(str, "0X") {
		return new(big.Int).SetBytes(common.FromHex(str)), nil
	} else {
		num, ok := new(big.Int).SetString(str, 10)
		if !ok {
			return nil, fmt.Errorf("failed to parse %s as decimal format", str)
		} else {
			return num, nil
		}
	}
}

func ParseUint64(str string) (uint64, error) {
	if strings.HasPrefix(str, "0x") || strings.HasPrefix(str, "0X") {
		return new(big.Int).SetBytes(common.FromHex(str)).Uint64(), nil
	} else {
		return strconv.ParseUint(str, 10, 64)
	}
}
