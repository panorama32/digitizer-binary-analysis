package digitizerdata

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type Data struct {
	DirPath string
	Buf []byte
	Reader *bytes.Reader
}

func New(path string) *Data {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	filepaths := []string{}
	for _, file := range files {
		if file.IsDir() {
			panic("Invalid path")
		}
		filename := file.Name()
		pos := strings.LastIndex(filename, ".")
		if filename[pos:] == ".raw" {
			filepaths = append(filepaths, path + "/" + filename)
		}
	}

	sort.Slice(filepaths, func(i, j int) bool {
		filename_i, filename_j := filepaths[i], filepaths[j]
		start_i, start_j := strings.LastIndex(filename_i, "_"), strings.LastIndex(filename_j, "_")
		end_i, end_j := strings.LastIndex(filename_i, "."), strings.LastIndex(filename_j, ".")
		num_i, _ := strconv.Atoi(filename_i[(start_i+1):end_i])
		num_j, _ := strconv.Atoi(filename_j[(start_j+1):end_j])
		return num_i < num_j
	})

	data := []byte{}
	for _, raw := range filepaths {
		buf, _ := ioutil.ReadFile(raw)
		data = append(data, buf...)
	}

	r := bytes.NewReader(data)

	return &Data{
		DirPath: path,
		Buf: data,
		Reader: r,
	}
}

// read n bytes from r
func (b *Data) Read(n int) ([]byte, error) {
	buf := make([]byte, n)
	count, err := b.Reader.Read(buf)
	if err != nil {
		return []byte{}, err
	}
	if n != count {
		return []byte{}, errors.New("cannot read n bytes")
	}
	return buf, nil
}

// read 2 bytes as int16
func (b *Data) ReadInt16() (int, error) {
	buf, err := b.Read(2)
	if err != nil {
		return 0, err
	}
	var s string
	for _, x := range buf {
		s += fmt.Sprintf("%02x", x)
	}
	i, err := strconv.ParseInt(s, 16, 16)
	if err != nil {
		return 0, err
	}
	return int(i), nil
	// return int(binary.BigEndian.Uint32(buf)), nil
}
