package ports

type PrompterItem struct {
	Value string
	Label string
}

type Prompter interface {
	ChooseFromList([]PrompterItem, string) (string, error)
}
