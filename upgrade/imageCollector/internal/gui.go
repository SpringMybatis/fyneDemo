package internal

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

type AppGUI struct {
	items     *widget.Select
	srcDir    *widget.Entry
	outputDir *widget.Entry
	btn       *widget.Button
}

// Gather 收集对应文件（采用直接移动的方式）
func (ap *AppGUI) Gather() {
	ap.srcDir.Disable()
	filters := make([]string, 0, 10)

	if ap.items.Selected != "" {
		filters = append(filters, "."+ap.items.Selected)
	}

	imgcl := &ImgCollector{}
	imgcl.SetDir(ap.srcDir.Text)
	imgcl.SetFilters(filters)
	imgcl.Run()
}

func (ap *AppGUI) Run() {
	ap.srcDir = widget.NewEntry()
	ap.outputDir = widget.NewEntry()
	ap.items = widget.NewSelect([]string{"bmp", "png", "jpg", "jpeg"}, nil)
	ap.btn = widget.NewButton("Run", ap.Gather)

	a := app.New()

	w := a.NewWindow("图片")
	w.SetTitle("图片收集器")
	w.SetContent(widget.NewVBox(
		ap.srcDir,
		ap.outputDir,
		ap.items,
		ap.btn,
	))
	w.ShowAndRun()
}
