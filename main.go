package main

import (
	"image"
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/qdorme/maze-go/maze"
)

func main() {

	a := app.New()
	w := a.NewWindow("Maze Size")

	c := container.NewVBox()

	widthInput := widget.NewEntry()
	widthInput.SetText("10")
	heightInput := widget.NewEntry()
	heightInput.SetText("10")

	c.Add(container.NewHBox(
		canvas.NewText("Width ", color.Black),
		widthInput,
	))
	c.Add(container.NewHBox(
		canvas.NewText("Height", color.Black),
		heightInput,
	))
	c.Add(widget.NewButton("Generate", func() {

		width, err := strconv.Atoi(widthInput.Text)
		if err != nil {
			return
		}
		height, err := strconv.Atoi(heightInput.Text)
		if err != nil {
			return
		}

		m := a.NewWindow("Maze")
		newMaze := maze.NewMaze(width, height)
		m.Resize(fyne.NewSize(float32(width*15), float32(height*15+10)))

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

		m.SetContent(canvasImage)

		m.Show()
	}))

	w.SetContent(c)
	w.SetMaster()

	w.ShowAndRun()
}
