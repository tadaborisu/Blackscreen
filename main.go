package main

import (
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/lxn/win"
)

var (
	windowWidth  = int(win.GetSystemMetrics(win.SM_CXSCREEN))
	windowHeight = int(win.GetSystemMetrics(win.SM_CYSCREEN))
)

func window() {
	file, _ := os.Open("icon.png")
	icon, _, _ := image.Decode(file)
	iconImage := []image.Image{icon}
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Blackscreen")
	ebiten.SetWindowDecorated(false)
	ebiten.SetMaxTPS(0)
	ebiten.SetWindowIcon(iconImage)
}

//Game :
type Game struct{}

//Update :
func (g *Game) Update() error {
	return nil
}

//Draw :
func (g *Game) Draw(screen *ebiten.Image) {
}

//Layout :
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1, 1
}

func main() {
	window()
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
