package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image"
	"image/color"
	"image/jpeg"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Fyne的canvas")
	scatter(w)
	w.Resize(fyne.NewSize(600, 600)) // 窗口大小
	w.ShowAndRun()                     // 渲染完毕
}

// 随机 小圆点

func scatter(w fyne.Window) {
	c := container.NewWithoutLayout()

	cdata := genData(20, 1000)
	var r float32 = 15
	ydata := genData(20, 600)

	for i, v := range cdata {
		c1 := canvas.NewCircle(genColor())
		c1.Position1 = fyne.NewPos(v-r, ydata[i]-r)
		c1.Position2 = fyne.NewPos(v+r, ydata[i]+r)
		fmt.Println(i, v)
		c.Objects = append(c.Objects, c1)
	}

	btn := widget.NewButton("image button", func() {
		for i := 0; i < 10; i++ {
			refreshPos(20, 1000, 600, c)
			saveImage(w.Canvas().Capture(), "test"+strconv.Itoa(i)+".jpeg")
		}
	})
	btnRefresh := widget.NewButtonWithIcon("", theme.ViewRefreshIcon(), func() {
		refreshPos(20, 1000, 600, c)
	})

	w.SetContent(container.NewVBox(container.NewHBox(btn, btnRefresh), c))
}

func saveImage(img image.Image, path string) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("os.create err")
	}
	defer file.Close()
	jpeg.Encode(file, img, nil)
	fmt.Println("done")
}

func refreshPos(counts, maxX, maxY int, c *fyne.Container) {
	cdata := genData(counts, maxX)
	time.Sleep(time.Millisecond * 100)
	ydata := genData(counts, maxY)
	for i, v := range c.Objects {
		v.Move(fyne.NewPos(cdata[i], ydata[i]))
	}
}

// 生存随机坐标

func genData(count, max int) (data []float32) {
	for i := 0; i < count; i++ {
		data = append(data, float32(rand.Intn(max)))
	}
	return
}

// 生存随机颜色

func genColor() color.Color {
	return color.RGBA{R: uint8(rand.Intn(256)), G: uint8(rand.Intn(256)), B: uint8(rand.Intn(256)), A: 255}
}