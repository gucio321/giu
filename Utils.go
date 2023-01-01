package giu

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"path/filepath"

	"github.com/AllenDang/cimgui-go"
	"github.com/pkg/browser"
)

// LoadImage loads image from file and returns *image.RGBA.
func LoadImage(imgPath string) (*image.RGBA, error) {
	imgFile, err := os.Open(filepath.Clean(imgPath))
	if err != nil {
		return nil, fmt.Errorf("LoadImage: error opening image file %s: %w", imgPath, err)
	}

	defer func() {
		//nolint:govet // we want to reuse this err variable here
		if err := imgFile.Close(); err != nil {
			panic(fmt.Sprintf("error closing image file: %s", imgPath))
		}
	}()

	img, err := png.Decode(imgFile)
	if err != nil {
		return nil, fmt.Errorf("LoadImage: error decoding png image: %w", err)
	}

	return ImageToRgba(img), nil
}

// ImageToRgba converts image.Image to *image.RGBA.
func ImageToRgba(img image.Image) *image.RGBA {
	switch trueImg := img.(type) {
	case *image.RGBA:
		return trueImg
	default:
		rgba := image.NewRGBA(trueImg.Bounds())
		draw.Draw(rgba, trueImg.Bounds(), trueImg, image.Pt(0, 0), draw.Src)

		return rgba
	}
}

// ToVec4Color converts rgba color to cimgui.ImVec4.
func ToVec4Color(col color.Color) cimgui.ImVec4 {
	const mask = 0xffff

	r, g, b, a := col.RGBA()

	return cimgui.ImVec4{
		X: float32(r) / mask,
		Y: float32(g) / mask,
		Z: float32(b) / mask,
		W: float32(a) / mask,
	}
}

// ToVec2 converts image.Point to cimgui.ImVec2.
func ToVec2(pt image.Point) cimgui.ImVec2 {
	return cimgui.ImVec2{
		X: float32(pt.X),
		Y: float32(pt.Y),
	}
}

// Vec4ToRGBA converts cimgui's Vec4 to golang rgba color.
func Vec4ToRGBA(vec4 cimgui.ImVec4) color.RGBA {
	return color.RGBA{
		R: uint8(vec4.X * 255),
		G: uint8(vec4.Y * 255),
		B: uint8(vec4.Z * 255),
		A: uint8(vec4.W * 255),
	}
}

// Update updates giu app
// it is done by default after each frame.
// However because frames stops rendering, when no user
// action is done, it may be necessary to
// Update ui manually at some point.
func Update() {
	if Context.isAlive {
		Context.platform.Update()
	}
}

// GetCursorScreenPos returns cimgui drawing cursor on the screen.
func GetCursorScreenPos() image.Point {
	pos := cimgui.CursorScreenPos()
	return image.Pt(int(pos.X), int(pos.Y))
}

// SetCursorScreenPos sets cimgui drawing cursor on the screen.
func SetCursorScreenPos(pos image.Point) {
	cimgui.SetCursorScreenPos(cimgui.ImVec2{X: float32(pos.X), Y: float32(pos.Y)})
}

// GetCursorPos gets cimgui drawing cursor inside of current window.
func GetCursorPos() image.Point {
	pos := cimgui.CursorPos()
	return image.Pt(int(pos.X), int(pos.Y))
}

// SetCursorPos sets cimgui drawing cursor inside of current window.
func SetCursorPos(pos image.Point) {
	cimgui.SetCursorPos(cimgui.ImVec2{X: float32(pos.X), Y: float32(pos.Y)})
}

// GetMousePos returns mouse position.
func GetMousePos() image.Point {
	pos := cimgui.MousePos()
	return image.Pt(int(pos.X), int(pos.Y))
}

// GetAvailableRegion returns region available for rendering.
// it is always WindowSize-WindowPadding*2.
func GetAvailableRegion() (width, height float32) {
	region := cimgui.ContentRegionAvail()
	return region.X, region.Y
}

// CalcTextSize calls CalcTextSizeV(text, false, -1).
func CalcTextSize(text string) (width, height float32) {
	return CalcTextSizeV(text, false, -1)
}

// CalcTextSizeV calculates text dimensions.
func CalcTextSizeV(text string, hideAfterDoubleHash bool, wrapWidth float32) (w, h float32) {
	size := cimgui.CalcTextSize(text, hideAfterDoubleHash, wrapWidth)
	return size.X, size.Y
}

// SetNextWindowSize sets size of the next window.
func SetNextWindowSize(width, height float32) {
	cimgui.SetNextWindowSize(cimgui.ImVec2{X: width, Y: height})
}

// ExecCondition represents cimgui.Cond.
type ExecCondition cimgui.Cond

// cimgui conditions.
const (
	ConditionAlways       ExecCondition = ExecCondition(cimgui.Cond_Always)
	ConditionOnce         ExecCondition = ExecCondition(cimgui.Cond_Once)
	ConditionFirstUseEver ExecCondition = ExecCondition(cimgui.Cond_FirstUseEver)
	ConditionAppearing    ExecCondition = ExecCondition(cimgui.Cond_Appearing)
)

// SetNextWindowPos sets position of next window.
func SetNextWindowPos(x, y float32) {
	cimgui.SetNextWindowPos(cimgui.ImVec2{X: x, Y: y})
}

// SetNextWindowSizeV does similar to SetNextWIndowSize but allows to specify cimgui.Cond.
func SetNextWindowSizeV(width, height float32, condition ExecCondition) {
	cimgui.SetNextWindowSizeV(
		cimgui.ImVec2{
			X: width,
			Y: height,
		},
		cimgui.Cond(condition),
	)
}

// SetItemDefaultFocus set the item focused by default.
func SetItemDefaultFocus() {
	cimgui.SetItemDefaultFocus()
}

// SetKeyboardFocusHere sets keyboard focus at *NEXT* widget.
func SetKeyboardFocusHere() {
	SetKeyboardFocusHereV(0)
}

// SetKeyboardFocusHereV sets keyboard on the next widget. Use positive 'offset' to access sub components of a multiple component widget. Use -1 to access previous widget.
func SetKeyboardFocusHereV(i int) {
	cimgui.SetKeyboardFocusHereV(i)
}

// PushClipRect pushes a clipping rectangle for both ImGui logic (hit-testing etc.) and low-level ImDrawList rendering.
func PushClipRect(clipRectMin, clipRectMax image.Point, intersectWithClipRect bool) {
	cimgui.PushClipRect(ToVec2(clipRectMin), ToVec2(clipRectMax), intersectWithClipRect)
}

// PopClipRect should be called to end PushClipRect.
func PopClipRect() {
	cimgui.PopClipRect()
}

// Assert checks if cond. If not cond, it alls golang panic.
func Assert(cond bool, t, method, msg string, args ...any) {
	if !cond {
		fatal(t, method, msg, args...)
	}
}

func fatal(widgetName, method, message string, args ...any) {
	if widgetName != "" {
		widgetName = fmt.Sprintf("(*%s)", widgetName)
	}

	log.Panicf("giu: %s.%s: %s", widgetName, method, fmt.Sprintf(message, args...))
}

// OpenURL opens `url` in default browser.
func OpenURL(url string) {
	if err := browser.OpenURL(url); err != nil {
		log.Printf("Error opening %s: %v", url, err)
	}
}

// ColorToUint converts GO color into Uint32 color
// it is 0xRRGGBBAA
func ColorToUint(col color.Color) uint32 {
	r, g, b, a := col.RGBA()
	mask := uint32(0xff)
	return r&mask<<24 + g&mask<<16 + b&mask<<8 + a&mask
}

// UintToColor converts uint32 of form 0xRRGGBB into color.RGBA
func UintToColor(col uint32) *color.RGBA {
	mask := 0xff
	r := byte(col >> 24 & uint32(mask))
	g := byte(col >> 16 & uint32(mask))
	b := byte(col >> 8 & uint32(mask))
	a := byte(col >> 0 & uint32(mask))
	return &color.RGBA{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}
