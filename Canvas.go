package giu

import (
	"image"
	"image/color"

	"github.com/AllenDang/cimgui-go"
)

// Canvas represents cimgui.ImDrawList
// from cimgui.h:
//
//	A single draw command list (generally one per window,
//	conceptually you may see this as a dynamic "mesh" builder)
//
// for more details and use cases see examples/canvas.
// NOTE: GetCanvas() method returns a window-level canvas, however
// you can convert any cimgui.ImDrawList into this type.
// The best example could be cimgui.GetPlotDrawList()
//
//	c := &Canvas{cimgui.GetXXXDrawList()}
type Canvas struct {
	DrawList cimgui.ImDrawList
}

// GetCanvas returns current draw list (for current window).
// it will fail if called out of window's layout.
func GetCanvas() *Canvas {
	return &Canvas{
		DrawList: cimgui.GetWindowDrawList(),
	}
}

// AddLine draws a line (from p1 to p2).
func (c *Canvas) AddLine(p1, p2 image.Point, col color.Color, thickness float32) {
	c.DrawList.AddLineV(ToVec2(p1), ToVec2(p2), ColorToUint(col), thickness)
}

// DrawFlags represents cimgui.DrawFlags.
type DrawFlags cimgui.DrawFlags

// draw flags enum:.
const (
	DrawFlagsNone DrawFlags = DrawFlags(cimgui.DrawFlags_None)
	// PathStroke(), AddPolyline(): specify that shape should be closed (note: this is always == 1 for legacy reasons).
	DrawFlagsClosed DrawFlags = DrawFlags(cimgui.DrawFlags_Closed)
	// AddRect(), AddRectFilled(), PathRect(): enable rounding top-left corner only (when rounding > 0.0f, we default to all corners).
	// Was 0x01.
	DrawFlagsRoundCornersTopLeft DrawFlags = DrawFlags(cimgui.DrawFlags_RoundCornersTopLeft)
	// AddRect(), AddRectFilled(), PathRect(): enable rounding top-right corner only (when rounding > 0.0f, we default to all corners).
	// Was 0x02.
	DrawFlagsRoundCornersTopRight DrawFlags = DrawFlags(cimgui.DrawFlags_RoundCornersTopRight)
	// AddRect(), AddRectFilled(), PathRect(): enable rounding bottom-left corner only (when rounding > 0.0f, we default to all corners).
	// Was 0x04.
	DrawFlagsRoundCornersBottomLeft DrawFlags = DrawFlags(cimgui.DrawFlags_RoundCornersBottomLeft)
	// AddRect(), AddRectFilled(), PathRect(): enable rounding bottom-right corner only (when rounding > 0.0f,
	// we default to all corners). Wax 0x08.
	DrawFlagsRoundCornersBottomRight DrawFlags = DrawFlags(cimgui.DrawFlags_RoundCornersBottomRight)
	// AddRect(), AddRectFilled(), PathRect(): disable rounding on all corners (when rounding > 0.0f). This is NOT zero, NOT an implicit flag!
	DrawFlagsRoundCornersNone   DrawFlags = DrawFlags(cimgui.DrawFlags_RoundCornersNone)
	DrawFlagsRoundCornersTop    DrawFlags = DrawFlags(cimgui.DrawFlags_RoundCornersTop)
	DrawFlagsRoundCornersBottom DrawFlags = DrawFlags(cimgui.DrawFlags_RoundCornersBottom)
	DrawFlagsRoundCornersLeft   DrawFlags = DrawFlags(cimgui.DrawFlags_RoundCornersLeft)
	DrawFlagsRoundCornersRight  DrawFlags = DrawFlags(cimgui.DrawFlags_RoundCornersRight)
	DrawFlagsRoundCornersAll    DrawFlags = DrawFlags(cimgui.DrawFlags_RoundCornersAll)

	// Default to ALL corners if none of the RoundCornersXX flags are specified.
	DrawFlagsRoundCornersDefault DrawFlags = DrawFlags(cimgui.DrawFlags_RoundCornersDefault)
	DrawFlagsRoundCornersMask    DrawFlags = DrawFlags(cimgui.DrawFlags_RoundCornersMask)
)

// AddRect draws a rectangle.
func (c *Canvas) AddRect(pMin, pMax image.Point, col color.Color, rounding float32, roundingCorners DrawFlags, thickness float32) {
	c.DrawList.AddRectV(ToVec2(pMin), ToVec2(pMax), ColorToUint(col), rounding, cimgui.DrawFlags(roundingCorners), thickness)
}

// AddRectFilled draws a rectangle filled with `col`.
func (c *Canvas) AddRectFilled(pMin, pMax image.Point, col color.Color, rounding float32, roundingCorners DrawFlags) {
	c.DrawList.AddRectFilledV(ToVec2(pMin), ToVec2(pMax), ColorToUint(col), rounding, cimgui.DrawFlags(roundingCorners))
}

// AddText draws text.
func (c *Canvas) AddText(pos image.Point, col color.Color, text string) {
	c.DrawList.AddText_Vec2(ToVec2(pos), ColorToUint(col), Context.FontAtlas.RegisterString(text))
}

// AddBezierCubic draws bezier cubic.
func (c *Canvas) AddBezierCubic(pos0, cp0, cp1, pos1 image.Point, col color.Color, thickness float32, numSegments int32) {
	c.DrawList.AddBezierCubicV(ToVec2(pos0), ToVec2(cp0), ToVec2(cp1), ToVec2(pos1), ColorToUint(col), thickness, numSegments)
}

// AddTriangle draws a triangle.
func (c *Canvas) AddTriangle(p1, p2, p3 image.Point, col color.Color, thickness float32) {
	c.DrawList.AddTriangleV(ToVec2(p1), ToVec2(p2), ToVec2(p3), ColorToUint(col), thickness)
}

// AddTriangleFilled draws a filled triangle.
func (c *Canvas) AddTriangleFilled(p1, p2, p3 image.Point, col color.Color) {
	c.DrawList.AddTriangleFilled(ToVec2(p1), ToVec2(p2), ToVec2(p3), ColorToUint(col))
}

// AddCircle draws a circle.
func (c *Canvas) AddCircle(center image.Point, radius float32, col color.Color, segments int32, thickness float32) {
	c.DrawList.AddCircleV(ToVec2(center), radius, ColorToUint(col), segments, thickness)
}

// AddCircleFilled draws a filled circle.
func (c *Canvas) AddCircleFilled(center image.Point, radius float32, col color.Color) {
	c.DrawList.AddCircleFilled(ToVec2(center), radius, ColorToUint(col))
}

// AddQuad draws a quad.
func (c *Canvas) AddQuad(p1, p2, p3, p4 image.Point, col color.Color, thickness float32) {
	c.DrawList.AddQuadV(ToVec2(p1), ToVec2(p2), ToVec2(p3), ToVec2(p4), ColorToUint(col), thickness)
}

// AddQuadFilled draws a filled quad.
func (c *Canvas) AddQuadFilled(p1, p2, p3, p4 image.Point, col color.Color) {
	c.DrawList.AddQuadFilled(ToVec2(p1), ToVec2(p2), ToVec2(p3), ToVec2(p4), ColorToUint(col))
}

// Stateful path API, add points then finish with PathFillConvex() or PathStroke().

func (c *Canvas) PathClear() {
	c.DrawList.PathClear()
}

func (c *Canvas) PathLineTo(pos image.Point) {
	c.DrawList.PathLineTo(ToVec2(pos))
}

func (c *Canvas) PathLineToMergeDuplicate(pos image.Point) {
	c.DrawList.PathLineToMergeDuplicate(ToVec2(pos))
}

func (c *Canvas) PathFillConvex(col color.Color) {
	c.DrawList.PathFillConvex(ColorToUint(col))
}

func (c *Canvas) PathStroke(col color.Color, flags DrawFlags, thickness float32) {
	c.DrawList.PathStrokeV(ColorToUint(col), cimgui.DrawFlags(flags), thickness)
}

func (c *Canvas) PathArcTo(center image.Point, radius, min, max float32, numSegments int32) {
	c.DrawList.PathArcToV(ToVec2(center), radius, min, max, numSegments)
}

func (c *Canvas) PathArcToFast(center image.Point, radius float32, min12, max12 int32) {
	c.DrawList.PathArcToFast(ToVec2(center), radius, min12, max12)
}

func (c *Canvas) PathBezierCubicCurveTo(p1, p2, p3 image.Point, numSegments int32) {
	c.DrawList.PathBezierCubicCurveToV(ToVec2(p1), ToVec2(p2), ToVec2(p3), numSegments)
}

func (c *Canvas) AddImage(texture *Texture, pMin, pMax image.Point) {
	c.DrawList.AddImage(texture.id, ToVec2(pMin), ToVec2(pMax))
}

func (c *Canvas) AddImageV(texture *Texture, pMin, pMax, uvMin, uvMax image.Point, col color.Color) {
	c.DrawList.AddImageV(texture.id, ToVec2(pMin), ToVec2(pMax), ToVec2(uvMin), ToVec2(uvMax), ColorToUint(col))
}
