package ui


import (
	"context"
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"netWatcher/internal"
	"time"
)

type AppMainWindow struct {
	ctx           context.Context
	playBtn       *widget.Button // 开始监控
	stopBtn       *widget.Button // 停止监控
	ethCards      *widget.Select // 网卡选择
	ipLabel       *widget.Label  // IP地址栏
	macLabel      *widget.Label  // MAC地址栏
	upLoadSpeed   *widget.Label  // 上行速度
	downLoadSpeed *widget.Label  // 下行速度
	anl           *internal.Analyzer
	// 要绘制图表
}

func (mw *AppMainWindow) Run() {
	a := app.New()

	ncs, _ := internal.GetNetCardsWithIPv4Addr()
	ethCards := make([]string, 0, 4)
	for _, nc := range ncs {
		ethCards = append(ethCards, nc.GetName())
	}

	mw.anl = &internal.Analyzer{}
	mw.anl.Init()

	mw.ethCards = widget.NewSelect(ethCards, func(s string) {
		var ncard internal.NetCard
		if mw.anl.Running {
			// 如果没能停止，Stop()会阻塞
			mw.anl.Stop()
		}
		validCard := false
		for _, nc := range ncs {
			if nc.GetName() == s {
				mw.ipLabel.SetText(nc.GetIPv4Addr())
				ncard = nc
				validCard = true
			}
		}

		// 如果是有效网卡，则进行捕捉
		if validCard {
			mw.anl.Nc = &ncard
			go mw.anl.Capture()
		}

	})

	//mw.ethCards.SetSelected(ethCards[0])
	mw.upLoadSpeed = widget.NewLabel("---")
	mw.downLoadSpeed = widget.NewLabel("----")

	go mw.UpdateSpeed()

	mw.ipLabel = widget.NewLabel("---")



	w := a.NewWindow("FyneNet")
	w.SetTitle("FyneNet")
	w.SetContent(fyne.NewContainerWithLayout(layout.NewGridLayout(1),
		fyne.NewContainerWithLayout(layout.NewGridLayout(2), mw.ethCards, mw.ipLabel),
		fyne.NewContainerWithLayout(layout.NewGridLayout(2), mw.upLoadSpeed, mw.downLoadSpeed),
	))
	w.ShowAndRun()
}


func (mw *AppMainWindow) UpdateSpeed() {
	for  {
		upspeed := fmt.Sprintf("Down:%.2fkb/s", mw.anl.GetDownSpeed())
		downspeed := fmt.Sprintf("Up:%.2fkb/s", mw.anl.GetUpSpeed())
		mw.upLoadSpeed.SetText(upspeed)
		mw.downLoadSpeed.SetText(downspeed)
		time.Sleep(1 * time.Second)
	}
}