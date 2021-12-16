package main

import (
	_ "image"
	"image/color"
	_ "math/rand"
	_ "image/png"
	_ "os"
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

}
