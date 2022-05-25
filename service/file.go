package service

import (
	"os"
)

//go:generate mockery --name File --filename file.go --outpkg mocks --structname FileMock --disable-version-string
type File interface {
	OpenFile(flag int, perm os.FileMode) (*os.File, error)
	Close()
}

type file struct {
	filename    string
	fileHandler *os.File
}

func NewFile(name string) *file {
	return &file{
		filename: name,
	}
}

func (f file) OpenFile(flag int, perm os.FileMode) (*os.File, error) {
	var err error
	f.fileHandler, err = os.OpenFile(f.filename, flag, perm)
	return f.fileHandler, err
}

func (f file) Close() {
	f.fileHandler.Close()
}
