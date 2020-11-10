package physarum

import (
	"fmt"
	"time"
)

func one(model *Model, iterations int) {
	now := time.Now().UTC().UnixNano()
	path := fmt.Sprintf("out%d.png", now)
	fmt.Println()
	fmt.Println(path)
	for _, config := range model.Configs {
		fmt.Println(*config)
	}
	for i := 0; i < iterations; i++ {
		// fmt.Println(i)
		model.Step()
	}
	SavePNG(path, Image(model.W, model.H, model.Colors(), 0, 0, 1/2.2))
}

func frames(model *Model, rate int) {
	saveImage := func(path string, w, h int, colors [][]float64, ch chan bool) {
		im := Image(w, h, colors, 0, 0, 1)
		SavePNG(path, im)
		if ch != nil {
			ch <- true
		}
	}

	ch := make(chan bool, 1)
	ch <- true
	for i := 0; ; i++ {
		fmt.Println(i)
		model.Step()
		if i%rate == 0 {
			<-ch
			path := fmt.Sprintf("out%08d.png", i/rate)
			go saveImage(path, model.W, model.H, model.Colors(), ch)
		}
	}
}

func Run() {
	for {
		configs := RandomConfigs(3)
		model := NewModel(1024, 1024, configs)
		one(model, 500)
	}
	// frames(model, 1)
}