package prompter

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/goodylabs/awxhelper/internal/services/dto"
	"github.com/manifoldco/promptui"
	"golang.org/x/term"
)

func (p *prompter) runPrompter(options []dto.PrompterItem, label string) (dto.PrompterItem, error) {
	sort.Slice(options, func(i, j int) bool {
		return options[i].Label < options[j].Label
	})

	_, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		panic(err)
	}

	hashKey := p.hashOptions(options)
	lastIndex := p.lastIndexes[hashKey]

	prompt := promptui.Select{
		Label:             label,
		Items:             options,
		Size:              height - 3,
		StartInSearchMode: false,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . }}",
			Active:   "▸ {{ .Label | cyan }}",
			Inactive: "  {{ .Label }}",
			Selected: "✔ {{ .Label | green }}",
		},
		Searcher: func(input string, index int) bool {
			option := options[index]
			return strings.Contains(option.Label, input)
		},
		CursorPos: lastIndex,
		Stdout:    noBellWriter{os.Stdout},
	}

	i, _, err := prompt.Run()
	if err != nil {
		return options[0], err
	}

	p.lastIndexes[hashKey] = i
	return options[i], nil
}

func (p *prompter) clear() {
	fmt.Print("\033[H\033[2J")
}

func (p *prompter) hashOptions(options []dto.PrompterItem) string {
	labels := make([]string, len(options))
	for i, opt := range options {
		labels[i] = opt.Label
	}
	hash := sha256.Sum256([]byte(strings.Join(labels, "|")))
	return hex.EncodeToString(hash[:])
}
