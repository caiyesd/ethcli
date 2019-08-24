package crypto

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/rlp"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/accounts"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func NewAccount(ksPath string) error {
	keyStore := keystore.NewKeyStore(ksPath, keystore.StandardScryptN, keystore.StandardScryptP)
	t := NewTerminalPrompter()
	passphrase, err := t.PromptPassphrase("Please input passphrase: ")
	if err != nil {
		return err
	}
	passphrase2, err := t.PromptPassphrase("Please input passphrase again: ")
	if err != nil {
		return err
	}
	if passphrase != passphrase2 {
		return errors.New("does not match")
	}
	account, err := keyStore.NewAccount(passphrase)
	if err != nil {
		return err
	}
	fmt.Println(account.Address.String())
	return nil
}

func ListAccount(ksPath string) error {
	keyStore := keystore.NewKeyStore(ksPath, keystore.StandardScryptN, keystore.StandardScryptP)
	for _, account := range keyStore.Accounts() {
		fmt.Println(account.Address.String())
	}
	return nil
}

func DeleteAccount(ksPath string, account string) error {
	keyStore := keystore.NewKeyStore(ksPath, keystore.StandardScryptN, keystore.StandardScryptP)
	t := NewTerminalPrompter()
	passphrase, err := t.PromptPassphrase("Please input passphrase: ")
	if err != nil {
		return err
	}
	return keyStore.Delete(accounts.Account{Address: common.HexToAddress(account)}, passphrase)
}

func UpdateAccount(ksPath string, account string) error {
	keyStore := keystore.NewKeyStore(ksPath, keystore.StandardScryptN, keystore.StandardScryptP)
	t := NewTerminalPrompter()
	passphrase, err := t.PromptPassphrase("Please input old passphrase: ")
	if err != nil {
		return err
	}
	passphrase2, err := t.PromptPassphrase("Please input new passphrase: ")
	if err != nil {
		return err
	}
	return keyStore.Update(accounts.Account{Address: common.HexToAddress(account)}, passphrase, passphrase2)
}

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
	_chainId, ok := new(big.Int).SetString(chainId, 10)
	if !ok {
		return ErrInvalidAmount
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
