package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

//Game struct is an instance of ebiten Game Inteface with 2 function Update and Layout.
//View more github.com\hajimehoshi\ebiten@v1.11.7\run.go
type Game struct {
	count int
}

//Update every frame
func (g *Game) Update(screen *ebiten.Image) error {
	g.count++
	return nil
}

//Draw (the recommended way), Update updates only the game logic and Draw draws the screen.
// In this case, the argument screen's updated content by Update is not adopted for the actual game screen,
// and the screen's updated content by Draw is adopted instead.
//
// Without Draw (the legacy way), Update updates the game logic and also draws the screen.
// In this case, the argument screen's updated content by Update is adopted for the actual game screen.
func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello!")
}

// Layout is called almost every frame.
// If Layout returns non-positive numbers, the caller can panic.
// You can return a fixed screen size if you don't care, or you can also return a calculated screen size
// adjusted with the given outside size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600
}

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("My Game")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
