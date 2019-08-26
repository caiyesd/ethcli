package crypto

import (
	"os"
	"path/filepath"

	cli "gopkg.in/urfave/cli.v1"
)

var (
	FlagKeystore = cli.StringFlag{
		Name:  "ks",
		Usage: "keystore directory",
		Value: DefaultKeystore(),
	}
)

func DefaultHome() string {
	home, err := os.UserHomeDir()
	if err != nil {
		home = ""
	}
	return home
}

func DefaultKeystore() string {
	return filepath.Join(DefaultHome(), ".ethereum", "keystore")
}
