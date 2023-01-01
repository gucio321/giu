package giu

import (
	"fmt"

	"github.com/AllenDang/cimgui-go"
)

// SingleWindow creates one window filling all available space
// in MasterWindow. If SingleWindow is set up, no other windows may be
// defined.
func SingleWindow() *WindowWidget {
	size := Context.platform.DisplaySize()
	title := fmt.Sprintf("SingleWindow_%d", Context.GetWidgetIndex())

	return Window(title).
		Flags(
			WindowFlags(cimgui.WindowFlags_NoTitleBar)|
				WindowFlags(cimgui.WindowFlags_NoCollapse)|
				WindowFlags(cimgui.WindowFlags_NoScrollbar)|
				WindowFlags(cimgui.WindowFlags_NoMove)|
				WindowFlags(cimgui.WindowFlags_NoResize)).
		Size(size[0], size[1])
}

// SingleWindowWithMenuBar creates a SingleWindow and allows to add menubar on its top.
func SingleWindowWithMenuBar() *WindowWidget {
	size := Context.platform.DisplaySize()
	title := fmt.Sprintf("SingleWindow_%d", Context.GetWidgetIndex())

	return Window(title).
		Flags(
			WindowFlags(cimgui.WindowFlags_NoTitleBar)|
				WindowFlags(cimgui.WindowFlags_NoCollapse)|
				WindowFlags(cimgui.WindowFlags_NoScrollbar)|
				WindowFlags(cimgui.WindowFlags_NoMove)|
				WindowFlags(cimgui.WindowFlags_MenuBar)|
				WindowFlags(cimgui.WindowFlags_NoResize)).Size(size[0], size[1])
}

var _ Disposable = &windowState{}

type windowState struct {
	hasFocus bool
	currentPosition,
	currentSize cimgui.ImVec2
}

// Dispose implements Disposable interface.
func (s *windowState) Dispose() {
	// noop
}

// WindowWidget represents cimgui.Window
// Windows are used to display ui widgets.
// They are in second place in the giu hierarchy (after the MasterWindow)
// NOTE: to disable multiple window, use SingleWindow.
type WindowWidget struct {
	title         string
	open          *bool
	flags         WindowFlags
	x, y          float32
	width, height float32
	bringToFront  bool
}

// Window creates a WindowWidget.
func Window(title string) *WindowWidget {
	return &WindowWidget{
		title: title,
	}
}

// IsOpen sets if window widget is `opened` (minimized).
func (w *WindowWidget) IsOpen(open *bool) *WindowWidget {
	w.open = open
	return w
}

// Flags sets window flags.
func (w *WindowWidget) Flags(flags WindowFlags) *WindowWidget {
	w.flags = flags
	return w
}

// Size sets window size
// NOTE: size can be changed by user, if you want to prevent
// user from changing window size, use NoResize flag.
func (w *WindowWidget) Size(width, height float32) *WindowWidget {
	w.width, w.height = width, height
	return w
}

// Pos sets the window start position
// NOTE: The position could be changed by user later.
// To prevent user from changing window position use
// WIndowFlagsNoMove.
func (w *WindowWidget) Pos(x, y float32) *WindowWidget {
	w.x, w.y = x, y
	return w
}

// Layout is a final step of the window setup.
// it should be called to add a layout to the window and build it.
func (w *WindowWidget) Layout(widgets ...Widget) {
	if widgets == nil {
		return
	}

	ws := w.getState()

	if w.flags&WindowFlags(cimgui.WindowFlags_NoMove) != 0 && w.flags&WindowFlags(cimgui.WindowFlags_NoResize) != 0 {
		cimgui.SetNextWindowPos(cimgui.ImVec2{X: w.x, Y: w.y})
		cimgui.SetNextWindowSize(cimgui.ImVec2{X: w.width, Y: w.height})
	} else {
		cimgui.SetNextWindowPosV(cimgui.ImVec2{X: w.x, Y: w.y}, cimgui.Cond_FirstUseEver, cimgui.ImVec2{X: 0, Y: 0})
		cimgui.SetNextWindowSizeV(cimgui.ImVec2{X: w.width, Y: w.height}, cimgui.Cond_FirstUseEver)
	}

	if w.bringToFront {
		cimgui.SetNextWindowFocus()

		w.bringToFront = false
	}

	widgets = append(widgets,
		Custom(func() {
			hasFocus := IsWindowFocused(0)
			if !hasFocus && ws.hasFocus {
				Context.InputHandler.UnregisterWindowShortcuts()
			}

			ws.hasFocus = hasFocus

			ws.currentPosition = cimgui.GetWindowPos()
			ws.currentSize = cimgui.GetWindowSize()
		}),
	)

	showed := cimgui.BeginV(Context.FontAtlas.RegisterString(w.title), w.open, cimgui.WindowFlags(w.flags))

	if showed {
		Layout(widgets).Build()
	}

	cimgui.End()
}

// CurrentPosition returns a current position of the window.
func (w *WindowWidget) CurrentPosition() (x, y float32) {
	pos := w.getState().currentPosition
	return pos.X, pos.Y
}

// CurrentSize returns current size of the window.
func (w *WindowWidget) CurrentSize() (width, height float32) {
	size := w.getState().currentSize
	return size.X, size.Y
}

// BringToFront sets window focused.
func (w *WindowWidget) BringToFront() {
	w.bringToFront = true
}

// HasFocus returns true if window is focused.
func (w *WindowWidget) HasFocus() bool {
	return w.getState().hasFocus
}

// RegisterKeyboardShortcuts adds local (window-level) keyboard shortcuts
// see InputHandler.go.
func (w *WindowWidget) RegisterKeyboardShortcuts(s ...WindowShortcut) *WindowWidget {
	if w.HasFocus() {
		for _, shortcut := range s {
			Context.InputHandler.RegisterKeyboardShortcuts(Shortcut{
				Key:      shortcut.Key,
				Modifier: shortcut.Modifier,
				Callback: shortcut.Callback,
				IsGlobal: LocalShortcut,
			})
		}
	}

	return w
}

func (w *WindowWidget) getStateID() string {
	return fmt.Sprintf("%s_windowState", w.title)
}

// returns window state.
func (w *WindowWidget) getState() (state *windowState) {
	if state = GetState[windowState](Context, w.getStateID()); state == nil {
		state = &windowState{}
		SetState(Context, w.getStateID(), state)
	}

	return state
}
