package main

import (
	_ "image"
	"image/color"
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
	red    = colorpicker{}
	yellow = colorpicker{}
	brown  = colorpicker{}
)

func main() {

}
