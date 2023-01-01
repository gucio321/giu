package main

import (
	"fmt"

	"github.com/AllenDang/cimgui-go"
	g "github.com/AllenDang/giu"
)

var (
	dropTarget string = "Drop here"
)

func loop() {
	g.SingleWindow().Layout(
		g.Row(
			g.Custom(func() {
				g.Button("Drag me: 9").Build()
				if cimgui.BeginDragDropSource() {
					cimgui.SetDragDropPayload("DND_DEMO", 9)
					g.Label("9").Build()
					cimgui.EndDragDropSource()
				}
			}),
			g.Custom(func() {
				g.Button("Drag me: 10").Build()
				if cimgui.BeginDragDropSource() {
					cimgui.SetDragDropPayload("DND_DEMO", 10)
					g.Label("10").Build()
					cimgui.EndDragDropSource()
				}
			}),
		),
		g.InputTextMultiline(&dropTarget).Size(g.Auto, g.Auto).Flags(g.InputTextFlagsReadOnly),
		g.Custom(func() {
			if cimgui.BeginDragDropTarget() {
				payload := cimgui.AcceptDragDropPayload("DND_DEMO")
				if payload != 0 {
					dropTarget = fmt.Sprintf("Dropped value: %d", payload.Data())
				}
				cimgui.EndDragDropTarget()
			}
		}),
	)
}

func main() {
	wnd := g.NewMasterWindow("Drag and Drop", 600, 400, g.MasterWindowFlagsNotResizable)
	wnd.Run(loop)
}
