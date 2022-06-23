package entity

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
)

type Graph struct {
	ID   uint
	Plot *plot.Plot
}

func New(id uint, title, xLabel, yLabel string, data plotter.XYs) (*Graph, error) {
	p := plot.New()
	p.Title.Text = title
	p.X.Label.Text = xLabel
	p.Y.Label.Text = yLabel
	if err := plotutil.AddLinePoints(p, data); err != nil {
		return nil, err
	}

	return &Graph{
		ID:   id,
		Plot: p,
	}, nil
}
