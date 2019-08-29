package account

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
)

func ParseFlagKeystore(c *cli.Context) string {
	return c.Parent().String(FlagKeystore.Name)
}

func ReadPassphrase(prompt string) (string, error) {
	return utils.ReadPassphraseFromPrompter2(prompt)
}

var ParseBigInt = utils.ParseBigInt
var ParseUint64 = utils.ParseUint64
var DefaultKeystore = utils.DefaultKeystore
