package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
	"math/rand"
	"time"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Fyne的canvas")
	dtBar(w)
	w.Resize(fyne.NewSize(600, 600)) // 窗口大小
	w.ShowAndRun()                     // 渲染完毕
}

// 柱状图
func dtBar(w fyne.Window) {
	xdata := genData1(10, 100)
	c := container.NewWithoutLayout()
	for i, v := range xdata {
		line := canvas.NewLine(genColor1())
		line.Position1 = fyne.NewPos(float32(100*i+10), 200) // 35-200 70-200
		line.Position2 = fyne.NewPos(float32(100*i+10), v)   // 35-随机 70-随机
		line.StrokeWidth = 30
		Bar(line)
		c.Objects = append(c.Objects, line)
	}

	w.SetContent(c)
}

// Bar 讲柱状图动态加载
func Bar(line *canvas.Line) {
	go func() {
		for i := 90; i > 10; i-- {
			line.Position2 = fyne.NewPos(line.Position2.X, float32(i))
			line.Refresh()
			time.Sleep(time.Millisecond * 100)
		}
	}()
}
// 生存随机坐标

func genData1(count, max int) (data []float32) {
	for i := 0; i < count; i++ {
		data = append(data, float32(rand.Intn(max)))
	}
	return
}

// 生存随机颜色

func genColor1() color.Color {
	return color.RGBA{R: uint8(rand.Intn(256)), G: uint8(rand.Intn(256)), B: uint8(rand.Intn(256)), A: 255}
}