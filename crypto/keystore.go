package crypto

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
)

func NewAccount(ksPath string) error {
	keyStore := keystore.NewKeyStore(ksPath, keystore.StandardScryptN, keystore.StandardScryptP)
	t := NewTerminalPrompter()
	passphrase, err := t.PromptPassphrase("input passphrase: ")
	if err != nil {
		return err
	}
	passphrase2, err := t.PromptPassphrase("input passphrase again: ")
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
	passphrase, err := t.PromptPassphrase("passphrase: ")
	if err != nil {
		return err
	}
	return keyStore.Delete(accounts.Account{Address: common.HexToAddress(account)}, passphrase)
}

func UpdateAccount(ksPath string, account string) error {
	keyStore := keystore.NewKeyStore(ksPath, keystore.StandardScryptN, keystore.StandardScryptP)
	t := NewTerminalPrompter()
	passphrase, err := t.PromptPassphrase("passphrase: ")
	if err != nil {
		return err
	}
	passphrase2, err := t.PromptPassphrase("new passphrase: ")
	if err != nil {
		return err
	}
	passphrase3, err := t.PromptPassphrase("new passphrase again: ")
	if err != nil {
		return err
	}
	if passphrase3 != passphrase2 {
		return errors.New("does not match")
	}
	return keyStore.Update(accounts.Account{Address: common.HexToAddress(account)}, passphrase, passphrase2)
}
