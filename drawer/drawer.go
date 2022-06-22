package drawer

import (
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func Draw(i int, title, xLabel, yLabel string, data plotter.XYs) error {
	p := plot.New()
	p.Title.Text = title
	p.X.Label.Text = xLabel
	p.Y.Label.Text = yLabel

	if err := plotutil.AddLinePoints(p, data); err != nil {
		return err
	}

	// fileName := "output/" + time.Now().String() + ".png"
	fileName := "output/" + strconv.Itoa(i) + ".png"

	if err := p.Save(5*vg.Inch, 5*vg.Inch, fileName); err != nil {
		return err
	}

	return nil
}
