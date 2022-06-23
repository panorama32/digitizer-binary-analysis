package repository

import "github.com/panorama32/digitizer-binary-analysis/domain/entity"

type FileRepository interface {
	Create(filename string, file *entity.File) error
}
