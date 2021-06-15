package main

import (
	"fmt"

	"github.com/gucio321/giu"
)

var checkbox = true

func loop() {
	flags := giu.FocusedFlagsRootAndChildWindows
	fmt.Println(giu.IsWindowFocused(flags))
	giu.Window("window 1").
		Shortcuts(
			giu.Shortcut{func() { fmt.Println("c") }, giu.KeyZ, giu.ModControl},
		).Layout(
		giu.Checkbox("Press ctrl+b to change my state", &checkbox),
	)
	fmt.Println(giu.IsWindowFocused(flags))

	giu.Window("window 2").Layout(
		giu.Checkbox("Press ctrl+b to change my state", &checkbox),
	)
	fmt.Println(giu.IsWindowFocused(flags))
}

func main() {
	wnd := giu.NewMasterWindow("keybard shortcuts", 640, 480, 0)

	giu.RegisterShortcut(func() {
		checkbox = !checkbox
	}, giu.KeyB, giu.ModControl, true)

	wnd.Run(loop)
}
