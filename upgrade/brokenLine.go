package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
	"math/rand"
	"strconv"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Fyne的canvas")
	zxLines(w)
	w.Resize(fyne.NewSize(600, 600)) // 窗口大小
	w.ShowAndRun()                     // 渲染完毕
}

// 折线图

func zxLines(w fyne.Window) {

	data := genData2(10, 100)
	ly := container.NewWithoutLayout()

	for i, d := range data {

		// 线条
		line := canvas.NewLine(genColor2())
		line.Position1 = fyne.NewPos(float32(100*i+10), 200) // 35-200 70-200
		line.Position2 = fyne.NewPos(float32(100*i+10), d)   // 35-随机 70-随机
		line.StrokeWidth = 30

		// 文字
		text := canvas.NewText(strconv.FormatFloat(float64(200-d), 'f', 3, 64), color.Black)
		text.Move(line.Position2)
		text.TextSize = 15
		text.Alignment = fyne.TextAlignCenter

		ly.Objects = append(ly.Objects, line)
		ly.Objects = append(ly.Objects, text)
	}

	for i := 0; i < len(data)-1; i++ {
		// 一跳折线
		linez := canvas.NewLine(genColor())
		linez.Position1 = fyne.NewPos(float32(100*i+10), data[i])       // 35-200 70-200  起始坐标
		linez.Position2 = fyne.NewPos(float32(100*(i+1)+10), data[i+1]) // 35-随机 70-随机  目标地点坐标
		linez.StrokeWidth = 5
		ly.Objects = append(ly.Objects, linez)

	}

	w.SetContent(ly)
}

// 生存随机坐标

func genData2(count, max int) (data []float32) {
	for i := 0; i < count; i++ {
		data = append(data, float32(rand.Intn(max)))
	}
	return
}

// 生存随机颜色

func genColor2() color.Color {
	return color.RGBA{R: uint8(rand.Intn(256)), G: uint8(rand.Intn(256)), B: uint8(rand.Intn(256)), A: 255}
}