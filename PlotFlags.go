package giu

import "github.com/AllenDang/cimgui-go"

// PlotFlags represents cimgui.ImPlotFlags.
type PlotFlags cimgui.ImPlotFlags

// plot flags.
const (
	PlotFlagsNone        = PlotFlags(cimgui.ImPlotFlags_None)
	PlotFlagsNoTitle     = PlotFlags(cimgui.ImPlotFlags_NoTitle)
	PlotFlagsNoLegend    = PlotFlags(cimgui.ImPlotFlags_NoLegend)
	PlotFlagsNoMenus     = PlotFlags(cimgui.ImPlotFlags_NoMenus)
	PlotFlagsNoBoxSelect = PlotFlags(cimgui.ImPlotFlags_NoBoxSelect)
	//PlotFlagsNoMousePos  = PlotFlags(cimgui.ImPlotFlags_NoMousePos)
	//PlotFlagsNoHighlight = PlotFlags(cimgui.ImPlotFlags_NoHighlight)
	PlotFlagsNoChild = PlotFlags(cimgui.ImPlotFlags_NoChild)
	PlotFlagsEqual   = PlotFlags(cimgui.ImPlotFlags_Equal)
	//PlotFlagsYAxis2      = PlotFlags(cimgui.ImPlotFlags_YAxis2)
	//PlotFlagsYAxis3      = PlotFlags(cimgui.ImPlotFlags_YAxis3)
	//PlotFlagsQuery       = PlotFlags(cimgui.ImPlotFlags_Query)
	PlotFlagsCrosshairs = PlotFlags(cimgui.ImPlotFlags_Crosshairs)
	//PlotFlagsAntiAliased = PlotFlags(cimgui.ImPlotFlags_AntiAliased)
	PlotFlagsCanvasOnly = PlotFlags(cimgui.ImPlotFlags_CanvasOnly)
)

// PlotAxisFlags represents imgui.ImPlotAxisFlags.
type PlotAxisFlags cimgui.ImPlotAxisFlags

// plot axis flags.
const (
	PlotAxisFlagsNone         PlotAxisFlags = PlotAxisFlags(cimgui.ImPlotAxisFlags_None)
	PlotAxisFlagsNoLabel      PlotAxisFlags = PlotAxisFlags(cimgui.ImPlotAxisFlags_NoLabel)
	PlotAxisFlagsNoGridLines  PlotAxisFlags = PlotAxisFlags(cimgui.ImPlotAxisFlags_NoGridLines)
	PlotAxisFlagsNoTickMarks  PlotAxisFlags = PlotAxisFlags(cimgui.ImPlotAxisFlags_NoTickMarks)
	PlotAxisFlagsNoTickLabels PlotAxisFlags = PlotAxisFlags(cimgui.ImPlotAxisFlags_NoTickLabels)
	PlotAxisFlagsForeground   PlotAxisFlags = PlotAxisFlags(cimgui.ImPlotAxisFlags_Foreground)
	//PlotAxisFlagsLogScale      PlotAxisFlags = PlotAxisFlags(cimgui.ImPlotAxisFlags_LogScale)
	//PlotAxisFlagsTime          PlotAxisFlags = PlotAxisFlags(cimgui.ImPlotAxisFlags_Time)
	PlotAxisFlagsInvert        PlotAxisFlags = PlotAxisFlags(cimgui.ImPlotAxisFlags_Invert)
	PlotAxisFlagsNoInitialFit  PlotAxisFlags = PlotAxisFlags(cimgui.ImPlotAxisFlags_NoInitialFit)
	PlotAxisFlagsAutoFit       PlotAxisFlags = PlotAxisFlags(cimgui.ImPlotAxisFlags_AutoFit)
	PlotAxisFlagsRangeFit      PlotAxisFlags = PlotAxisFlags(cimgui.ImPlotAxisFlags_RangeFit)
	PlotAxisFlagsLockMin       PlotAxisFlags = PlotAxisFlags(cimgui.ImPlotAxisFlags_LockMin)
	PlotAxisFlagsLockMax       PlotAxisFlags = PlotAxisFlags(cimgui.ImPlotAxisFlags_LockMax)
	PlotAxisFlagsLock          PlotAxisFlags = PlotAxisFlags(cimgui.ImPlotAxisFlags_Lock)
	PlotAxisFlagsNoDecorations PlotAxisFlags = PlotAxisFlags(cimgui.ImPlotAxisFlags_NoDecorations)
)
