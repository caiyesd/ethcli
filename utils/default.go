package utils

import (
	"os"
	"path/filepath"
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

func DefaultPassword() string {
	return filepath.Join(DefaultHome(), ".ethereum", "password")
}
