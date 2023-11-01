package main

import (
	"github.com/CalvinWalmer/img_mod/Colors"
	"github.com/CalvinWalmer/img_mod/GetPic"
	"github.com/CalvinWalmer/img_mod/Grayscale"
	"github.com/CalvinWalmer/img_mod/Text"
)

func main() {

	GetPic.GetPic()
	Text.Text()
	Grayscale.Grayscale()
	Colors.Colors()
}
