package ports

type Prompter interface {
	ChooseFromList([]PrompterItem, string) (PrompterItem, error)
	PromptForString(message string) (string, error)
}

type PrompterItem struct {
	Value string
	Label string
}
