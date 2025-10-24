package fileadapter

import "github.com/goodylabs/awxhelper/internal/ports"

type fileAdapterMock struct {
	Path string
	Obj  any
}

func NewFileAdapter(obj any) ports.FileAdapter {
	return &fileAdapterMock{
		Obj: obj,
	}
}

func (fa *fileAdapterMock) ReadJSONFile(path string, target any) error {
	target = fa.Obj
	return nil
}

func (fa *fileAdapterMock) WriteJSONFile(path string, data any) error {
	fa.Obj = data
	return nil
}

func (fa *fileAdapterMock) GetAwxConfigPath() string {
	return ""
}
