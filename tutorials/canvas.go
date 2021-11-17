package tutorials

import (
	"fyne.io/fyne/v2/container"
	"image/color"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
)

func rgbGradient(x, y, w, h int) color.Color {
	g := int(float32(x) / float32(w) * float32(255))
	b := int(float32(y) / float32(h) * float32(255))

	return color.NRGBA{R: uint8(255 - b), G: uint8(g), B: uint8(b), A: 0xff}
}

// canvasScreen loads a graphics example panel for the demo app
func canvasScreen(_ fyne.Window) fyne.CanvasObject {
	gradient := canvas.NewHorizontalGradient(color.NRGBA{R: 0x80, A: 0xff}, color.NRGBA{G: 0x80, A: 0xff})
	go func() {
		for {
			time.Sleep(time.Second)

			gradient.Angle += 45
			if gradient.Angle >= 360 {
				gradient.Angle -= 360
			}
			canvas.Refresh(gradient)
		}
	}()

	//lines := dtBar()
	return container.New(layout.NewGridWrapLayout(fyne.NewSize(90, 90)),
		canvas.NewImageFromResource(theme.FyneLogo()),
		&canvas.Rectangle{
			FillColor:   color.NRGBA{R: 0x80, A: 0xff},
			StrokeColor: color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff},
			StrokeWidth: 1,
		},
		&canvas.Line{
			StrokeColor: color.NRGBA{B: 0x80, A: 0xff},
			StrokeWidth: 5,
		},
		&canvas.Circle{
			StrokeColor: color.NRGBA{B: 0x80, A: 0xff},
			FillColor:   color.NRGBA{R: 0x30, G: 0x30, B: 0x30, A: 0x60},
			StrokeWidth: 2,
		},
		canvas.NewText("Text", color.NRGBA{G: 0x80, A: 0xff}),
		canvas.NewRasterWithPixels(rgbGradient),
		gradient,
		canvas.NewRadialGradient(color.NRGBA{R: 0x80, A: 0xff}, color.NRGBA{G: 0x80, B: 0x80, A: 0xff}),
		canvas.NewImageFromFile("./data/images/girl.jpg"),
		//&canvas.Image{},
	)
}

// 柱状图
func dtBar() []canvas.Line {
	var lines []canvas.Line
	xData := genData(5, 100)
	for i, v := range xData {
		line := &canvas.Line{}
		line.StrokeColor = genColor()
		line.Position1 = fyne.NewPos(float32(100*i+10), 200) // 35-200 70-200
		line.Position2 = fyne.NewPos(float32(100*i+10), v)   // 35-随机 70-随机
		line.StrokeWidth = 20
		Bar(line)
		lines = append(lines, *line)
	}
	return lines
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
