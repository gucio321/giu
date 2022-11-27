package giu

import "github.com/AllenDang/cimgui-go"

// ExecCondition represents cimgui.Condition.
type ExecCondition cimgui.ImGuiCond

// cimgui conditions.
const (
	ConditionAlways       ExecCondition = ExecCondition(cimgui.ImGuiCond_Always)
	ConditionOnce         ExecCondition = ExecCondition(cimgui.ImGuiCond_Once)
	ConditionFirstUseEver ExecCondition = ExecCondition(cimgui.ImGuiCond_FirstUseEver)
	ConditionAppearing    ExecCondition = ExecCondition(cimgui.ImGuiCond_Appearing)
)
