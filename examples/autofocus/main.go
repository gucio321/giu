// Package main shows how to set auto-focus on an input field.
// Reference: https://github.com/AllenDang/giu/issues/501
package main

import (
	"github.com/AllenDang/giu"
)

var (
	command     string
	shouldFocus bool
)

func loop() {
	giu.SingleWindow().Layout(
		giu.Custom(func() {
			if shouldFocus {
				shouldFocus = false

				giu.SetKeyboardFocusHere()
			}
		}),
		giu.Row(
			giu.InputText(&command).Size(200),
			giu.Label("<- press any key to start typing here!"),
		),
	)
}

func onAnyKeyPressed(_ giu.Key, _ giu.Modifier, action giu.Action) {
	if action == giu.Press {
		shouldFocus = true
	}
}

func main() {
	wnd := giu.NewMasterWindow("Handle any key event", 640, 480, 0)
	wnd.SetAdditionalInputHandlerCallback(onAnyKeyPressed)
	wnd.Run(loop)
}
