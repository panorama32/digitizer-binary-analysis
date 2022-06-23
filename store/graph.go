package store

import (
	"strconv"

	"github.com/panorama32/digitizer-binary-analysis/domain/entity"
	"github.com/panorama32/digitizer-binary-analysis/domain/repository"
	"gonum.org/v1/plot/vg"
)

type GraphRepository struct{}

func NewGraphRepository() repository.GraphRepository {
	return &GraphRepository{}
}

func (r *GraphRepository) SaveAsPng(filename string, g *entity.Graph) error {
	if filename == "" {
		filename = "output/" + strconv.Itoa(int(g.ID)) + ".png"
	}

	if err := g.Plot.Save(5*vg.Inch, 5*vg.Inch, filename); err != nil {
		return err
	}

	return nil
}
