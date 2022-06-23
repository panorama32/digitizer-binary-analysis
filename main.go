package main

import (
	"fmt"

	"github.com/panorama32/digitizer-binary-analysis/domain/data"
	"github.com/panorama32/digitizer-binary-analysis/store"
	"github.com/panorama32/digitizer-binary-analysis/usecase"
)

func main() {
	fileRepo := store.NewFileRepository()
	graphRepo := store.NewGraphRepository()
	processUC := usecase.NewProcessUseCase(fileRepo, graphRepo)

	path := "data/20220616_16"

	d, err := data.New(path)
	if err != nil {
		fmt.Println(err)
	}

	// loopN := len(d.Bytes) / 32776

	if err := processUC.SetUp(); err != nil {
		fmt.Println(err)
	}

	for i := 0; i < 5; i++ {
		err = processUC.ProcessChunk(i, d)
		if err != nil {
			fmt.Println(err)
		}
	}

	// files.CreateFile("sample.dat", datax)
}
