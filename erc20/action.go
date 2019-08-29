package erc20

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	cli "gopkg.in/urfave/cli.v1"
)

var (
	ErrInsufficientArguments = fmt.Errorf("insufficient arguments")
)

func TotalSupplyAction(c *cli.Context) error {
	cli, err := ethclient.DialContext(context.Background(), c.Parent().String(FlagRpcAddr.Name))
	if err != nil {
		return err
	}
	defer cli.Close()
	erc20, err := NewErc20(common.HexToAddress(c.Parent().String(FlagErc20Addr.Name)), cli)
	if err != nil {
		return err
	}
	supply, err := erc20.TotalSupply(nil)
	if err != nil {
		return err
	}
	fmt.Println(supply)
	return nil
}

func BalanceOfAction(c *cli.Context) error {
	if c.NArg() < 1 {
		return ErrInsufficientArguments
	}
	cli, err := ethclient.DialContext(context.Background(), c.Parent().String(FlagRpcAddr.Name))
	if err != nil {
		return err
	}
	defer cli.Close()
	erc20, err := NewErc20(common.HexToAddress(c.Parent().String(FlagErc20Addr.Name)), cli)
	if err != nil {
		return err
	}
	balance, err := erc20.BalanceOf(nil, common.HexToAddress(c.Args()[0]))
	if err != nil {
		return err
	}
	fmt.Println(balance)
	return nil
}

func AllowanceAction(c *cli.Context) error {
	if c.NArg() < 2 {
		return ErrInsufficientArguments
	}
	cli, err := ethclient.DialContext(context.Background(), c.Parent().String(FlagRpcAddr.Name))
	if err != nil {
		return err
	}
	defer cli.Close()
	erc20, err := NewErc20(common.HexToAddress(c.Parent().String(FlagErc20Addr.Name)), cli)
	if err != nil {
		return err
	}
	allowance, err := erc20.Allowance(nil, common.HexToAddress(c.Args()[0]), common.HexToAddress(c.Args()[1]))
	if err != nil {
		return err
	}
	fmt.Println(allowance)
	return nil
}

func TransferAction(c *cli.Context) error {
	if c.NArg() < 3 {
		return ErrInsufficientArguments
	}
	cli, err := ethclient.DialContext(context.Background(), c.Parent().String(FlagRpcAddr.Name))
	if err != nil {
		return err
	}
	defer cli.Close()
	erc20, err := NewErc20(common.HexToAddress(c.Parent().String(FlagErc20Addr.Name)), cli)
	if err != nil {
		return err
	}

	keyStore := keystore.NewKeyStore(c.Parent().String(FlagKeystore.Name), keystore.StandardScryptN, keystore.StandardScryptP)

	passphrase, err := ReadPassphrase(c)
	if err != nil {
		return err
	}

	account := accounts.Account{Address: common.HexToAddress(c.Args()[0])}

	err = keyStore.Unlock(account, passphrase)
	if err != nil {
		return err
	}
	defer keyStore.Lock(account.Address)

	auth, err := bind.NewKeyStoreTransactor(keyStore, account)
	if err != nil {
		return err
	}

	value, err := ParseBigInt(c.Args()[2])
	if err != nil {
		return err
	}

	tx, err := erc20.Transfer(auth, common.HexToAddress(c.Args()[1]), value)
	if err != nil {
		return err
	}
	fmt.Println(tx.Hash().String())
	return nil
}

func TransferFromAction(c *cli.Context) error {
	if c.NArg() < 4 {
		return ErrInsufficientArguments
	}
	cli, err := ethclient.DialContext(context.Background(), c.Parent().String(FlagRpcAddr.Name))
	if err != nil {
		return err
	}
	defer cli.Close()
	erc20, err := NewErc20(common.HexToAddress(c.Parent().String(FlagErc20Addr.Name)), cli)
	if err != nil {
		return err
	}

	keyStore := keystore.NewKeyStore(c.Parent().String(FlagKeystore.Name), keystore.StandardScryptN, keystore.StandardScryptP)

	passphrase, err := ReadPassphrase(c)
	if err != nil {
		return err
	}

	account := accounts.Account{Address: common.HexToAddress(c.Args()[0])}

	err = keyStore.Unlock(account, passphrase)
	if err != nil {
		return err
	}
	defer keyStore.Lock(account.Address)

	auth, err := bind.NewKeyStoreTransactor(keyStore, account)
	if err != nil {
		return err
	}

	value, err := ParseBigInt(c.Args()[3])
	if err != nil {
		return err
	}

	tx, err := erc20.TransferFrom(auth, common.HexToAddress(c.Args()[1]), common.HexToAddress(c.Args()[2]), value)
	if err != nil {
		return err
	}
	fmt.Println(tx.Hash().String())
	return nil
}

func ApproveAction(c *cli.Context) error {
	if c.NArg() < 3 {
		return ErrInsufficientArguments
	}
	cli, err := ethclient.DialContext(context.Background(), c.Parent().String(FlagRpcAddr.Name))
	if err != nil {
		return err
	}
	defer cli.Close()
	erc20, err := NewErc20(common.HexToAddress(c.Parent().String(FlagErc20Addr.Name)), cli)
	if err != nil {
		return err
	}

	keyStore := keystore.NewKeyStore(c.Parent().String(FlagKeystore.Name), keystore.StandardScryptN, keystore.StandardScryptP)

	passphrase, err := ReadPassphrase(c)
	if err != nil {
		return err
	}

	account := accounts.Account{Address: common.HexToAddress(c.Args()[0])}

	err = keyStore.Unlock(account, passphrase)
	if err != nil {
		return err
	}
	defer keyStore.Lock(account.Address)

	auth, err := bind.NewKeyStoreTransactor(keyStore, account)
	if err != nil {
		return err
	}

	value, err := ParseBigInt(c.Args()[2])
	if err != nil {
		return err
	}

	tx, err := erc20.Approve(auth, common.HexToAddress(c.Args()[1]), value)
	if err != nil {
		return err
	}
	fmt.Println(tx.Hash().String())
	return nil
}

func PausedAction(c *cli.Context) error {
	cli, err := ethclient.DialContext(context.Background(), c.Parent().String(FlagRpcAddr.Name))
	if err != nil {
		return err
	}
	defer cli.Close()
	erc20, err := NewErc20(common.HexToAddress(c.Parent().String(FlagErc20Addr.Name)), cli)
	if err != nil {
		return err
	}
	paused, err := erc20.Paused(nil)
	if err != nil {
		return err
	}
	fmt.Println(paused)
	return nil
}

func IsPauserAction(c *cli.Context) error {
	if c.NArg() < 1 {
		return ErrInsufficientArguments
	}
	cli, err := ethclient.DialContext(context.Background(), c.Parent().String(FlagRpcAddr.Name))
	if err != nil {
		return err
	}
	defer cli.Close()
	erc20, err := NewErc20(common.HexToAddress(c.Parent().String(FlagErc20Addr.Name)), cli)
	if err != nil {
		return err
	}
	paused, err := erc20.IsPauser(nil, common.HexToAddress(c.Args()[0]))
	if err != nil {
		return err
	}
	fmt.Println(paused)
	return nil
}

func AddPauserAction(c *cli.Context) error {
	if c.NArg() < 2 {
		return ErrInsufficientArguments
	}
	cli, err := ethclient.DialContext(context.Background(), c.Parent().String(FlagRpcAddr.Name))
	if err != nil {
		return err
	}
	defer cli.Close()
	erc20, err := NewErc20(common.HexToAddress(c.Parent().String(FlagErc20Addr.Name)), cli)
	if err != nil {
		return err
	}

	keyStore := keystore.NewKeyStore(c.Parent().String(FlagKeystore.Name), keystore.StandardScryptN, keystore.StandardScryptP)

	passphrase, err := ReadPassphrase(c)
	if err != nil {
		return err
	}

	account := accounts.Account{Address: common.HexToAddress(c.Args()[0])}

	err = keyStore.Unlock(account, passphrase)
	if err != nil {
		return err
	}
	defer keyStore.Lock(account.Address)

	auth, err := bind.NewKeyStoreTransactor(keyStore, account)
	if err != nil {
		return err
	}

	tx, err := erc20.AddPauser(auth, common.HexToAddress(c.Args()[1]))
	if err != nil {
		return err
	}
	fmt.Println(tx.Hash().String())
	return nil
}

func RenouncePauserAction(c *cli.Context) error {
	if c.NArg() < 1 {
		return ErrInsufficientArguments
	}
	cli, err := ethclient.DialContext(context.Background(), c.Parent().String(FlagRpcAddr.Name))
	if err != nil {
		return err
	}
	defer cli.Close()
	erc20, err := NewErc20(common.HexToAddress(c.Parent().String(FlagErc20Addr.Name)), cli)
	if err != nil {
		return err
	}

	keyStore := keystore.NewKeyStore(c.Parent().String(FlagKeystore.Name), keystore.StandardScryptN, keystore.StandardScryptP)

	passphrase, err := ReadPassphrase(c)
	if err != nil {
		return err
	}

	account := accounts.Account{Address: common.HexToAddress(c.Args()[0])}

	err = keyStore.Unlock(account, passphrase)
	if err != nil {
		return err
	}
	defer keyStore.Lock(account.Address)

	auth, err := bind.NewKeyStoreTransactor(keyStore, account)
	if err != nil {
		return err
	}

	tx, err := erc20.RenouncePauser(auth)
	if err != nil {
		return err
	}
	fmt.Println(tx.Hash().String())
	return nil
}

func PauseAction(c *cli.Context) error {
	if c.NArg() < 1 {
		return ErrInsufficientArguments
	}
	cli, err := ethclient.DialContext(context.Background(), c.Parent().String(FlagRpcAddr.Name))
	if err != nil {
		return err
	}
	defer cli.Close()
	erc20, err := NewErc20(common.HexToAddress(c.Parent().String(FlagErc20Addr.Name)), cli)
	if err != nil {
		return err
	}

	keyStore := keystore.NewKeyStore(c.Parent().String(FlagKeystore.Name), keystore.StandardScryptN, keystore.StandardScryptP)

	passphrase, err := ReadPassphrase(c)
	if err != nil {
		return err
	}

	account := accounts.Account{Address: common.HexToAddress(c.Args()[0])}

	err = keyStore.Unlock(account, passphrase)
	if err != nil {
		return err
	}
	defer keyStore.Lock(account.Address)

	auth, err := bind.NewKeyStoreTransactor(keyStore, account)
	if err != nil {
		return err
	}

	tx, err := erc20.Pause(auth)
	if err != nil {
		return err
	}
	fmt.Println(tx.Hash().String())
	return nil
}

func UnpauseAction(c *cli.Context) error {
	if c.NArg() < 1 {
		return ErrInsufficientArguments
	}
	cli, err := ethclient.DialContext(context.Background(), c.Parent().String(FlagRpcAddr.Name))
	if err != nil {
		return err
	}
	defer cli.Close()
	erc20, err := NewErc20(common.HexToAddress(c.Parent().String(FlagErc20Addr.Name)), cli)
	if err != nil {
		return err
	}

	keyStore := keystore.NewKeyStore(c.Parent().String(FlagKeystore.Name), keystore.StandardScryptN, keystore.StandardScryptP)

	passphrase, err := ReadPassphrase(c)
	if err != nil {
		return err
	}

	account := accounts.Account{Address: common.HexToAddress(c.Args()[0])}

	err = keyStore.Unlock(account, passphrase)
	if err != nil {
		return err
	}
	defer keyStore.Lock(account.Address)

	auth, err := bind.NewKeyStoreTransactor(keyStore, account)
	if err != nil {
		return err
	}

	tx, err := erc20.Unpause(auth)
	if err != nil {
		return err
	}
	fmt.Println(tx.Hash().String())
	return nil
}

func IsMinterAction(c *cli.Context) error {
	if c.NArg() < 1 {
		return ErrInsufficientArguments
	}
	cli, err := ethclient.DialContext(context.Background(), c.Parent().String(FlagRpcAddr.Name))
	if err != nil {
		return err
	}
	defer cli.Close()
	erc20, err := NewErc20(common.HexToAddress(c.Parent().String(FlagErc20Addr.Name)), cli)
	if err != nil {
		return err
	}
	paused, err := erc20.IsMinter(nil, common.HexToAddress(c.Args()[0]))
	if err != nil {
		return err
	}
	fmt.Println(paused)
	return nil
}

func MintAction(c *cli.Context) error {
	if c.NArg() < 3 {
		return ErrInsufficientArguments
	}
	cli, err := ethclient.DialContext(context.Background(), c.Parent().String(FlagRpcAddr.Name))
	if err != nil {
		return err
	}
	defer cli.Close()
	erc20, err := NewErc20(common.HexToAddress(c.Parent().String(FlagErc20Addr.Name)), cli)
	if err != nil {
		return err
	}

	keyStore := keystore.NewKeyStore(c.Parent().String(FlagKeystore.Name), keystore.StandardScryptN, keystore.StandardScryptP)

	passphrase, err := ReadPassphrase(c)
	if err != nil {
		return err
	}

	account := accounts.Account{Address: common.HexToAddress(c.Args()[0])}

	err = keyStore.Unlock(account, passphrase)
	if err != nil {
		return err
	}
	defer keyStore.Lock(account.Address)

	auth, err := bind.NewKeyStoreTransactor(keyStore, account)
	if err != nil {
		return err
	}

	value, err := ParseBigInt(c.Args()[2])
	if err != nil {
		return err
	}
	tx, err := erc20.Mint(auth, common.HexToAddress(c.Args()[1]), value)
	if err != nil {
		return err
	}
	fmt.Println(tx.Hash().String())
	return nil
}

func AddMinterAction(c *cli.Context) error {
	if c.NArg() < 2 {
		return ErrInsufficientArguments
	}
	cli, err := ethclient.DialContext(context.Background(), c.Parent().String(FlagRpcAddr.Name))
	if err != nil {
		return err
	}
	defer cli.Close()
	erc20, err := NewErc20(common.HexToAddress(c.Parent().String(FlagErc20Addr.Name)), cli)
	if err != nil {
		return err
	}

	keyStore := keystore.NewKeyStore(c.Parent().String(FlagKeystore.Name), keystore.StandardScryptN, keystore.StandardScryptP)

	passphrase, err := ReadPassphrase(c)
	if err != nil {
		return err
	}

	account := accounts.Account{Address: common.HexToAddress(c.Args()[0])}

	err = keyStore.Unlock(account, passphrase)
	if err != nil {
		return err
	}
	defer keyStore.Lock(account.Address)

	auth, err := bind.NewKeyStoreTransactor(keyStore, account)
	if err != nil {
		return err
	}

	tx, err := erc20.AddMinter(auth, common.HexToAddress(c.Args()[1]))
	if err != nil {
		return err
	}
	fmt.Println(tx.Hash().String())
	return nil
}

func RenounceMinterAction(c *cli.Context) error {
	if c.NArg() < 1 {
		return ErrInsufficientArguments
	}
	cli, err := ethclient.DialContext(context.Background(), c.Parent().String(FlagRpcAddr.Name))
	if err != nil {
		return err
	}
	defer cli.Close()
	erc20, err := NewErc20(common.HexToAddress(c.Parent().String(FlagErc20Addr.Name)), cli)
	if err != nil {
		return err
	}

	keyStore := keystore.NewKeyStore(c.Parent().String(FlagKeystore.Name), keystore.StandardScryptN, keystore.StandardScryptP)

	passphrase, err := ReadPassphrase(c)
	if err != nil {
		return err
	}

	account := accounts.Account{Address: common.HexToAddress(c.Args()[0])}

	err = keyStore.Unlock(account, passphrase)
	if err != nil {
		return err
	}
	defer keyStore.Lock(account.Address)

	auth, err := bind.NewKeyStoreTransactor(keyStore, account)
	if err != nil {
		return err
	}

	tx, err := erc20.RenounceMinter(auth)
	if err != nil {
		return err
	}
	fmt.Println(tx.Hash().String())
	return nil
}

func BurnAction(c *cli.Context) error {
	if c.NArg() < 2 {
		return ErrInsufficientArguments
	}
	cli, err := ethclient.DialContext(context.Background(), c.Parent().String(FlagRpcAddr.Name))
	if err != nil {
		return err
	}
	defer cli.Close()
	erc20, err := NewErc20(common.HexToAddress(c.Parent().String(FlagErc20Addr.Name)), cli)
	if err != nil {
		return err
	}

	keyStore := keystore.NewKeyStore(c.Parent().String(FlagKeystore.Name), keystore.StandardScryptN, keystore.StandardScryptP)

	passphrase, err := ReadPassphrase(c)
	if err != nil {
		return err
	}

	account := accounts.Account{Address: common.HexToAddress(c.Args()[0])}

	err = keyStore.Unlock(account, passphrase)
	if err != nil {
		return err
	}
	defer keyStore.Lock(account.Address)

	auth, err := bind.NewKeyStoreTransactor(keyStore, account)
	if err != nil {
		return err
	}

	value, err := ParseBigInt(c.Args()[1])
	if err != nil {
		return err
	}

	tx, err := erc20.Burn(auth, value)
	if err != nil {
		return err
	}
	fmt.Println(tx.Hash().String())
	return nil
}

func BurnFromAction(c *cli.Context) error {
	if c.NArg() < 3 {
		return ErrInsufficientArguments
	}
	cli, err := ethclient.DialContext(context.Background(), c.Parent().String(FlagRpcAddr.Name))
	if err != nil {
		return err
	}
	defer cli.Close()
	erc20, err := NewErc20(common.HexToAddress(c.Parent().String(FlagErc20Addr.Name)), cli)
	if err != nil {
		return err
	}

	keyStore := keystore.NewKeyStore(c.Parent().String(FlagKeystore.Name), keystore.StandardScryptN, keystore.StandardScryptP)

	passphrase, err := ReadPassphrase(c)
	if err != nil {
		return err
	}

	account := accounts.Account{Address: common.HexToAddress(c.Args()[0])}

	err = keyStore.Unlock(account, passphrase)
	if err != nil {
		return err
	}
	defer keyStore.Lock(account.Address)

	auth, err := bind.NewKeyStoreTransactor(keyStore, account)
	if err != nil {
		return err
	}

	value, err := ParseBigInt(c.Args()[2])
	if err != nil {
		return err
	}

	tx, err := erc20.BurnFrom(auth, common.HexToAddress(c.Args()[1]), value)
	if err != nil {
		return err
	}
	fmt.Println(tx.Hash().String())
	return nil
}
