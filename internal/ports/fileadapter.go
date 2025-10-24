package ports

type FileAdapter interface {
	ReadJSONFile(path string, target any) error
	WriteJSONFile(path string, data any) error
}
