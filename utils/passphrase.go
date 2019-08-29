package utils

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func ReadPassphraseFromFile(passwordFile string) (string, error) {
	f, err := os.Open(passwordFile)
	if err != nil {
		return "", err
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	line, err := rd.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", err
	}
	return strings.TrimSpace(line), nil
}

func ReadPassphraseFromPrompter() (string, error) {
	t := NewTerminalPrompter()
	passphrase, err := t.PromptPassphrase("passphrase: ")
	return passphrase, err
}

func ReadPassphraseFromPrompter2(prompt string) (string, error) {
	t := NewTerminalPrompter()
	passphrase, err := t.PromptPassphrase(prompt)
	return passphrase, err
}
