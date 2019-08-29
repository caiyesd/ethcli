package crypto

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

func SignHash(ksPath string, account, passphrase string, hash string) error {
	keyStore := keystore.NewKeyStore(ksPath, keystore.StandardScryptN, keystore.StandardScryptP)
	h := common.HexToHash(hash)
	sig, err := keyStore.SignHashWithPassphrase(accounts.Account{Address: common.HexToAddress(account)}, passphrase, h[:])
	if err != nil {
		return err
	}
	fmt.Println(common.ToHex(sig))
	return nil
}

func SignTx(ksPath string, account, passphrase string, rawTxStr string, chainId string) error {
	keyStore := keystore.NewKeyStore(ksPath, keystore.StandardScryptN, keystore.StandardScryptP)
	tx := new(types.Transaction)
	err := tx.DecodeRLP(rlp.NewStream(bytes.NewBuffer(common.FromHex(rawTxStr)), 0))
	if err != nil {
		return err
	}
	var _chainId *big.Int = nil
	if chainId != "" {
		tmp, err := ParseBigInt(chainId)
		if err != nil {
			return err
		}
		_chainId = tmp
	}

	tx, err = keyStore.SignTxWithPassphrase(accounts.Account{Address: common.HexToAddress(account)}, passphrase, tx, _chainId)
	if err != nil {
		return err
	}
	buffer := bytes.NewBuffer(nil)
	err = tx.EncodeRLP(buffer)
	if err != nil {
		return err
	}
	fmt.Println(common.ToHex(buffer.Bytes()))
	return nil
}

func Verify(pub, hash, sig string) error {
	h := common.HexToHash(hash)
	p := common.FromHex(pub)
	s := common.FromHex(sig)
	ok := crypto.VerifySignature(p, h[:], s)
	fmt.Println(ok)
	return nil
}

func Ecrecover(hash string, sig string) error {
	h := common.HexToHash(hash)
	s := common.FromHex(sig)
	pub, err := crypto.SigToPub(h[:], s)
	if err != nil {
		return err
	}
	fmt.Println(crypto.PubkeyToAddress(*pub).String())
	return nil
}

func PackArguments(args []string) error {
	arguments := abi.Arguments{}
	values := []interface{}{}
	for i := 0; i < len(args)/2; i++ {
		_type, err := abi.NewType(args[i*2], nil)
		if err != nil {
			return err
		}
		arguments = append(arguments, abi.Argument{Type: _type})
		switch args[i*2] {
		case "address":
			values = append(values, common.HexToAddress(args[i*2+1]))
			break
		case "bytes":
			values = append(values, common.FromHex(args[i*2+1]))
			break
		case "bytes32":
			values = append(values, common.HexToHash(args[i*2+1]))
			break
		case "uint64":
			value, err := ParseUint64(args[i*2+1])
			if err != nil {
				return err
			}
			values = append(values, value)
			break
		case "uint256":
			value, err := ParseBigInt(args[i*2+1])
			if err != nil {
				return err
			}
			values = append(values, value)
			break
		case "string":
			values = append(values, args[i*2+1])
			break
		default:
			return fmt.Errorf("unsupported type %s", args[i*2])
		}
	}
	data, err := arguments.Pack(values...)
	if err != nil {
		return err
	}
	fmt.Println(common.ToHex(data))
	return nil
}
