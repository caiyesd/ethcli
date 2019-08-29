package crypto

import (
	"github.com/caiyesd/ethcli/utils"
	cli "gopkg.in/urfave/cli.v1"
)

var (
	FlagKeystore = cli.StringFlag{
		Name:  "keystore",
		Usage: "keystore directory",
		Value: DefaultKeystore(),
	}
	FlagPassword = cli.StringFlag{
		Name:  "password",
		Usage: "password file",
		Value: DefaultPassword(),
	}
)

func ParseFlagKeystore(c *cli.Context) string {
	return c.Parent().String(FlagKeystore.Name)
}

func ParseFlagPassword(c *cli.Context) string {
	return c.Parent().String(FlagPassword.Name)
}

func ReadPassphrase(c *cli.Context) (string, error) {
	passwordFile := ParseFlagPassword(c)
	if passwordFile == "" {
		return utils.ReadPassphraseFromPrompter()
	} else {
		passphrase, err := utils.ReadPassphraseFromFile(passwordFile)
		if err != nil {
			return utils.ReadPassphraseFromPrompter()
		} else {
			return passphrase, nil
		}
	}
}

var ParseBigInt = utils.ParseBigInt
var ParseUint64 = utils.ParseUint64
var DefaultKeystore = utils.DefaultKeystore
var DefaultPassword = utils.DefaultPassword
