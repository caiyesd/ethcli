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

type ChainExecutor struct {
	rpcUrl string
}

func (e *ChainExecutor) Execute(handler ExecutorHandler) error {
	cli, err := ethclient.DialContext(context.Background(), e.rpcUrl)
	if err != nil {
		return err
	}
	defer cli.Close()
	return handler(cli, nil)
}

type ContractExecutorWithKeystore struct {
	ChainExecutor
	keystoreDir string
	account     common.Address
	passphrase  string
}

type ContractExecutor struct {
	ChainExecutor
	privKey *ecdsa.PrivateKey
}

func NewChainExecutor(rpcUrl string) Executor {
	return &ChainExecutor{rpcUrl}
}

func ExecuteChainRpcWithKeystore(rpcUrl string, keystoreDir string, account common.Address, passphrase string,
	handler func(*ethclient.Client, *bind.TransactOpts) error) error {
	cli, err := ethclient.DialContext(context.Background(), rpcUrl)
	if err != nil {
		return err
	}
	tmp := accounts.Account{Address: account}
	keyStore := keystore.NewKeyStore(keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP)
	err = keyStore.Unlock(tmp, passphrase)
	if err != nil {
		return err
	}
	defer keyStore.Lock(account)

	auth, err := bind.NewKeyStoreTransactor(keyStore, tmp)
	if err != nil {
		return err
	}
	return handler(cli, auth)
}

func ExecuteChainRpcWithPrivateKey(rpcUrl string, sk []byte,
	handler func(*ethclient.Client, *bind.TransactOpts) error) error {
	cli, err := ethclient.DialContext(context.Background(), rpcUrl)
	if err != nil {
		return err
	}

	priv, err := crypto.ToECDSA(sk)
	if err != nil {
		return err
	}

	auth := bind.NewKeyedTransactor(priv)

	return handler(cli, auth)
}
