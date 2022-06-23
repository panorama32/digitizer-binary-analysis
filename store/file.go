package store

import (
	"os"

	"github.com/panorama32/digitizer-binary-analysis/domain/entity"
	"github.com/panorama32/digitizer-binary-analysis/domain/repository"
)

type FileRepository struct{}

func NewFileRepository() repository.FileRepository {
	return &FileRepository{}
}

func (r *FileRepository) Create(filename string, file *entity.File) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, line := range file.Lines {
		_, err := f.WriteString(line)
		// fmt.Fprint()の場合
		// _, err := fmt.Fprint(file, line)
		if err != nil {
			return err
		}
	}
	return nil
}
