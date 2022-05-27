package service

import (
	"os"

	"github.com/gocarina/gocsv"
)

//go:generate mockery --name CSV --filename csv.go --outpkg mocks --structname CSVMock --disable-version-string
type CSV interface {
	UnmarshalFile(in *os.File, out interface{}) error
	MarshalFile(in interface{}, file *os.File) (err error)
}

type csv struct {
}

func NewCSV() *csv {
	return &csv{}
}

func (csv csv) UnmarshalFile(in *os.File, out interface{}) error {
	return gocsv.UnmarshalFile(in, out)
}

func (csv csv) MarshalFile(in interface{}, file *os.File) (err error) {
	return gocsv.MarshalFile(in, file)
}
