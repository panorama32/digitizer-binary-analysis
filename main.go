package main

func main() {
	// d := digitizerdata.New("data/20220616_16")

	// loopN := len(d.Buf)/32776

	// data0 := plotter.XYs{}
	// data1 := plotter.XYs{}
	// data2 := plotter.XYs{}
	// data3 := plotter.XYs{}
	// data4 := plotter.XYs{}
	// data5 := plotter.XYs{}
	// data6 := plotter.XYs{}
	// data7 := plotter.XYs{}

	// datax := []string{}

	// for i := 0; i < 1; i++ {
	// 	d.Read(2)
	// 	d.Read(2)
	// 	if i%2 == 1 {
	// 		for j := 0; j < 8; j++ {
	// 			for k := 0; k < 1024; k++ {
	// 				d.ReadInt16()
	// 			}
	// 		}
	// 	} else {
	// 		for j := 0; j < 8; j++ {
	// 			data := plotter.XYs{}
	// 			for k := 0; k < 1024; k++ {
	// 				x, _ := d.ReadInt16()
	// 				if k == 0 {
	// 					// data = append(data, plotter.XY{X: float64(k), Y: float64(x)})
	// 					datax = append(datax, strconv.Itoa(x))
	// 				} else {
	// 					datax = append(datax, strconv.Itoa(x)+"\n")
	// 					data = append(data, plotter.XY{X: float64(k), Y: float64(x)})
	// 				}
	// 			}
	// 			switch j {
	// 			case 0:
	// 				data0 = append(data0, data...)
	// 			case 1:
	// 				data1 = append(data1, data...)
	// 			case 2:
	// 				data2 = append(data2, data...)
	// 			case 3:
	// 				data3 = append(data3, data...)
	// 			case 4:
	// 				data4 = append(data4, data...)
	// 			case 5:
	// 				data5 = append(data5, data...)
	// 			case 6:
	// 				data6 = append(data6, data...)
	// 			case 7:
	// 				data7 = append(data7, data...)
	// 			}
	// 		}
	// 	}
	// }

	// files.CreateFile("sample.dat", datax)

	// if err := drawer.Draw(0, "test", "x-axis", "y-axis", data0); err != nil {
	// 	fmt.Errorf("%w", err)
	// }
	// if err := drawer.Draw(1, "test", "x-axis", "y-axis", data1); err != nil {
	// 	fmt.Errorf("%w", err)
	// }
	// if err := drawer.Draw(2, "test", "x-axis", "y-axis", data2); err != nil {
	// 	fmt.Errorf("%w", err)
	// }
	// if err := drawer.Draw(3, "test", "x-axis", "y-axis", data3); err != nil {
	// 	fmt.Errorf("%w", err)
	// }
	// if err := drawer.Draw(4, "test", "x-axis", "y-axis", data4); err != nil {
	// 	fmt.Errorf("%w", err)
	// }
	// if err := drawer.Draw(5, "test", "x-axis", "y-axis", data5); err != nil {
	// 	fmt.Errorf("%w", err)
	// }
	// if err := drawer.Draw(6, "test", "x-axis", "y-axis", data6); err != nil {
	// 	fmt.Errorf("%w", err)
	// }
	// if err := drawer.Draw(7, "test", "x-axis", "y-axis", data7); err != nil {
	// 	fmt.Errorf("%w", err)
	// }

	return
}
