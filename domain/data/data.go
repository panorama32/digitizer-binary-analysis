package data

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"

	"github.com/panorama32/digitizer-binary-analysis/domain/rawdata"
	"github.com/panorama32/digitizer-binary-analysis/domain/service"
)

type Data struct {
	DirPath string
	Bytes   []byte
	Reader  *bytes.Reader
}

func New(dirpath string) (*Data, error) {
	files, err := ioutil.ReadDir(dirpath)
	if err != nil {
		return nil, err
	}

	rds := []*rawdata.RawData{}
	for _, file := range files {
		if file.IsDir() {
			return nil, errors.New("invalid dir path")
		}
		filename := file.Name()
		extPos := strings.LastIndex(filename, ".")
		if extPos == -1 {
			return nil, errors.New("extension not found")
		}
		if filename[extPos:] != ".raw" {
			continue
		}
		rd, _ := rawdata.New(dirpath + "/" + filename)
		rds = append(rds, rd)
	}

	sort.Slice(rds, func(i, j int) bool {
		return rds[i].BranchNum < rds[j].BranchNum
	})

	data := service.NewRawDataService(rds).Join().Bytes

	r := bytes.NewReader(data)

	return &Data{
		DirPath: dirpath,
		Bytes:   data,
		Reader:  r,
	}, nil
}

// read n bytes from Data.Reader
func (d *Data) Read(n int) ([]byte, error) {
	buf := make([]byte, n)
	count, err := d.Reader.Read(buf)
	if err != nil {
		return []byte{}, err
	}
	if n != count {
		return []byte{}, errors.New("cannot read n bytes")
	}
	return buf, nil
}

// read 2 bytes from Data.Reader as int16
func (d *Data) ReadInt16() (int, error) {
	buf, err := d.Read(2)
	if err != nil {
		return 0, err
	}
	var s string
	for _, b := range buf {
		s += fmt.Sprintf("%02x", b)
	}
	i, err := strconv.ParseInt(s, 16, 16)
	if err != nil {
		return 0, err
	}
	// return int(binary.BigEndian.Uint32(buf)), nil
	return int(i), nil
}
