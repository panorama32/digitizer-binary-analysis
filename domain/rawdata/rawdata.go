package rawdata

import (
	"errors"
	"io/ioutil"
	"strconv"
	"strings"
)

type RawData struct {
	BranchNum uint
	FilePath  string
	Bytes     []byte
}

func New(filepath string) (*RawData, error) {
	extPos := strings.LastIndex(filepath, ".")
	if extPos == -1 {
		return nil, errors.New("extension not found")
	}
	if filepath[extPos:] != ".raw" {
		return nil, errors.New("file is not .raw file")
	}

	branchNumPos := strings.LastIndex(filepath, "_")
	branchNum, err := strconv.Atoi(filepath[(branchNumPos + 1):extPos])
	if err != nil {
		return nil, errors.New("invalid branch num")
	}

	buf, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, errors.New("cannot read file")
	}

	return &RawData{
		BranchNum: uint(branchNum),
		FilePath:  filepath,
		Bytes:     buf,
	}, nil
}
