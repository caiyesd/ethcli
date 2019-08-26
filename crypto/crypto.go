package crypto

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

func SignHash(ksPath string, account string, hash string) error {
	keyStore := keystore.NewKeyStore(ksPath, keystore.StandardScryptN, keystore.StandardScryptP)
	t := NewTerminalPrompter()
	passphrase, err := t.PromptPassphrase("Please input passphrase: ")
	if err != nil {
		return err
	}
	h := common.HexToHash(hash)
	sig, err := keyStore.SignHashWithPassphrase(accounts.Account{Address: common.HexToAddress(account)}, passphrase, h[:])
	if err != nil {
		return err
	}
	fmt.Println(common.ToHex(sig))
	return nil
}

func SignTx(ksPath string, account string, rawTxStr string, chainId string) error {
	keyStore := keystore.NewKeyStore(ksPath, keystore.StandardScryptN, keystore.StandardScryptP)
	t := NewTerminalPrompter()
	passphrase, err := t.PromptPassphrase("Please input passphrase: ")
	if err != nil {
		return err
	}
	tx := new(types.Transaction)
	err = tx.DecodeRLP(rlp.NewStream(bytes.NewBuffer(common.FromHex(rawTxStr)), 0))
	if err != nil {
		return err
	}
	var _chainId *big.Int = nil
	if chainId != "" {
		tmp, ok := new(big.Int).SetString(chainId, 10)
		if !ok {
			return fmt.Errorf("failed to parse %s", chainId)
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
