package main

import (
	"image"
	_ "image"
	"image/color"
	_ "image/png"
	"math/rand"
	_ "math/rand"
	_ "os"
	"time"

	"golang.org/x/text/cases"
)

var (
	width  = 600
	height = 400
)

type colorpicker struct {
	color1 color.Color
	color2 color.Color
	color3 color.Color
}

var (
	red = colorpicker{
		//#e2471d
		//#e6140e
		//#e66352

	}
	yellow = colorpicker{
		//#f69e2c
		//#dec443
		//#d4aa6d
	}
	brown = colorpicker{
		//#9a7b55
		//#da8b40
		//#db8468
	}
)

func main() {
	rand.Seed(time.Now().Unix())
	img := image.NewNRGBA(image.Rect(0,0,width,height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++{
			srand := rand.Intn(3)
			switch rand.Intn(3){
			case 0:
				switch srand{
					case 0:
					case 1:
					case 2:
				}		
			case 1:
				switch srand{
					case 0:
					case 1:
					case 2:
				}
			case 2:
				switch srand{
					case 0:
					case 1:
					case 2:
				}
		 }		
		}
	}
}
