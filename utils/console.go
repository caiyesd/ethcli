package utils

import (
	"fmt"
	"strings"

	"github.com/peterh/liner"
)

// UserPrompter handle console user input interactive
type UserPrompter interface {
	Prompt(prompt string) (string, error)
	PromptPassphrase(prompt string) (passwd string, err error)
	PromptConfirm(prompt string) (bool, error)
	SetHistory(history []string)
	AppendHistory(command string)
	SetWordCompleter(completer liner.WordCompleter)
}

// TerminalPrompter terminal prompter
type TerminalPrompter struct {
	liner     *liner.State
	supported bool
	origMode  liner.ModeApplier
	rawMode   liner.ModeApplier
}

// NewTerminalPrompter create a terminal prompter
func NewTerminalPrompter() *TerminalPrompter {
	p := new(TerminalPrompter)
	// Get the original mode before calling NewLiner.
	origMode, _ := liner.TerminalMode()
	// Turn on liner.
	p.liner = liner.NewLiner()
	rawMode, err := liner.TerminalMode()
	if err != nil || !liner.TerminalSupported() {
		p.supported = false
	} else {
		p.supported = true
		p.origMode = origMode
		p.rawMode = rawMode
		// Switch back to normal mode while we're not prompting.
		origMode.ApplyMode()
	}
	p.liner.SetCtrlCAborts(true)
	p.liner.SetTabCompletionStyle(liner.TabPrints)
	p.liner.SetMultiLineMode(true)
	return p
}

// Prompt shows the prompt and requests text input
// returning the input.
func (p *TerminalPrompter) Prompt(prompt string) (string, error) {
	if p.supported {
		p.rawMode.ApplyMode()
		defer p.origMode.ApplyMode()
	} else {
		fmt.Print(prompt)
		defer fmt.Println()
	}
	return p.liner.Prompt(prompt)
}

// PromptPassphrase shows the prompt and request passphrase text input, the passphrase
// not show, returns the passphrase
func (p *TerminalPrompter) PromptPassphrase(prompt string) (passwd string, err error) {
	if p.supported {
		p.rawMode.ApplyMode()
		defer p.origMode.ApplyMode()
		return p.liner.PasswordPrompt(prompt)
	}

	fmt.Print(prompt)
	passwd, err = p.liner.Prompt("")
	fmt.Println()
	return passwd, err
}

// PromptConfirm shows the prompt to the user and requests a boolean
// choice to be made, returning that choice.
func (p *TerminalPrompter) PromptConfirm(prompt string) (bool, error) {
	input, err := p.Prompt(prompt + " [y/N] ")
	if len(input) > 0 && strings.ToUpper(input[:1]) == "Y" {
		return true, nil
	}
	return false, err
}

// SetHistory sets the history that the prompter will allow
// the user to scroll back to.
func (p *TerminalPrompter) SetHistory(history []string) {
	p.liner.ReadHistory(strings.NewReader(strings.Join(history, "\n")))
}

// AppendHistory appends an entry to the scrollback history.
func (p *TerminalPrompter) AppendHistory(command string) {
	p.liner.AppendHistory(command)
}

// SetWordCompleter sets the completion function that the prompter will call to
// fetch completion candidates when the user presses tab.
func (p *TerminalPrompter) SetWordCompleter(completer liner.WordCompleter) {
	p.liner.SetWordCompleter(completer)
}
