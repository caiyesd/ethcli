package account

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
)

func NewAccount(ksPath string) error {
	keyStore := keystore.NewKeyStore(ksPath, keystore.StandardScryptN, keystore.StandardScryptP)
	passphrase, err := ReadPassphrase("input passphrase: ")
	if err != nil {
		return err
	}
	passphrase2, err := ReadPassphrase("input again: ")
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

func ListAccount(ksPath string, index string) error {
	keyStore := keystore.NewKeyStore(ksPath, keystore.StandardScryptN, keystore.StandardScryptP)
	if index == "" {
		for _, account := range keyStore.Accounts() {
			fmt.Println(account.Address.String())
		}
	} else {
		_index, err := ParseUint64(index)
		if err != nil {
			return err
		}
		accounts := keyStore.Accounts()
		if _index >= uint64(len(accounts)) {
			return fmt.Errorf("invalid index %s", index)
		}

		fmt.Println(accounts[_index].Address.String())
		return nil
	}
	return nil
}

func DeleteAccount(ksPath, account string) error {
	keyStore := keystore.NewKeyStore(ksPath, keystore.StandardScryptN, keystore.StandardScryptP)
	_account, err := indexOrHexToAddress(keyStore, account)
	if err != nil {
		return err
	}

	passphrase, err := ReadPassphrase("passphrase: ")
	if err != nil {
		return err
	}
	return keyStore.Delete(accounts.Account{Address: _account}, passphrase)
}

func UpdateAccount(ksPath, account string) error {
	keyStore := keystore.NewKeyStore(ksPath, keystore.StandardScryptN, keystore.StandardScryptP)
	_account, err := indexOrHexToAddress(keyStore, account)
	if err != nil {
		return err
	}

	passphrase, err := ReadPassphrase("passphrase: ")
	if err != nil {
		return err
	}
	passphrase2, err := ReadPassphrase("input new passphrase: ")
	if err != nil {
		return err
	}
	passphrase3, err := ReadPassphrase("intput again: ")
	if err != nil {
		return err
	}
	if passphrase3 != passphrase2 {
		return errors.New("does not match")
	}

	return keyStore.Update(accounts.Account{Address: _account}, passphrase, passphrase2)
}

func indexOrHexToAddress(keyStore *keystore.KeyStore, account string) (common.Address, error) {
	if strings.HasPrefix(account, "0x") || strings.HasPrefix(account, "0X") {
		return common.HexToAddress(account), nil
	} else {
		_index, err := ParseUint64(account)
		if err != nil {
			return common.Address{}, err
		}
		accounts := keyStore.Accounts()
		if _index >= uint64(len(accounts)) {
			return common.Address{}, fmt.Errorf("invalid index %s", account)
		}
		return accounts[_index].Address, nil
	}
}
