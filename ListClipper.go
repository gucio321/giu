package giu

import "github.com/AllenDang/cimgui-go/imgui"

var _ Widget = &ListClipperWrapper{}

// ListClipperWrapper is a ImGuiListClipper implementation.
// it can be used to display a large, vertical list of items and
// avoid rendering them.
type ListClipperWrapper struct {
	layout Layout
}

// ListClipper creates list clipper.
func ListClipper() *ListClipperWrapper {
	return &ListClipperWrapper{}
}

// Layout sets layout for list clipper.
func (l *ListClipperWrapper) Layout(layout ...Widget) *ListClipperWrapper {
	l.layout = layout
	return l
}

// Build implements widget interface.
func (l *ListClipperWrapper) Build() {
	// read all the layout widgets and (eventually) split nested layouts
	var layout Layout

	l.layout.Range(func(w Widget) {
		layout = append(layout, w)
	})

	clipper := imgui.NewListClipper()
	defer clipper.Destroy()

	clipper.Begin(int32(len(layout)))

	for clipper.Step() {
		for i := clipper.DisplayStart(); i < clipper.DisplayEnd(); i++ {
			layout[i].Build()
		}
	}

	clipper.End()
}
