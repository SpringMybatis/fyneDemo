package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"image/color"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Fyne的canvas")
	Rectangles(w)
	w.Resize(fyne.NewSize(600, 600)) // 窗口大小
	w.ShowAndRun()                     // 渲染完毕
}

// 长方形

func Rectangles(w fyne.Window) {
	rect := canvas.NewRectangle(color.White) // 画一个长方形
	rect.StrokeColor = color.Black           // 边框颜色
	rect.StrokeWidth = 5                     // 边框大小

	w.SetContent(rect) //窗口添加控件
}