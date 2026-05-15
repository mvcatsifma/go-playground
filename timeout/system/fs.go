package system

import (
	"github.com/spf13/afero"
	"os"
)

func NewAferoFs() *aferoFs {
	return &aferoFs{
		fs: afero.NewOsFs(),
	}
}

type aferoFs struct {
	fs afero.Fs
}

func (a aferoFs) Stat(name string) (os.FileInfo, error) {
	return a.fs.Stat(name)
}
