package prompter

import (
	"github.com/goodylabs/awxhelper/internal/ports"
	"github.com/manifoldco/promptui"
)

type prompter struct {
	lastIndexes map[string]int
}

func NewPrompter() ports.Prompter {
	return &prompter{
		lastIndexes: make(map[string]int),
	}
}

func (p *prompter) ChooseFromList(options []ports.PrompterItem, label string) (ports.PrompterItem, error) {
	return p.runPrompter(options, label)
}

func (p *prompter) PromptForString(message string) (string, error) {
	prompt := promptui.Prompt{
		Label:   message,
		Pointer: promptui.PipeCursor,
	}
	return prompt.Run()
}
