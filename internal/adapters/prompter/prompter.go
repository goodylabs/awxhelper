package prompter

import (
	"github.com/goodylabs/awxhelper/internal/services/ports"
)

type prompter struct {
	lastIndexes map[string]int
}

func NewPrompter() ports.Prompter {
	return &prompter{
		lastIndexes: make(map[string]int),
	}
}

func (p *prompter) ChooseFromList(options []ports.PrompterItem, label string) (string, error) {
	p.clear()
	return p.runPrompter(options, label)
}

// func (p *prompter) ChooseFromMap(options map[string]string, label string) (string, error) {
// 	keys := make([]string, 0, len(options))
// 	for k := range options {
// 		keys = append(keys, k)
// 	}
// 	optionsPrompterItem := make([]ports.PrompterItem, len(keys))
// 	for i, key := range keys {
// 		optionsPrompterItem[i] = ports.PrompterItem{
// 			Label: key,
// 			Value: key,
// 		}
// 	}
// 	resultKey, err := p.runPrompter(optionsPrompterItem, label)
// 	if err != nil {
// 		return "", err
// 	}
// 	return options[resultKey], nil
// }
