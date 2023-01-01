package giu

import (
	"image/color"

	"github.com/AllenDang/cimgui-go"
)

// You may want to use styles in order to make your app looking more beautiful.
// You have two ways to apply style to a widget:
// 1. Use the StyleSetter e.g.:
//    ```golang
//   	giu.Style().
//  		SetStyle(giu.StyleVarWindowPadding, cimgui.ImVec2{10, 10})
//  		SetStyleFloat(giu.StyleVarGrabRounding, 5)
//  		SetColor(giu.StyleColorButton, colornames.Red).
// 			To(/*your widgets here*/),
//   ```
// NOTE/TODO: style variables could be Vec2 or float32 for details see comments
// 2. use PushStyle/PushStyleColor in giu.Custom widget
//    NOTE: remember about calling PopStyle/PopStyleColor at the end of styled section!
//    example:
//    ```golang
// 	  	giu.Custom(func() {
// 		  	cimgui.PushStyleVarFlot(giu.StyleVarFrameRounding, 2)
//    	}),
// 		/*your widgets here*/
//   	giu.Custom(func() {
//   		cimgui.PopStyleVar()
//   	}),
//    ```
// below, you can find a few giu wrappers like PushItemSpacing PushColorFrameBG that
// can be used in a similar way as shown above but without specifying style ID.
//
// See also:
// - examples/setstyle for code example
// - StyleIDs.go for list of all style/color IDs
// - StyleSetter.go for user-friendly giu api for styles

// PushFont sets font to "font"
// NOTE: PopFont has to be called
// NOTE: Don't use PushFont. use StyleSetter instead.
func PushFont(font *FontInfo) bool {
	if font == nil {
		return false
	}

	if f, ok := Context.FontAtlas.extraFontMap[font.String()]; ok {
		cimgui.PushFont(*f)
		return true
	}

	return false
}

// PopFont pops the font (should be called after PushFont).
func PopFont() {
	cimgui.PopFont()
}

// PushStyleColor wraps cimgui.PushStyleColor_Vec4
// NOTE: don't forget to call PopStyleColor()!
func PushStyleColor(id StyleColorID, col color.Color) {
	cimgui.PushStyleColor_Vec4(cimgui.Col(id), ToVec4Color(col))
}

// PushColorText calls PushStyleColor(StyleColorText,...)
// NOTE: don't forget to call PopStyleColor()!
func PushColorText(col color.Color) {
	cimgui.PushStyleColor_Vec4(cimgui.Col_Text, ToVec4Color(col))
}

// PushColorTextDisabled calls PushStyleColor(StyleColorTextDisabled,...)
// NOTE: don't forget to call PopStyleColor()!
func PushColorTextDisabled(col color.Color) {
	cimgui.PushStyleColor_Vec4(cimgui.Col_TextDisabled, ToVec4Color(col))
}

// PushColorWindowBg calls PushStyleColor(StyleColorWindowBg,...)
// NOTE: don't forget to call PopStyleColor()!
func PushColorWindowBg(col color.Color) {
	cimgui.PushStyleColor_Vec4(cimgui.Col_WindowBg, ToVec4Color(col))
}

// PushColorFrameBg calls PushStyleColor(StyleColorFrameBg,...)
// NOTE: don't forget to call PopStyleColor()!
func PushColorFrameBg(col color.Color) {
	cimgui.PushStyleColor_Vec4(cimgui.Col_FrameBg, ToVec4Color(col))
}

// PushColorButton calls PushStyleColor(StyleColorButton,...)
// NOTE: don't forget to call PopStyleColor()!
func PushColorButton(col color.Color) {
	cimgui.PushStyleColor_Vec4(cimgui.Col_Button, ToVec4Color(col))
}

// PushColorButtonHovered calls PushStyleColor(StyleColorButtonHovered,...)
// NOTE: don't forget to call PopStyleColor()!
func PushColorButtonHovered(col color.Color) {
	cimgui.PushStyleColor_Vec4(cimgui.Col_ButtonHovered, ToVec4Color(col))
}

// PushColorButtonActive calls PushStyleColor(StyleColorButtonActive,...)
// NOTE: don't forget to call PopStyleColor()!
func PushColorButtonActive(col color.Color) {
	cimgui.PushStyleColor_Vec4(cimgui.Col_ButtonActive, ToVec4Color(col))
}

// PushWindowPadding calls PushStyleVar(StyleWindowPadding,...)
func PushWindowPadding(width, height float32) {
	cimgui.PushStyleVar_Vec2(cimgui.StyleVar_WindowPadding, cimgui.ImVec2{X: width, Y: height})
}

// PushFramePadding calls PushStyleVar(StyleFramePadding,...)
func PushFramePadding(width, height float32) {
	cimgui.PushStyleVar_Vec2(cimgui.StyleVar_FramePadding, cimgui.ImVec2{X: width, Y: height})
}

// PushItemSpacing calls PushStyleVar(StyleVarItemSpacing,...)
func PushItemSpacing(width, height float32) {
	cimgui.PushStyleVar_Vec2(cimgui.StyleVar_ItemSpacing, cimgui.ImVec2{X: width, Y: height})
}

// PushButtonTextAlign sets alignment for button text. Defaults to (0.0f,0.5f) for left-aligned,vertically centered.
func PushButtonTextAlign(width, height float32) {
	cimgui.PushStyleVar_Vec2(cimgui.StyleVar_ButtonTextAlign, cimgui.ImVec2{X: width, Y: height})
}

// PushSelectableTextAlign sets alignment for selectable text. Defaults to (0.0f,0.5f) for left-aligned,vertically centered.
func PushSelectableTextAlign(width, height float32) {
	cimgui.PushStyleVar_Vec2(cimgui.StyleVar_SelectableTextAlign, cimgui.ImVec2{X: width, Y: height})
}

// PopStyle should be called to stop applying style.
// It should be called as many times, as you called PushStyle...
// NOTE: If you don't call PopStyle cimgui will panic.
func PopStyle() {
	cimgui.PopStyleVar()
}

// PopStyleV does similarly to PopStyle, but allows to specify number
// of styles you're going to pop.
func PopStyleV(count int) {
	cimgui.PopStyleVarV(int32(count))
}

// PopStyleColor is used to stop applying colors styles.
// It should be called after each PushStyleColor... (for each push)
// If PopStyleColor wasn't called after PushColor... or was called
// improperly, cimgui will panic.
func PopStyleColor() {
	cimgui.PopStyleColor()
}

// PopStyleColorV does similar to PopStyleColor, but allows to specify
// how much style colors would you like to pop.
func PopStyleColorV(count int) {
	cimgui.PopStyleColorV(int32(count))
}

// AlignTextToFramePadding vertically aligns upcoming text baseline to
// FramePadding.y so that it will align properly to regularly framed
// items. Call if you have text on a line before a framed item.
func AlignTextToFramePadding() {
	cimgui.AlignTextToFramePadding()
}

// PushItemWidth sets following item's widths
// NOTE: don't forget to call PopItemWidth! If you don't do so, cimgui
// will panic.
func PushItemWidth(width float32) {
	cimgui.PushItemWidth(width)
}

// PopItemWidth should be called to stop applying PushItemWidth effect
// If it isn't called cimgui will panic.
func PopItemWidth() {
	cimgui.PopItemWidth()
}

// PushTextWrapPos adds the position, where the text should be wrapped.
// use PushTextWrapPos, render text. If text reaches frame end,
// rendering will be continued at the start pos in line below.
// NOTE: Don't forget to call PopWrapTextPos
// NOTE: it is done automatically in LabelWidget (see (*LabelWidget).Wrapped()).
func PushTextWrapPos() {
	cimgui.PushTextWrapPos()
}

// PopTextWrapPos should be called as many times as PushTextWrapPos
// on each frame.
func PopTextWrapPos() {
	cimgui.PopTextWrapPos()
}

// MouseCursorType represents a type (layout) of mouse cursor.
type MouseCursorType cimgui.MouseCursor

// cursor types.
const (
	// MouseCursorNone no mouse cursor.
	MouseCursorNone MouseCursorType = cimgui.MouseCursor_None
	// MouseCursorArrow standard arrow mouse cursor.
	MouseCursorArrow MouseCursorType = cimgui.MouseCursor_Arrow
	// MouseCursorTextInput when hovering over InputText, etc.
	MouseCursorTextInput MouseCursorType = cimgui.MouseCursor_TextInput
	// MouseCursorResizeAll (Unused by cimgui functions).
	MouseCursorResizeAll MouseCursorType = cimgui.MouseCursor_ResizeAll
	// MouseCursorResizeNS when hovering over an horizontal border.
	MouseCursorResizeNS MouseCursorType = cimgui.MouseCursor_ResizeNS
	// MouseCursorResizeEW when hovering over a vertical border or a column.
	MouseCursorResizeEW MouseCursorType = cimgui.MouseCursor_ResizeEW
	// MouseCursorResizeNESW when hovering over the bottom-left corner of a window.
	MouseCursorResizeNESW MouseCursorType = cimgui.MouseCursor_ResizeNESW
	// MouseCursorResizeNWSE when hovering over the bottom-right corner of a window.
	MouseCursorResizeNWSE MouseCursorType = cimgui.MouseCursor_ResizeNWSE
	// MouseCursorHand (Unused by cimgui functions. Use for e.g. hyperlinks).
	MouseCursorHand  MouseCursorType = cimgui.MouseCursor_Hand
	MouseCursorCount MouseCursorType = cimgui.MouseCursor_COUNT
)

// SetMouseCursor sets mouse cursor layout.
func SetMouseCursor(cursor MouseCursorType) {
	cimgui.SetMouseCursor(cimgui.MouseCursor(cursor))
}

// GetWindowPadding returns window padding.
func GetWindowPadding() (x, y float32) {
	vec2 := cimgui.GetStyle().GetWindowPadding()
	return vec2.X, vec2.Y
}

// GetItemSpacing returns current item spacing.
func GetItemSpacing() (w, h float32) {
	vec2 := cimgui.GetStyle().GetItemSpacing()
	return vec2.X, vec2.Y
}

// GetItemInnerSpacing returns current item inner spacing.
func GetItemInnerSpacing() (w, h float32) {
	vec2 := cimgui.GetStyle().GetItemInnerSpacing()
	return vec2.X, vec2.Y
}

// GetFramePadding returns current frame padding.
func GetFramePadding() (x, y float32) {
	vec2 := cimgui.GetStyle().GetFramePadding()
	return vec2.X, vec2.Y
}
