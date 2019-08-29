package utils

import (
	"context"
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ExecutorHandler func(*ethclient.Client, *bind.TransactOpts) error

type Executor interface {
	Execute(handler ExecutorHandler) error
}

// ------------------------------------------------------

type ChainExecutor struct {
	rpcUrl string
}

func NewChainExecutor(rpcUrl string) Executor {
	return &ChainExecutor{rpcUrl}
}

func (e *ChainExecutor) Execute(handler ExecutorHandler) error {
	cli, err := ethclient.DialContext(context.Background(), e.rpcUrl)
	if err != nil {
		return err
	}
	defer cli.Close()
	return handler(cli, nil)
}

// ------------------------------------------------------

type ContractExecutorWithKeystore struct {
	ChainExecutor
	keystoreDir string
	account     common.Address
	passphrase  string
}

func NewContractExecutorWithKeystore(rpcUrl string, keystoreDir string, account string, passphrase string) Executor {
	return &ContractExecutorWithKeystore{ChainExecutor{rpcUrl}, keystoreDir, common.HexToAddress(account), passphrase}
}

func (e *ContractExecutorWithKeystore) Execute(handler ExecutorHandler) error {
	h := func(cli *ethclient.Client, auth *bind.TransactOpts) error {
		tmp := accounts.Account{Address: e.account}
		keyStore := keystore.NewKeyStore(e.keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP)
		err := keyStore.Unlock(tmp, e.passphrase)
		if err != nil {
			return err
		}
		defer keyStore.Lock(e.account)
		return handler(cli, auth)
	}
	return e.ChainExecutor.Execute(h)
}

// ------------------------------------------------------

type ContractExecutorWithPrivateKey struct {
	ChainExecutor
	privKey *ecdsa.PrivateKey
}

func NewContractExecutorWithPrivateKey(rpcUrl string, privateKey []byte) Executor {
	key, err := crypto.ToECDSA(privateKey)
	if err != nil {
		panic(err)
	}
	return &ContractExecutorWithPrivateKey{ChainExecutor{rpcUrl}, key}
}

func (e *ContractExecutorWithPrivateKey) Execute(handler ExecutorHandler) error {
	h := func(cli *ethclient.Client, auth *bind.TransactOpts) error {
		return handler(cli, bind.NewKeyedTransactor(e.privKey))
	}
	return e.ChainExecutor.Execute(h)
}
