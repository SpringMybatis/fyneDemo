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
	cles(w)
	w.Resize(fyne.NewSize(600, 600)) // 窗口大小
	w.ShowAndRun()                     // 渲染完毕
}

// 圆形

func cles(w fyne.Window) {
	c := canvas.NewCircle(color.RGBA{R: 255, G: 11, B: 11, A: 235})
	c.StrokeWidth = 10
	c.StrokeColor = color.RGBA{R: 10, G: 10, B: 10, A: 235}
	c.Position1 = fyne.NewPos(10, 10)
	c.Position2 = fyne.NewPos(100, 100)

	c1 := canvas.NewCircle(color.RGBA{R: 255, G: 11, B: 11, A: 235})
	c1.StrokeWidth = 10
	c1.StrokeColor = color.RGBA{R: 10, G: 10, B: 10, A: 235}
	c1.Position1 = fyne.NewPos(100, 100)
	c1.Position2 = fyne.NewPos(200, 200)
	w.SetContent(container.NewWithoutLayout(c, c1)) // 有定位必须在  newWithoutLayout 上面画
}
