package repository

import "github.com/panorama32/digitizer-binary-analysis/domain/entity"

type GraphRepository interface {
	SaveAsPng(filename string, g *entity.Graph) error
}
