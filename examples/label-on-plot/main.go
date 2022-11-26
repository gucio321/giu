package main

import (
	"image"

	"github.com/AllenDang/giu"
	"golang.org/x/image/colornames"
)

func loop() {
	giu.SingleWindow().Layout(
		giu.Plot("my plot").Plots(
			giu.PieChart([]string{"chart 1", "2"}, []float64{0.45, 0.55}, 0.5, 0.5, 0.45),
			giu.Custom(func() {
				c := cimgui
				c.AddText(image.Pt(20, 80), colornames.Red, "hello world")
			}),
		),
	)
}

func main() {
	wnd := giu.NewMasterWindow("label on plot canvas", 640, 480, 0)
	wnd.Run(loop)
}
