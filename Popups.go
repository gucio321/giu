package giu

import (
	"github.com/AllenDang/cimgui-go"
)

// OpenPopup opens a popup with specified id.
// NOTE: you need to build this popup first (see Pop(Modal)Widget).
func OpenPopup(name string) {
	cimgui.OpenPopup_Str(name)
}

// CloseCurrentPopup closes currently opened popup.
// If no popups opened, no action will be taken.
func CloseCurrentPopup() {
	cimgui.CloseCurrentPopup()
}

var _ Widget = &PopupWidget{}

// PopupWidget  is a window which appears next to the mouse cursor.
// For instance it is used to display color palette in ColorSelectWidget.
type PopupWidget struct {
	name   string
	flags  WindowFlags
	layout Layout
}

// Popup creates new popup widget.
func Popup(name string) *PopupWidget {
	return &PopupWidget{
		name:   Context.FontAtlas.RegisterString(name),
		flags:  0,
		layout: nil,
	}
}

// Flags sets popup's flags.
func (p *PopupWidget) Flags(flags WindowFlags) *PopupWidget {
	p.flags = flags
	return p
}

// Layout sets popup's layout.
func (p *PopupWidget) Layout(widgets ...Widget) *PopupWidget {
	p.layout = Layout(widgets)
	return p
}

// Build implements Widget interface.
func (p *PopupWidget) Build() {
	if cimgui.BeginPopupV(p.name, cimgui.WindowFlags(p.flags)) {
		p.layout.Build()
		cimgui.EndPopup()
	}
}

var _ Widget = &PopupModalWidget{}

// PopupModalWidget is a popup window that block every interactions behind it, cannot be closed by
// user, adds a dimming background, has a title bar.
type PopupModalWidget struct {
	name   string
	open   *bool
	flags  WindowFlags
	layout Layout
}

// PopupModal creates new popup modal widget.
func PopupModal(name string) *PopupModalWidget {
	return &PopupModalWidget{
		name:   Context.FontAtlas.RegisterString(name),
		open:   nil,
		flags:  WindowFlagsNoResize,
		layout: nil,
	}
}

// IsOpen allows to control popup's state
// NOTE: changing opens' value will not result in changing popup's state
// if OpenPopup(...) wasn't called!
func (p *PopupModalWidget) IsOpen(open *bool) *PopupModalWidget {
	p.open = open
	return p
}

// Flags allows to specify popup's flags.
func (p *PopupModalWidget) Flags(flags WindowFlags) *PopupModalWidget {
	p.flags = flags
	return p
}

// Layout sets layout.
func (p *PopupModalWidget) Layout(widgets ...Widget) *PopupModalWidget {
	p.layout = Layout(widgets)
	return p
}

// Build implements Widget interface.
func (p *PopupModalWidget) Build() {
	if cimgui.BeginPopupModalV(p.name, p.open, cimgui.WindowFlags(p.flags)) {
		p.layout.Build()
		cimgui.EndPopup()
	}
}
