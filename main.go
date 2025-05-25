package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// ball type
type Ball struct {
	X, Y   float64
	VX, VY float64
	Radius float64
	Color  color.Color
}

type Game struct {
	Balls []*Ball
	Canvas *ebiten.Image
}

func (g *Game) Draw(screen *ebiten.Image) {
	// redering logic
	for _, ball := range g.Balls {
		op := &ebiten.DrawImageOptions{}
		circle := ebiten.NewImage(int(ball.Radius*2), int(ball.Radius*2))
		circle.Fill(ball.Color)
		op.GeoM.Translate(ball.X - ball.Radius, ball.Y - ball.Radius)
		g.Canvas.DrawImage(circle, op)
	}

	screen.DrawImage(g.Canvas, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600
}

func main() {

	canvas := ebiten.NewImage(800, 600)
	canvas.Fill(color.RGBA{128, 128, 255, 255})

	game := &Game{
		Balls: []*Ball{
			{
				X:      200,
				Y:      150,
				VX:     2,
				VY:     1.5,
				Radius: 3,
				Color:  color.RGBA{255, 255, 255, 255}, // Day: white
			},
			{
				X:      600,
				Y:      450,
				VX:     -2,
				VY:     -1.2,
				Radius: 3,
				Color:  color.RGBA{0, 0, 0, 255}, // Night: black
			},
		},
		Canvas: canvas,
	}

	ebiten.SetWindowTitle("Pong Wars in Go")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Update() error {
	screenWidth, screenHeight := 800, 600

	for _, ball := range g.Balls {
		ball.X += ball.VX
		ball.Y += ball.VY

		if ball.X - ball.Radius < 0 {
			ball.X = ball.Radius
			ball.VX = -ball.VX
		}
		if ball.X + ball.Radius > float64(screenWidth) {
			ball.X = float64(screenWidth) - ball.Radius
			ball.VX = -ball.VX
		}
		if ball.Y - ball.Radius < 0 {
			ball.Y = ball.Radius
			ball.VY = -ball.VY
		}
		if ball.Y + ball.Radius > float64(screenHeight) {
			ball.Y = float64(screenHeight) - ball.Radius
			ball.VY = -ball.VY
		}
	}
	return nil
}

/*
func (g *Game) Update() error {
	screenWidth, screenHeight := 800, 600

	for _, ball := range g.Balls {
		ball.X += ball.VX
		ball.Y += ball.VY

		if ball.X - ball.Radius < 0 || ball.X + ball.Radius > float64(screenWidth) {
			ball.VX = -ball.VX
		}

		if ball.Y - ball.Radius < 0 || ball.Y + ball.Radius > float64(screenHeight) {
			ball.VY = -ball.VY
		}
	}
	return nil
}
*/