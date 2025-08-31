package main

import (
	"image"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"github.com/qdorme/maze-go/maze"
)

func main() {

	width, height := 30, 20

	newMaze := maze.NewMaze(width, height)

	a := app.New()
	w := a.NewWindow("Maze")
	w.Resize(fyne.NewSize(float32(width*15), float32(height*15+10)))

	canvasImage := canvas.NewImageFromImage(image.NewRGBA(image.Rect(0, 0, 10, 10)))

	go func() {
		imageMaze := newMaze.Start()

		for {
			select {
			case img := <-imageMaze:
				fyne.Do(func() {
					canvasImage.Image = img
					canvasImage.Refresh()
				})
			}
		}
	}()

	w.SetContent(canvasImage)

	w.ShowAndRun()
}
