package usecase

import (
	"os"
	"strconv"

	"github.com/panorama32/digitizer-binary-analysis/domain/data"
	"github.com/panorama32/digitizer-binary-analysis/domain/entity"
	"github.com/panorama32/digitizer-binary-analysis/domain/repository"
	"gonum.org/v1/plot/plotter"
)

type ProcessUseCase struct {
	fileRepo  repository.FileRepository
	graphRepo repository.GraphRepository
}

func NewProcessUseCase(fileRepo repository.FileRepository, graphRepo repository.GraphRepository) *ProcessUseCase {
	return &ProcessUseCase{
		fileRepo:  fileRepo,
		graphRepo: graphRepo,
	}
}

func (u *ProcessUseCase) SetUp() error {
	if _, err := os.Stat("output"); err == nil {
		if err := os.RemoveAll("output"); err != nil {
			return err
		}
	}

	if err := os.Mkdir("output", 0777); err != nil {
		return err
	}

	return nil
}

func (u *ProcessUseCase) ProcessChunk(chunkID int, d *data.Data) error {
	data0 := plotter.XYs{}
	data1 := plotter.XYs{}
	data2 := plotter.XYs{}
	data3 := plotter.XYs{}
	data4 := plotter.XYs{}
	data5 := plotter.XYs{}
	data6 := plotter.XYs{}
	data7 := plotter.XYs{}

	// dispose 2bytes of "5555-aaaa"
	_, err := d.Read(2)
	if err != nil {
		return err
	}
	_, err = d.Read(2)
	if err != nil {
		return err
	}

	for i := 0; i < 8; i++ {
		for k := 0; k < 1024; k++ {
			x, err := d.ReadInt16()
			if err != nil {
				// fmt.Println(err)
				switch i {
				case 0:
					data0 = append(data0, plotter.XY{X: float64(k), Y: float64(0)})
				case 1:
					data1 = append(data1, plotter.XY{X: float64(k), Y: float64(0)})
				case 2:
					data2 = append(data2, plotter.XY{X: float64(k), Y: float64(0)})
				case 3:
					data3 = append(data3, plotter.XY{X: float64(k), Y: float64(0)})
				case 4:
					data4 = append(data4, plotter.XY{X: float64(k), Y: float64(0)})
				case 5:
					data5 = append(data5, plotter.XY{X: float64(k), Y: float64(0)})
				case 6:
					data6 = append(data6, plotter.XY{X: float64(k), Y: float64(0)})
				case 7:
					data7 = append(data7, plotter.XY{X: float64(k), Y: float64(0)})
				}
				continue
			}
			// if k == 0 {
			// 	switch i {
			// 	case 0:
			// 		data0 = append(data0, plotter.XY{X: float64(k), Y: float64(0)})
			// 	case 1:
			// 		data1 = append(data1, plotter.XY{X: float64(k), Y: float64(0)})
			// 	case 2:
			// 		data2 = append(data2, plotter.XY{X: float64(k), Y: float64(0)})
			// 	case 3:
			// 		data3 = append(data3, plotter.XY{X: float64(k), Y: float64(0)})
			// 	case 4:
			// 		data4 = append(data4, plotter.XY{X: float64(k), Y: float64(0)})
			// 	case 5:
			// 		data5 = append(data5, plotter.XY{X: float64(k), Y: float64(0)})
			// 	case 6:
			// 		data6 = append(data6, plotter.XY{X: float64(k), Y: float64(0)})
			// 	case 7:
			// 		data7 = append(data7, plotter.XY{X: float64(k), Y: float64(0)})
			// 	}
			// }
			switch i {
			case 0:
				data0 = append(data0, plotter.XY{X: float64(k), Y: float64(x)})
			case 1:
				data1 = append(data1, plotter.XY{X: float64(k), Y: float64(x)})
			case 2:
				data2 = append(data2, plotter.XY{X: float64(k), Y: float64(x)})
			case 3:
				data3 = append(data3, plotter.XY{X: float64(k), Y: float64(x)})
			case 4:
				data4 = append(data4, plotter.XY{X: float64(k), Y: float64(x)})
			case 5:
				data5 = append(data5, plotter.XY{X: float64(k), Y: float64(x)})
			case 6:
				data6 = append(data6, plotter.XY{X: float64(k), Y: float64(x)})
			case 7:
				data7 = append(data7, plotter.XY{X: float64(k), Y: float64(x)})
			}
			// if k != 0 {
			// 	switch i {
			// 	case 0:
			// 		data0 = append(data0, plotter.XY{X: float64(k), Y: float64(x)})
			// 	case 1:
			// 		data1 = append(data1, plotter.XY{X: float64(k), Y: float64(x)})
			// 	case 2:
			// 		data2 = append(data2, plotter.XY{X: float64(k), Y: float64(x)})
			// 	case 3:
			// 		data3 = append(data3, plotter.XY{X: float64(k), Y: float64(x)})
			// 	case 4:
			// 		data4 = append(data4, plotter.XY{X: float64(k), Y: float64(x)})
			// 	case 5:
			// 		data5 = append(data5, plotter.XY{X: float64(k), Y: float64(x)})
			// 	case 6:
			// 		data6 = append(data6, plotter.XY{X: float64(k), Y: float64(x)})
			// 	case 7:
			// 		data7 = append(data7, plotter.XY{X: float64(k), Y: float64(x)})
			// 	}
			// }
		}
	}

	// ch0
	g0, err := entity.NewGraph(0, "sample", "x", "y", data0)
	if err != nil {
		return err
	}
	if err := u.graphRepo.SaveAsPng("output/"+strconv.Itoa(chunkID)+"-"+"0"+".png", g0); err != nil {
		return err
	}

	lines0 := []string{}
	for _, x := range data0 {
		lines0 = append(lines0, strconv.Itoa(int(x.Y))+"\n")
	}
	f0 := entity.NewFile(lines0)
	if err := u.fileRepo.Create("output/"+strconv.Itoa(chunkID)+"-"+"0"+".dat", f0); err != nil {
		return err
	}

	// ch1
	g1, err := entity.NewGraph(1, "sample", "x", "y", data1)
	if err != nil {
		return err
	}
	if err := u.graphRepo.SaveAsPng("output/"+strconv.Itoa(chunkID)+"-"+"1"+".png", g1); err != nil {
		return err
	}

	lines1 := []string{}
	for _, x := range data1 {
		lines1 = append(lines1, strconv.Itoa(int(x.Y))+"\n")
	}
	f1 := entity.NewFile(lines1)
	if err := u.fileRepo.Create("output/"+strconv.Itoa(chunkID)+"-"+"1"+".dat", f1); err != nil {
		return err
	}

	// ch2
	g2, err := entity.NewGraph(2, "sample", "x", "y", data2)
	if err != nil {
		return err
	}
	if err := u.graphRepo.SaveAsPng("output/"+strconv.Itoa(chunkID)+"-"+"2"+".png", g2); err != nil {
		return err
	}

	lines2 := []string{}
	for _, x := range data2 {
		lines2 = append(lines2, strconv.Itoa(int(x.Y))+"\n")
	}
	f2 := entity.NewFile(lines2)
	if err := u.fileRepo.Create("output/"+strconv.Itoa(chunkID)+"-"+"2"+".dat", f2); err != nil {
		return err
	}

	// ch3
	g3, err := entity.NewGraph(3, "sample", "x", "y", data3)
	if err != nil {
		return err
	}
	if err := u.graphRepo.SaveAsPng("output/"+strconv.Itoa(chunkID)+"-"+"3"+".png", g3); err != nil {
		return err
	}

	lines3 := []string{}
	for _, x := range data3 {
		lines3 = append(lines3, strconv.Itoa(int(x.Y))+"\n")
	}
	f3 := entity.NewFile(lines3)
	if err := u.fileRepo.Create("output/"+strconv.Itoa(chunkID)+"-"+"3"+".dat", f3); err != nil {
		return err
	}

	// ch4
	g4, err := entity.NewGraph(4, "sample", "x", "y", data4)
	if err != nil {
		return err
	}
	if err := u.graphRepo.SaveAsPng("output/"+strconv.Itoa(chunkID)+"-"+"4"+".png", g4); err != nil {
		return err
	}

	lines4 := []string{}
	for _, x := range data4 {
		lines4 = append(lines4, strconv.Itoa(int(x.Y))+"\n")
	}
	f4 := entity.NewFile(lines4)
	if err := u.fileRepo.Create("output/"+strconv.Itoa(chunkID)+"-"+"4"+".dat", f4); err != nil {
		return err
	}

	// ch5
	g5, err := entity.NewGraph(5, "sample", "x", "y", data5)
	if err != nil {
		return err
	}
	if err := u.graphRepo.SaveAsPng("output/"+strconv.Itoa(chunkID)+"-"+"5"+".png", g5); err != nil {
		return err
	}

	lines5 := []string{}
	for _, x := range data5 {
		lines5 = append(lines5, strconv.Itoa(int(x.Y))+"\n")
	}
	f5 := entity.NewFile(lines5)
	if err := u.fileRepo.Create("output/"+strconv.Itoa(chunkID)+"-"+"5"+".dat", f5); err != nil {
		return err
	}

	// ch6
	g6, err := entity.NewGraph(6, "sample", "x", "y", data6)
	if err != nil {
		return err
	}
	if err := u.graphRepo.SaveAsPng("output/"+strconv.Itoa(chunkID)+"-"+"6"+".png", g6); err != nil {
		return err
	}

	lines6 := []string{}
	for _, x := range data6 {
		lines6 = append(lines6, strconv.Itoa(int(x.Y))+"\n")
	}
	f6 := entity.NewFile(lines6)
	if err := u.fileRepo.Create("output/"+strconv.Itoa(chunkID)+"-"+"6"+".dat", f6); err != nil {
		return err
	}

	// ch7
	g7, err := entity.NewGraph(7, "sample", "x", "y", data7)
	if err != nil {
		return err
	}
	if err := u.graphRepo.SaveAsPng("output/"+strconv.Itoa(chunkID)+"-"+"7"+".png", g7); err != nil {
		return err
	}

	lines7 := []string{}
	for _, x := range data7 {
		lines7 = append(lines7, strconv.Itoa(int(x.Y))+"\n")
	}
	f7 := entity.NewFile(lines7)
	if err := u.fileRepo.Create("output/"+strconv.Itoa(chunkID)+"-"+"7"+".dat", f7); err != nil {
		return err
	}

	return nil
}
