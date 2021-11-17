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
	txt(w)
	w.Resize(fyne.NewSize(600, 600)) // 窗口大小
	w.ShowAndRun()                     // 渲染完毕
}

// 渲染文本框
func txt(w fyne.Window) {
	// 文本框
	txt := canvas.NewText("12123", color.RGBA{R: 255, A: 255})
	txt.Alignment = fyne.TextAlignCenter // 文字对齐方式
	txt.TextSize = 10
	txt.TextStyle = fyne.TextStyle{ // 文字状态
		Bold:      true,
		Italic:    true,
		Monospace: false,
	}

	txt1 := canvas.NewText("456456", color.RGBA{R: 255, G: 205, B: 95, A: 255})
	txt1.Alignment = fyne.TextAlignCenter // 文字对齐方式
	txt1.TextSize = 20
	txt1.TextStyle = fyne.TextStyle{ // 文字状态
		Bold:      false, // 加粗
		Italic:    false, // 斜体
		Monospace: true,
	}
	w.SetContent(container.NewVBox(txt, txt1)) //窗口添加控件
}