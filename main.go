package main

import (
	"fmt"
	"image/color"
	"math"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func imagToIndex(imag float64) float64 {
	gamma := 0.57721566490153286060651209008240243104215933593992
	e := 2.7182818284590452353602874713526624977572
	gamma_to_the_e := math.Pow(gamma, e)
	two_root_3_pi := 2 * math.Sqrt(3*math.Pi)
	return_this := math.Sqrt(6*gamma_to_the_e/imag+6*imag+math.Pi)/two_root_3_pi - 1.0/2.0
	return return_this
}

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Line")

	text := canvas.NewText("Text Object", color.White)
	text.Alignment = fyne.TextAlignTrailing
	text.TextStyle = fyne.TextStyle{Italic: true}

	c := container.NewWithoutLayout(&canvas.Line{
		Position1:   fyne.NewPos(500, 500),
		Position2:   fyne.NewPos(501, 500),
		StrokeColor: color.White,
		StrokeWidth: 1,
	})

	v := container.NewVBox(text, c)
	w.SetContent(v)

	imag := 10000.0

	go func() {
		for {
			numLines := int(imag / math.Pi)
			for i := 1; i < numLines; i++ {

				prevLine := c.Objects[i-1].(*canvas.Line)
				x := 100 * math.Cos(imag*math.Log(float64(i))) / math.Pow(float64(i), .5)
				y := 100 * math.Sin(imag*math.Log(float64(i))) / math.Pow(float64(i), .5)

				// if we need to add a new line
				if i == len(c.Objects) {
					c.Add(&canvas.Line{
						StrokeColor: color.White,
						StrokeWidth: 1,
					})
				}

				// update existing line
				line := c.Objects[i].(*canvas.Line)
				line.Position1 = prevLine.Position2
				line.Position2 = prevLine.Position2.AddXY(float32(x), float32(y))
			}
			text.Text = fmt.Sprintf("%.6f", imagToIndex(imag))
			text.Refresh()
			imag += 0.001
			time.Sleep(1000)
			c.Refresh()
		}
	}()

	w.Resize(fyne.NewSize(1000, 1000))
	w.ShowAndRun()
}
