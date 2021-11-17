package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Fyne的canvas")
	lines(w)
	w.Resize(fyne.NewSize(600, 600)) // 窗口大小
	w.ShowAndRun()                     // 渲染完毕
}

// 线条
func lines(w fyne.Window) {
	l := canvas.NewLine(color.RGBA{R: 255, A: 155})
	l.StrokeWidth = 25
	l.Position1 = fyne.NewPos(0, 0)
	l.Position2 = fyne.NewPos(0, 100)
	l.Refresh()

	l1 := canvas.NewLine(color.RGBA{R: 255, G: 100, B: 100, A: 155})
	l1.StrokeWidth = 15
	l1.Position1 = fyne.NewPos(0, 0)
	l1.Position2 = fyne.NewPos(100, 100)
	l.Refresh()

	w.SetContent(container.NewWithoutLayout(l, l1))
}