package main

import (
	"image"
	_ "image"
	"image/color"
	"image/png"
	_ "image/png"
	"math/rand"
	_ "math/rand"
	"os"
	_ "os"
	"time"
)

var (
	width  = 60
	height = 40
)

type colorpicker struct {
	color1 color.Color
	color2 color.Color
	color3 color.Color
}

var (
	red = colorpicker{
		//#e2471d
		color1: color.RGBA{226, 71, 29, 255},
		//#e6140e
		color2: color.RGBA{230, 20, 14, 255},
		//#e66352
		color3: color.RGBA{230, 99, 82, 255},
	}

	yellow = colorpicker{
		//#f69e2c
		color1: color.RGBA{246, 158, 44, 255},
		//#dec443
		color2: color.RGBA{222, 196, 67, 255},
		//#d4aa6d
		color3: color.RGBA{212, 170, 109, 255},
	}
	brown = colorpicker{
		//#9a7b55
		color1: color.RGBA{154, 123, 85, 255},
		//#da8b40
		color2: color.RGBA{218, 139, 64, 255},
		//#db8468
		color3: color.RGBA{208, 132, 104, 255},
	}
)

func main() {
	rand.Seed(time.Now().Unix())
	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			srand := rand.Intn(3)
			switch rand.Intn(3) {
			case 0:
				switch srand {
				case 0:
					img.Set(x, y, red.color1)
				case 1:
					img.Set(x, y, red.color2)
				case 2:
					img.Set(x, y, red.color3)
				}
			case 1:
				switch srand {
				case 0:
					img.Set(x, y, yellow.color1)
				case 1:
					img.Set(x, y, red.color2)
				case 2:
					img.Set(x, y, red.color3)
				}
			case 2:
				switch srand {
				case 0:
					img.Set(x, y, brown.color1)
				case 1:
					img.Set(x, y, brown.color2)
				case 2:
					img.Set(x, y, brown.color3)
				}
			}
		}

		f, _ := os.Create("./image2.png")
		png.Encode(f, img)
		f.Close()
	}
}
