package giu

import "github.com/AllenDang/cimgui-go"

// InputTextFlags represents input text flags.
type InputTextFlags cimgui.InputTextFlags

// input text flags.
const (
	// InputTextFlagsNone sets everything default.
	InputTextFlagsNone InputTextFlags = cimgui.InputTextFlags_None
	// InputTextFlagsCharsDecimal allows 0123456789.+-.
	InputTextFlagsCharsDecimal InputTextFlags = cimgui.InputTextFlags_CharsDecimal
	// InputTextFlagsCharsHexadecimal allow 0123456789ABCDEFabcdef.
	InputTextFlagsCharsHexadecimal InputTextFlags = cimgui.InputTextFlags_CharsHexadecimal
	// InputTextFlagsCharsUppercase turns a..z into A..Z.
	InputTextFlagsCharsUppercase InputTextFlags = cimgui.InputTextFlags_CharsUppercase
	// InputTextFlagsCharsNoBlank filters out spaces, tabs.
	InputTextFlagsCharsNoBlank InputTextFlags = cimgui.InputTextFlags_CharsNoBlank
	// InputTextFlagsAutoSelectAll selects entire text when first taking mouse focus.
	InputTextFlagsAutoSelectAll InputTextFlags = cimgui.InputTextFlags_AutoSelectAll
	// InputTextFlagsEnterReturnsTrue returns 'true' when Enter is pressed (as opposed to when the value was modified).
	InputTextFlagsEnterReturnsTrue InputTextFlags = cimgui.InputTextFlags_EnterReturnsTrue
	// InputTextFlagsCallbackCompletion for callback on pressing TAB (for completion handling).
	InputTextFlagsCallbackCompletion InputTextFlags = cimgui.InputTextFlags_CallbackCompletion
	// InputTextFlagsCallbackHistory for callback on pressing Up/Down arrows (for history handling).
	InputTextFlagsCallbackHistory InputTextFlags = cimgui.InputTextFlags_CallbackHistory
	// InputTextFlagsCallbackAlways for callback on each iteration. User code may query cursor position, modify text buffer.
	InputTextFlagsCallbackAlways InputTextFlags = cimgui.InputTextFlags_CallbackAlways
	// InputTextFlagsCallbackCharFilter for callback on character inputs to replace or discard them.
	// Modify 'EventChar' to replace or discard, or return 1 in callback to discard.
	InputTextFlagsCallbackCharFilter InputTextFlags = cimgui.InputTextFlags_CallbackCharFilter
	// InputTextFlagsAllowTabInput when pressing TAB to input a '\t' character into the text field.
	InputTextFlagsAllowTabInput InputTextFlags = cimgui.InputTextFlags_AllowTabInput
	// InputTextFlagsCtrlEnterForNewLine in multi-line mode, unfocus with Enter, add new line with Ctrl+Enter
	// (default is opposite: unfocus with Ctrl+Enter, add line with Enter).
	InputTextFlagsCtrlEnterForNewLine InputTextFlags = cimgui.InputTextFlags_CtrlEnterForNewLine
	// InputTextFlagsNoHorizontalScroll disables following the cursor horizontally.
	InputTextFlagsNoHorizontalScroll InputTextFlags = cimgui.InputTextFlags_NoHorizontalScroll
	// InputTextFlagsAlwaysInsertMode sets insert mode.

	// InputTextFlagsReadOnly sets read-only mode.
	InputTextFlagsReadOnly InputTextFlags = cimgui.InputTextFlags_ReadOnly
	// InputTextFlagsPassword sets password mode, display all characters as '*'.
	InputTextFlagsPassword InputTextFlags = cimgui.InputTextFlags_Password
	// InputTextFlagsNoUndoRedo disables undo/redo. Note that input text owns the text data while active,
	// if you want to provide your own undo/redo stack you need e.g. to call ClearActiveID().
	InputTextFlagsNoUndoRedo InputTextFlags = cimgui.InputTextFlags_NoUndoRedo
	// InputTextFlagsCharsScientific allows 0123456789.+-*/eE (Scientific notation input).
	InputTextFlagsCharsScientific InputTextFlags = cimgui.InputTextFlags_CharsScientific
)

// WindowFlags represents a window flags (see (*WindowWidget).Flags.
type WindowFlags cimgui.GLFWWindowFlags

// window flags.
const (
	// WindowFlagsNone default = 0.
	WindowFlagsNone WindowFlags = WindowFlags(cimgui.WindowFlags_None)
	// WindowFlagsNoTitleBar disables title-bar.
	WindowFlagsNoTitleBar WindowFlags = WindowFlags(cimgui.WindowFlags_NoTitleBar)
	// WindowFlagsNoResize disables user resizing with the lower-right grip.
	WindowFlagsNoResize WindowFlags = WindowFlags(cimgui.WindowFlags_NoResize)
	// WindowFlagsNoMove disables user moving the window.
	WindowFlagsNoMove WindowFlags = WindowFlags(cimgui.WindowFlags_NoMove)
	// WindowFlagsNoScrollbar disables scrollbars. Window can still scroll with mouse or programmatically.
	WindowFlagsNoScrollbar WindowFlags = WindowFlags(cimgui.WindowFlags_NoScrollbar)
	// WindowFlagsNoScrollWithMouse disables user vertically scrolling with mouse wheel. On child window, mouse wheel
	// will be forwarded to the parent unless NoScrollbar is also set.
	WindowFlagsNoScrollWithMouse WindowFlags = WindowFlags(cimgui.WindowFlags_NoScrollWithMouse)
	// WindowFlagsNoCollapse disables user collapsing window by double-clicking on it.
	WindowFlagsNoCollapse WindowFlags = WindowFlags(cimgui.WindowFlags_NoCollapse)
	// WindowFlagsAlwaysAutoResize resizes every window to its content every frame.
	WindowFlagsAlwaysAutoResize WindowFlags = WindowFlags(cimgui.WindowFlags_AlwaysAutoResize)
	// WindowFlagsNoBackground disables drawing background color (WindowBg, etc.) and outside border. Similar as using
	// SetNextWindowBgAlpha(0.0f).
	WindowFlagsNoBackground WindowFlags = WindowFlags(cimgui.WindowFlags_NoBackground)
	// WindowFlagsNoSavedSettings will never load/save settings in .ini file.
	WindowFlagsNoSavedSettings WindowFlags = WindowFlags(cimgui.WindowFlags_NoSavedSettings)
	// WindowFlagsNoMouseInputs disables catching mouse, hovering test with pass through.
	WindowFlagsNoMouseInputs WindowFlags = WindowFlags(cimgui.WindowFlags_NoMouseInputs)
	// WindowFlagsMenuBar has a menu-bar.
	WindowFlagsMenuBar WindowFlags = WindowFlags(cimgui.WindowFlags_MenuBar)
	// WindowFlagsHorizontalScrollbar allows horizontal scrollbar to appear (off by default). You may use
	// SetNextWindowContentSize(ImVec2(width,0.0f)); prior to calling Begin() to specify width. Read code in cimgui_demo
	// in the "Horizontal Scrolling" section.
	WindowFlagsHorizontalScrollbar WindowFlags = WindowFlags(cimgui.WindowFlags_HorizontalScrollbar)
	// WindowFlagsNoFocusOnAppearing disables taking focus when transitioning from hidden to visible state.
	WindowFlagsNoFocusOnAppearing WindowFlags = WindowFlags(cimgui.WindowFlags_NoFocusOnAppearing)
	// WindowFlagsNoBringToFrontOnFocus disables bringing window to front when taking focus. e.g. clicking on it or
	// programmatically giving it focus.
	WindowFlagsNoBringToFrontOnFocus WindowFlags = WindowFlags(cimgui.WindowFlags_NoBringToFrontOnFocus)
	// WindowFlagsAlwaysVerticalScrollbar always shows vertical scrollbar, even if ContentSize.y < Size.y .
	WindowFlagsAlwaysVerticalScrollbar WindowFlags = WindowFlags(cimgui.WindowFlags_AlwaysVerticalScrollbar)
	// WindowFlagsAlwaysHorizontalScrollbar always shows horizontal scrollbar, even if ContentSize.x < Size.x .
	WindowFlagsAlwaysHorizontalScrollbar WindowFlags = WindowFlags(cimgui.WindowFlags_AlwaysHorizontalScrollbar)
	// WindowFlagsAlwaysUseWindowPadding ensures child windows without border uses style.WindowPadding (ignored by
	// default for non-bordered child windows, because more convenient).
	WindowFlagsAlwaysUseWindowPadding WindowFlags = WindowFlags(cimgui.WindowFlags_AlwaysUseWindowPadding)
	// WindowFlagsNoNavInputs has no gamepad/keyboard navigation within the window.
	WindowFlagsNoNavInputs WindowFlags = WindowFlags(cimgui.WindowFlags_NoNavInputs)
	// WindowFlagsNoNavFocus has no focusing toward this window with gamepad/keyboard navigation
	// (e.g. skipped by CTRL+TAB).
	WindowFlagsNoNavFocus WindowFlags = WindowFlags(cimgui.WindowFlags_NoNavFocus)
	// WindowFlagsUnsavedDocument appends '*' to title without affecting the ID, as a convenience to avoid using the
	// ### operator. When used in a tab/docking context, tab is selected on closure and closure is deferred by one
	// frame to allow code to cancel the closure (with a confirmation popup, etc.) without flicker.
	WindowFlagsUnsavedDocument WindowFlags = WindowFlags(cimgui.WindowFlags_UnsavedDocument)

	// WindowFlagsNoNav combines WindowFlagsNoNavInputs and WindowFlagsNoNavFocus.
	WindowFlagsNoNav WindowFlags = WindowFlags(cimgui.WindowFlags_NoNav)
	// WindowFlagsNoDecoration combines WindowFlagsNoTitleBar, WindowFlagsNoResize, WindowFlagsNoScrollbar and
	// WindowFlagsNoCollapse.
	WindowFlagsNoDecoration WindowFlags = WindowFlags(cimgui.WindowFlags_NoDecoration)
	// WindowFlagsNoInputs combines WindowFlagsNoMouseInputs, WindowFlagsNoNavInputs and WindowFlagsNoNavFocus.
	WindowFlagsNoInputs WindowFlags = WindowFlags(cimgui.WindowFlags_NoInputs)
)

// ComboFlags represents cimgui.ComboFlags.
type ComboFlags cimgui.ComboFlags

// combo flags list.
const (
	// ComboFlagsNone default = 0.
	ComboFlagsNone ComboFlags = cimgui.ComboFlags_None
	// ComboFlagsPopupAlignLeft aligns the popup toward the left by default.
	ComboFlagsPopupAlignLeft ComboFlags = cimgui.ComboFlags_PopupAlignLeft
	// ComboFlagsHeightSmall has max ~4 items visible.
	// Tip: If you want your combo popup to be a specific size you can use SetNextWindowSizeConstraints() prior to calling BeginCombo().
	ComboFlagsHeightSmall ComboFlags = cimgui.ComboFlags_HeightSmall
	// ComboFlagsHeightRegular has max ~8 items visible (default).
	ComboFlagsHeightRegular ComboFlags = cimgui.ComboFlags_HeightRegular
	// ComboFlagsHeightLarge has max ~20 items visible.
	ComboFlagsHeightLarge ComboFlags = cimgui.ComboFlags_HeightLarge
	// ComboFlagsHeightLargest has as many fitting items as possible.
	ComboFlagsHeightLargest ComboFlags = cimgui.ComboFlags_HeightLargest
	// ComboFlagsNoArrowButton displays on the preview box without the square arrow button.
	ComboFlagsNoArrowButton ComboFlags = cimgui.ComboFlags_NoArrowButton
	// ComboFlagsNoPreview displays only a square arrow button.
	ComboFlagsNoPreview ComboFlags = cimgui.ComboFlags_NoPreview
)

// SelectableFlags represents cimgui.SelectableFlags.
type SelectableFlags cimgui.SelectableFlags

// selectable flags list.
const (
	// SelectableFlagsNone default = 0.
	SelectableFlagsNone SelectableFlags = cimgui.SelectableFlags_None
	// SelectableFlagsDontClosePopups makes clicking the selectable not close any parent popup windows.
	SelectableFlagsDontClosePopups SelectableFlags = cimgui.SelectableFlags_DontClosePopups
	// SelectableFlagsSpanAllColumns allows the selectable frame to span all columns (text will still fit in current column).
	SelectableFlagsSpanAllColumns SelectableFlags = cimgui.SelectableFlags_SpanAllColumns
	// SelectableFlagsAllowDoubleClick generates press events on double clicks too.
	SelectableFlagsAllowDoubleClick SelectableFlags = cimgui.SelectableFlags_AllowDoubleClick
	// SelectableFlagsDisabled disallows selection and displays text in a greyed out color.
	SelectableFlagsDisabled SelectableFlags = cimgui.SelectableFlags_Disabled
)

// TabItemFlags represents tab item flags.
type TabItemFlags cimgui.TabItemFlags

// tab item flags list.
const (
	// TabItemFlagsNone default = 0.
	TabItemFlagsNone TabItemFlags = cimgui.TabItemFlags_None
	// TabItemFlagsUnsavedDocument Append '*' to title without affecting the ID, as a convenience to avoid using the
	// ### operator. Also: tab is selected on closure and closure is deferred by one frame to allow code to undo it
	// without flicker.
	TabItemFlagsUnsavedDocument TabItemFlags = cimgui.TabItemFlags_UnsavedDocument
	// TabItemFlagsSetSelected Trigger flag to programmatically make the tab selected when calling BeginTabItem().
	TabItemFlagsSetSelected TabItemFlags = cimgui.TabItemFlags_SetSelected
	// TabItemFlagsNoCloseWithMiddleMouseButton  Disable behavior of closing tabs (that are submitted with
	// p_open != NULL) with middle mouse button. You can still repro this behavior on user's side with if
	// (IsItemHovered() && IsMouseClicked(2)) *p_open = false.
	TabItemFlagsNoCloseWithMiddleMouseButton TabItemFlags = cimgui.TabItemFlags_NoCloseWithMiddleMouseButton
	// TabItemFlagsNoPushID Don't call PushID(tab->ID)/PopID() on BeginTabItem()/EndTabItem().

)

// TabBarFlags represents cimgui.TabBarFlags.
type TabBarFlags cimgui.TabBarFlags

// tab bar flags list.
const (
	// TabBarFlagsNone default = 0.
	TabBarFlagsNone TabBarFlags = cimgui.TabBarFlags_None
	// TabBarFlagsReorderable Allow manually dragging tabs to re-order them + New tabs are appended at the end of list.
	TabBarFlagsReorderable TabBarFlags = cimgui.TabBarFlags_Reorderable
	// TabBarFlagsAutoSelectNewTabs Automatically select new tabs when they appear.
	TabBarFlagsAutoSelectNewTabs TabBarFlags = cimgui.TabBarFlags_AutoSelectNewTabs
	// TabBarFlagsTabListPopupButton Disable buttons to open the tab list popup.
	TabBarFlagsTabListPopupButton TabBarFlags = cimgui.TabBarFlags_TabListPopupButton
	// TabBarFlagsNoCloseWithMiddleMouseButton Disable behavior of closing tabs (that are submitted with p_open != NULL)
	// with middle mouse button. You can still repro this behavior on user's side with if
	// (IsItemHovered() && IsMouseClicked(2)) *p_open = false.
	TabBarFlagsNoCloseWithMiddleMouseButton TabBarFlags = cimgui.TabBarFlags_NoCloseWithMiddleMouseButton
	// TabBarFlagsNoTabListScrollingButtons Disable scrolling buttons (apply when fitting policy is
	// TabBarFlagsFittingPolicyScroll).
	TabBarFlagsNoTabListScrollingButtons TabBarFlags = cimgui.TabBarFlags_NoTabListScrollingButtons
	// TabBarFlagsNoTooltip Disable tooltips when hovering a tab.
	TabBarFlagsNoTooltip TabBarFlags = cimgui.TabBarFlags_NoTooltip
	// TabBarFlagsFittingPolicyResizeDown Resize tabs when they don't fit.
	TabBarFlagsFittingPolicyResizeDown TabBarFlags = cimgui.TabBarFlags_FittingPolicyResizeDown
	// TabBarFlagsFittingPolicyScroll Add scroll buttons when tabs don't fit.
	TabBarFlagsFittingPolicyScroll TabBarFlags = cimgui.TabBarFlags_FittingPolicyScroll
	// TabBarFlagsFittingPolicyMask combines
	// TabBarFlagsFittingPolicyResizeDown and TabBarFlagsFittingPolicyScroll.
	TabBarFlagsFittingPolicyMask TabBarFlags = cimgui.TabBarFlags_FittingPolicyMask
	// TabBarFlagsFittingPolicyDefault alias for TabBarFlagsFittingPolicyResizeDown.
	TabBarFlagsFittingPolicyDefault TabBarFlags = cimgui.TabBarFlags_FittingPolicyDefault
)

// TreeNodeFlags represents tree node widget flags.
type TreeNodeFlags cimgui.TreeNodeFlags

// tree node flags list.
const (
	// TreeNodeFlagsNone default = 0.
	TreeNodeFlagsNone TreeNodeFlags = cimgui.TreeNodeFlags_None
	// TreeNodeFlagsSelected draws as selected.
	TreeNodeFlagsSelected TreeNodeFlags = cimgui.TreeNodeFlags_Selected
	// TreeNodeFlagsFramed draws full colored frame (e.g. for CollapsingHeader).
	TreeNodeFlagsFramed TreeNodeFlags = cimgui.TreeNodeFlags_Framed
	// TreeNodeFlagsAllowItemOverlap hit testing to allow subsequent widgets to overlap this one.
	TreeNodeFlagsAllowItemOverlap TreeNodeFlags = cimgui.TreeNodeFlags_AllowItemOverlap
	// TreeNodeFlagsNoTreePushOnOpen doesn't do a TreePush() when open
	// (e.g. for CollapsingHeader) = no extra indent nor pushing on ID stack.
	TreeNodeFlagsNoTreePushOnOpen TreeNodeFlags = cimgui.TreeNodeFlags_NoTreePushOnOpen
	// TreeNodeFlagsNoAutoOpenOnLog doesn't automatically and temporarily open node when Logging is active
	// (by default logging will automatically open tree nodes).
	TreeNodeFlagsNoAutoOpenOnLog TreeNodeFlags = cimgui.TreeNodeFlags_NoAutoOpenOnLog
	// TreeNodeFlagsDefaultOpen defaults node to be open.
	TreeNodeFlagsDefaultOpen TreeNodeFlags = cimgui.TreeNodeFlags_DefaultOpen
	// TreeNodeFlagsOpenOnDoubleClick needs double-click to open node.
	TreeNodeFlagsOpenOnDoubleClick TreeNodeFlags = cimgui.TreeNodeFlags_OpenOnDoubleClick
	// TreeNodeFlagsOpenOnArrow opens only when clicking on the arrow part.
	// If TreeNodeFlagsOpenOnDoubleClick is also set, single-click arrow or double-click all box to open.
	TreeNodeFlagsOpenOnArrow TreeNodeFlags = cimgui.TreeNodeFlags_OpenOnArrow
	// TreeNodeFlagsLeaf allows no collapsing, no arrow (use as a convenience for leaf nodes).
	TreeNodeFlagsLeaf TreeNodeFlags = cimgui.TreeNodeFlags_Leaf
	// TreeNodeFlagsBullet displays a bullet instead of an arrow.
	TreeNodeFlagsBullet TreeNodeFlags = cimgui.TreeNodeFlags_Bullet
	// TreeNodeFlagsFramePadding uses FramePadding (even for an unframed text node) to
	// vertically align text baseline to regular widget height. Equivalent to calling AlignTextToFramePadding().
	TreeNodeFlagsFramePadding TreeNodeFlags = cimgui.TreeNodeFlags_FramePadding
	// TreeNodeFlagsSpanAvailWidth extends hit box to the right-most edge, even if not framed.
	// This is not the default in order to allow adding other items on the same line.
	// In the future we may refactor the hit system to be front-to-back, allowing natural overlaps
	// and then this can become the default.
	TreeNodeFlagsSpanAvailWidth TreeNodeFlags = cimgui.TreeNodeFlags_SpanAvailWidth
	// TreeNodeFlagsSpanFullWidth extends hit box to the left-most and right-most edges (bypass the indented area).
	TreeNodeFlagsSpanFullWidth TreeNodeFlags = cimgui.TreeNodeFlags_SpanFullWidth
	// TreeNodeFlagsNavLeftJumpsBackHere (WIP) Nav: left direction may move to this TreeNode() from any of its child
	// (items submitted between TreeNode and TreePop).
	TreeNodeFlagsNavLeftJumpsBackHere TreeNodeFlags = cimgui.TreeNodeFlags_NavLeftJumpsBackHere
	// TreeNodeFlagsCollapsingHeader combines TreeNodeFlagsFramed and TreeNodeFlagsNoAutoOpenOnLog.
	TreeNodeFlagsCollapsingHeader TreeNodeFlags = cimgui.TreeNodeFlags_CollapsingHeader
)

// FocusedFlags represents cimgui.FocusedFlags.
type FocusedFlags cimgui.FocusedFlags

// focused flags list.
const (
	FocusedFlagsNone             = cimgui.FocusedFlags_None
	FocusedFlagsChildWindows     = cimgui.FocusedFlags_ChildWindows     // Return true if any children of the window is focused
	FocusedFlagsRootWindow       = cimgui.FocusedFlags_RootWindow       // Test from root window (top most parent of the current hierarchy)
	FocusedFlagsAnyWindow        = cimgui.FocusedFlags_AnyWindow        // Return true if any window is focused. Important: If you are trying to tell how to dispatch your low-level inputs do NOT use this. Use 'io.WantCaptureMouse' instead! Please read the FAQ!
	FocusedFlagsNoPopupHierarchy = cimgui.FocusedFlags_NoPopupHierarchy // Do not consider popup hierarchy (do not treat popup emitter as parent of popup) (when used with ChildWindows or RootWindow)
	// FocusedFlagsDockHierarchy               = 1 << 4   // Consider docking hierarchy (treat dockspace host as parent of docked window) (when used with ChildWindows or RootWindow).
	FocusedFlagsRootAndChildWindows = cimgui.FocusedFlags_RootAndChildWindows
)

// HoveredFlags represents a hovered flags.
type HoveredFlags cimgui.HoveredFlags

// hovered flags list.
const (
	// HoveredFlagsNone Return true if directly over the item/window, not obstructed by another window,
	// not obstructed by an active popup or modal blocking inputs under them.
	HoveredFlagsNone HoveredFlags = cimgui.HoveredFlags_None
	// HoveredFlagsChildWindows IsWindowHovered() only: Return true if any children of the window is hovered.
	HoveredFlagsChildWindows HoveredFlags = cimgui.HoveredFlags_ChildWindows
	// HoveredFlagsRootWindow IsWindowHovered() only: Test from root window (top most parent of the current hierarchy).
	HoveredFlagsRootWindow HoveredFlags = cimgui.HoveredFlags_RootWindow
	// HoveredFlagsAnyWindow IsWindowHovered() only: Return true if any window is hovered.
	HoveredFlagsAnyWindow HoveredFlags = cimgui.HoveredFlags_AnyWindow
	// HoveredFlagsAllowWhenBlockedByPopup Return true even if a popup window is normally blocking access to this item/window.
	HoveredFlagsAllowWhenBlockedByPopup HoveredFlags = cimgui.HoveredFlags_AllowWhenBlockedByPopup
	// HoveredFlagsAllowWhenBlockedByActiveItem Return true even if an active item is blocking access to this item/window.
	// Useful for Drag and Drop patterns.
	HoveredFlagsAllowWhenBlockedByActiveItem HoveredFlags = cimgui.HoveredFlags_AllowWhenBlockedByActiveItem
	// HoveredFlagsAllowWhenOverlapped Return true even if the position is overlapped by another window.
	HoveredFlagsAllowWhenOverlapped HoveredFlags = cimgui.HoveredFlags_AllowWhenOverlapped
	// HoveredFlagsAllowWhenDisabled Return true even if the item is disabled.
	HoveredFlagsAllowWhenDisabled HoveredFlags = cimgui.HoveredFlags_AllowWhenDisabled
)

// ColorEditFlags for ColorEdit3V(), etc.
type ColorEditFlags int

// list of color edit flags.
const (
// 	// ColorEditFlagsNone default = 0.
// 	ColorEditFlagsNone ColorEditFlags = cimgui.ColorEditFlagsNone
// ColorEditFlagsNoAlpha ignores Alpha component (read 3 components from the input pointer).
// 	ColorEditFlagsNoAlpha ColorEditFlags = cimgui.ColorEditFlagsNoAlpha
// ColorEditFlagsNoPicker disables picker when clicking on colored square.
// 	ColorEditFlagsNoPicker ColorEditFlags = cimgui.ColorEditFlagsNoPicker
// ColorEditFlagsNoOptions disables toggling options menu when right-clicking on inputs/small preview.
// 	ColorEditFlagsNoOptions ColorEditFlags = cimgui.ColorEditFlagsNoOptions
// ColorEditFlagsNoSmallPreview disables colored square preview next to the inputs. (e.g. to show only the inputs).
// 	ColorEditFlagsNoSmallPreview ColorEditFlags = cimgui.ColorEditFlagsNoSmallPreview
// ColorEditFlagsNoInputs disables inputs sliders/text widgets (e.g. to show only the small preview colored square).
// 	ColorEditFlagsNoInputs ColorEditFlags = cimgui.ColorEditFlagsNoInputs
// ColorEditFlagsNoTooltip disables tooltip when hovering the preview.
// 	ColorEditFlagsNoTooltip ColorEditFlags = cimgui.ColorEditFlagsNoTooltip
// ColorEditFlagsNoLabel disables display of inline text label (the label is still forwarded to the tooltip and picker).
// 	ColorEditFlagsNoLabel ColorEditFlags = cimgui.ColorEditFlagsNoLabel
// ColorEditFlagsNoDragDrop disables drag and drop target. ColorButton: disable drag and drop source.
// 	ColorEditFlagsNoDragDrop ColorEditFlags = cimgui.ColorEditFlagsNoDragDrop

// User Options (right-click on widget to change some of them). You can set application defaults using SetColorEditOptions().
// The idea is that you probably don't want to override them in most of your calls, let the user choose and/or call SetColorEditOptions()
// during startup.

// ColorEditFlagsAlphaBar shows vertical alpha bar/gradient in picker.
// 	ColorEditFlagsAlphaBar ColorEditFlags = cimgui.ColorEditFlagsAlphaBar
// ColorEditFlagsAlphaPreview displays preview as a transparent color over a checkerboard, instead of opaque.
// 	ColorEditFlagsAlphaPreview ColorEditFlags = cimgui.ColorEditFlagsAlphaPreview
// ColorEditFlagsAlphaPreviewHalf displays half opaque / half checkerboard, instead of opaque.
// 	ColorEditFlagsAlphaPreviewHalf ColorEditFlags = cimgui.ColorEditFlagsAlphaPreviewHalf
// 	// ColorEditFlagsHDR = (WIP) currently only disable 0.0f..1.0f limits in RGBA edition (note: you probably want to use
// ImGuiColorEditFlags_Float flag as well).
// 	ColorEditFlagsHDR ColorEditFlags = cimgui.ColorEditFlagsHDR
// ColorEditFlagsRGB sets the format as RGB.
// 	ColorEditFlagsRGB ColorEditFlags = cimgui.ColorEditFlagsRGB
// ColorEditFlagsHSV sets the format as HSV.
// 	ColorEditFlagsHSV ColorEditFlags = cimgui.ColorEditFlagsHSV
// ColorEditFlagsHEX sets the format as HEX.
// 	ColorEditFlagsHEX ColorEditFlags = cimgui.ColorEditFlagsHEX
// ColorEditFlagsUint8 _display_ values formatted as 0..255.
// 	ColorEditFlagsUint8 ColorEditFlags = cimgui.ColorEditFlagsUint8
// ColorEditFlagsFloat _display_ values formatted as 0.0f..1.0f floats instead of 0..255 integers. No round-trip of value via integers.
// 	ColorEditFlagsFloat ColorEditFlags = cimgui.ColorEditFlagsFloat
)

// TableFlags represents table flags.
type TableFlags cimgui.TableFlags

// Table flags enum:.
const (
	TableFlagsNone                       TableFlags = TableFlags(cimgui.TableFlags_None)
	TableFlagsResizable                  TableFlags = TableFlags(cimgui.TableFlags_Resizable)
	TableFlagsReorderable                TableFlags = TableFlags(cimgui.TableFlags_Reorderable)
	TableFlagsHideable                   TableFlags = TableFlags(cimgui.TableFlags_Hideable)
	TableFlagsSortable                   TableFlags = TableFlags(cimgui.TableFlags_Sortable)
	TableFlagsNoSavedSettings            TableFlags = TableFlags(cimgui.TableFlags_NoSavedSettings)
	TableFlagsContextMenuInBody          TableFlags = TableFlags(cimgui.TableFlags_ContextMenuInBody)
	TableFlagsRowBg                      TableFlags = TableFlags(cimgui.TableFlags_RowBg)
	TableFlagsBordersInnerH              TableFlags = TableFlags(cimgui.TableFlags_BordersInnerH)
	TableFlagsBordersOuterH              TableFlags = TableFlags(cimgui.TableFlags_BordersOuterH)
	TableFlagsBordersInnerV              TableFlags = TableFlags(cimgui.TableFlags_BordersInnerV)
	TableFlagsBordersOuterV              TableFlags = TableFlags(cimgui.TableFlags_BordersOuterV)
	TableFlagsBordersH                   TableFlags = TableFlags(cimgui.TableFlags_BordersH)
	TableFlagsBordersV                   TableFlags = TableFlags(cimgui.TableFlags_BordersV)
	TableFlagsBordersInner               TableFlags = TableFlags(cimgui.TableFlags_BordersInner)
	TableFlagsBordersOuter               TableFlags = TableFlags(cimgui.TableFlags_BordersOuter)
	TableFlagsBorders                    TableFlags = TableFlags(cimgui.TableFlags_Borders)
	TableFlagsNoBordersInBody            TableFlags = TableFlags(cimgui.TableFlags_NoBordersInBody)
	TableFlagsNoBordersInBodyUntilResize TableFlags = TableFlags(cimgui.TableFlags_NoBordersInBodyUntilResize)
	TableFlagsSizingFixedFit             TableFlags = TableFlags(cimgui.TableFlags_SizingFixedFit)
	TableFlagsSizingFixedSame            TableFlags = TableFlags(cimgui.TableFlags_SizingFixedSame)
	TableFlagsSizingStretchProp          TableFlags = TableFlags(cimgui.TableFlags_SizingStretchProp)
	TableFlagsSizingStretchSame          TableFlags = TableFlags(cimgui.TableFlags_SizingStretchSame)
	TableFlagsNoHostExtendX              TableFlags = TableFlags(cimgui.TableFlags_NoHostExtendX)
	TableFlagsNoHostExtendY              TableFlags = TableFlags(cimgui.TableFlags_NoHostExtendY)
	TableFlagsNoKeepColumnsVisible       TableFlags = TableFlags(cimgui.TableFlags_NoKeepColumnsVisible)
	TableFlagsPreciseWidths              TableFlags = TableFlags(cimgui.TableFlags_PreciseWidths)
	TableFlagsNoClip                     TableFlags = TableFlags(cimgui.TableFlags_NoClip)
	TableFlagsPadOuterX                  TableFlags = TableFlags(cimgui.TableFlags_PadOuterX)
	TableFlagsNoPadOuterX                TableFlags = TableFlags(cimgui.TableFlags_NoPadOuterX)
	TableFlagsNoPadInnerX                TableFlags = TableFlags(cimgui.TableFlags_NoPadInnerX)
	TableFlagsScrollX                    TableFlags = TableFlags(cimgui.TableFlags_ScrollX)
	TableFlagsScrollY                    TableFlags = TableFlags(cimgui.TableFlags_ScrollY)
	TableFlagsSortMulti                  TableFlags = TableFlags(cimgui.TableFlags_SortMulti)
	TableFlagsSortTristate               TableFlags = TableFlags(cimgui.TableFlags_SortTristate)
	TableFlagsSizingMask                 TableFlags = TableFlags(cimgui.TableFlags_SizingMask_)
)

// TableRowFlags represents table row flags.
type TableRowFlags cimgui.TableRowFlags

// table row flags:.
const (
	TableRowFlagsNone TableRowFlags = TableRowFlags(cimgui.TableRowFlags_None)
	// Identify header row (set default background color + width of its contents accounted different for auto column width).
	TableRowFlagsHeaders TableRowFlags = TableRowFlags(cimgui.TableRowFlags_Headers)
)

// TableColumnFlags represents a flags for table column (see (*TableColumnWidget).Flags()).
type TableColumnFlags cimgui.TableColumnFlags

// table column flags list.
const (
	// Input configuration flags.
	TableColumnFlagsNone                 TableColumnFlags = TableColumnFlags(cimgui.TableColumnFlags_None)
	TableColumnFlagsDefaultHide          TableColumnFlags = TableColumnFlags(cimgui.TableColumnFlags_DefaultHide)
	TableColumnFlagsDefaultSort          TableColumnFlags = TableColumnFlags(cimgui.TableColumnFlags_DefaultSort)
	TableColumnFlagsWidthStretch         TableColumnFlags = TableColumnFlags(cimgui.TableColumnFlags_WidthStretch)
	TableColumnFlagsWidthFixed           TableColumnFlags = TableColumnFlags(cimgui.TableColumnFlags_WidthFixed)
	TableColumnFlagsNoResize             TableColumnFlags = TableColumnFlags(cimgui.TableColumnFlags_NoResize)
	TableColumnFlagsNoReorder            TableColumnFlags = TableColumnFlags(cimgui.TableColumnFlags_NoReorder)
	TableColumnFlagsNoHide               TableColumnFlags = TableColumnFlags(cimgui.TableColumnFlags_NoHide)
	TableColumnFlagsNoClip               TableColumnFlags = TableColumnFlags(cimgui.TableColumnFlags_NoClip)
	TableColumnFlagsNoSort               TableColumnFlags = TableColumnFlags(cimgui.TableColumnFlags_NoSort)
	TableColumnFlagsNoSortAscending      TableColumnFlags = TableColumnFlags(cimgui.TableColumnFlags_NoSortAscending)
	TableColumnFlagsNoSortDescending     TableColumnFlags = TableColumnFlags(cimgui.TableColumnFlags_NoSortDescending)
	TableColumnFlagsNoHeaderWidth        TableColumnFlags = TableColumnFlags(cimgui.TableColumnFlags_NoHeaderWidth)
	TableColumnFlagsPreferSortAscending  TableColumnFlags = TableColumnFlags(cimgui.TableColumnFlags_PreferSortAscending)
	TableColumnFlagsPreferSortDescending TableColumnFlags = TableColumnFlags(cimgui.TableColumnFlags_PreferSortDescending)
	TableColumnFlagsIndentEnable         TableColumnFlags = TableColumnFlags(cimgui.TableColumnFlags_IndentEnable)
	TableColumnFlagsIndentDisable        TableColumnFlags = TableColumnFlags(cimgui.TableColumnFlags_IndentDisable)

	// Output status flags read-only via TableGetColumnFlags().
	TableColumnFlagsIsEnabled TableColumnFlags = TableColumnFlags(cimgui.TableColumnFlags_IsEnabled)
	TableColumnFlagsIsVisible TableColumnFlags = TableColumnFlags(cimgui.TableColumnFlags_IsVisible)
	TableColumnFlagsIsSorted  TableColumnFlags = TableColumnFlags(cimgui.TableColumnFlags_IsSorted)
	TableColumnFlagsIsHovered TableColumnFlags = TableColumnFlags(cimgui.TableColumnFlags_IsHovered)

	// [Internal] Combinations and masks.
	TableColumnFlagsWidthMask      TableColumnFlags = TableColumnFlags(cimgui.TableColumnFlags_WidthMask_)
	TableColumnFlagsIndentMask     TableColumnFlags = TableColumnFlags(cimgui.TableColumnFlags_IndentMask_)
	TableColumnFlagsStatusMask     TableColumnFlags = TableColumnFlags(cimgui.TableColumnFlags_StatusMask_)
	TableColumnFlagsNoDirectResize TableColumnFlags = TableColumnFlags(cimgui.TableColumnFlags_NoDirectResize_)
)

// SliderFlags represents cimgui.SliderFlags
// TODO: Hard-reffer to these constants.
type SliderFlags cimgui.SliderFlags

// slider flags.
const (
	SliderFlagsNone SliderFlags = cimgui.SliderFlags_None
	// Clamp value to min/max bounds when input manually with CTRL+Click. By default CTRL+Click allows going out of bounds.
	SliderFlagsAlwaysClamp SliderFlags = cimgui.SliderFlags_AlwaysClamp
	// Make the widget logarithmic (linear otherwise). Consider using ImGuiSliderFlagsNoRoundToFormat SliderFlags = cimgui.SliderFlags_NoRoundToFormat
	// a format-string with small amount of digits.
	SliderFlagsLogarithmic SliderFlags = cimgui.SliderFlags_Logarithmic
	// Disable rounding underlying value to match precision of the display format string (e.g. %.3f values are rounded to those 3 digits).
	SliderFlagsNoRoundToFormat SliderFlags = cimgui.SliderFlags_NoRoundToFormat
	// Disable CTRL+Click or Enter key allowing to input text directly into the widget.
	SliderFlagsNoInput SliderFlags = cimgui.SliderFlags_NoInput
	// [Internal] We treat using those bits as being potentially a 'float power' argument from the previous API that has got miscast
	// to this enum, and will trigger an assert if needed.
	SliderFlagsInvalidMask SliderFlags = cimgui.SliderFlags_InvalidMask
)

// PlotFlags represents cimgui.PlotFlags.
type PlotFlags cimgui.PlotFlags

// plot flags.
const (
	PlotFlagsNone        = PlotFlags(cimgui.PlotFlags_None)
	PlotFlagsNoTitle     = PlotFlags(cimgui.PlotFlags_NoTitle)
	PlotFlagsNoLegend    = PlotFlags(cimgui.PlotFlags_NoLegend)
	PlotFlagsNoMenus     = PlotFlags(cimgui.PlotFlags_NoMenus)
	PlotFlagsNoBoxSelect = PlotFlags(cimgui.PlotFlags_NoBoxSelect)
	// 	PlotFlagsNoMousePos  = PlotFlags(cimgui.PlotFlags_NoMousePos)
	// 	PlotFlagsNoHighlight = PlotFlags(cimgui.PlotFlags_NoHighlight)
	PlotFlagsNoChild = PlotFlags(cimgui.PlotFlags_NoChild)
	PlotFlagsEqual   = PlotFlags(cimgui.PlotFlags_Equal)
	// 	PlotFlagsYAxis2      = PlotFlags(cimgui.PlotFlags_YAxis2)
	// 	PlotFlagsYAxis3      = PlotFlags(cimgui.PlotFlags_YAxis3)
	// 	PlotFlagsQuery       = PlotFlags(cimgui.PlotFlags_Query)
	PlotFlagsCrosshairs = PlotFlags(cimgui.PlotFlags_Crosshairs)
	// 	PlotFlagsAntiAliased = PlotFlags(cimgui.PlotFlags_AntiAliased)
	PlotFlagsCanvasOnly = PlotFlags(cimgui.PlotFlags_CanvasOnly)
)

// PlotAxisFlags represents cimgui.PlotAxisFlags.
type PlotAxisFlags cimgui.PlotAxisFlags

// plot axis flags.
const (
	PlotAxisFlagsNone         PlotAxisFlags = PlotAxisFlags(cimgui.PlotAxisFlags_None)
	PlotAxisFlagsNoLabel      PlotAxisFlags = PlotAxisFlags(cimgui.PlotAxisFlags_NoLabel)
	PlotAxisFlagsNoGridLines  PlotAxisFlags = PlotAxisFlags(cimgui.PlotAxisFlags_NoGridLines)
	PlotAxisFlagsNoTickMarks  PlotAxisFlags = PlotAxisFlags(cimgui.PlotAxisFlags_NoTickMarks)
	PlotAxisFlagsNoTickLabels PlotAxisFlags = PlotAxisFlags(cimgui.PlotAxisFlags_NoTickLabels)
	PlotAxisFlagsForeground   PlotAxisFlags = PlotAxisFlags(cimgui.PlotAxisFlags_Foreground)
	// 	PlotAxisFlagsLogScale      PlotAxisFlags = PlotAxisFlags(cimgui.PlotAxisFlags_LogScale)
	// 	PlotAxisFlagsTime          PlotAxisFlags = PlotAxisFlags(cimgui.PlotAxisFlags_Time)
	PlotAxisFlagsInvert        PlotAxisFlags = PlotAxisFlags(cimgui.PlotAxisFlags_Invert)
	PlotAxisFlagsNoInitialFit  PlotAxisFlags = PlotAxisFlags(cimgui.PlotAxisFlags_NoInitialFit)
	PlotAxisFlagsAutoFit       PlotAxisFlags = PlotAxisFlags(cimgui.PlotAxisFlags_AutoFit)
	PlotAxisFlagsRangeFit      PlotAxisFlags = PlotAxisFlags(cimgui.PlotAxisFlags_RangeFit)
	PlotAxisFlagsLockMin       PlotAxisFlags = PlotAxisFlags(cimgui.PlotAxisFlags_LockMin)
	PlotAxisFlagsLockMax       PlotAxisFlags = PlotAxisFlags(cimgui.PlotAxisFlags_LockMax)
	PlotAxisFlagsLock          PlotAxisFlags = PlotAxisFlags(cimgui.PlotAxisFlags_Lock)
	PlotAxisFlagsNoDecorations PlotAxisFlags = PlotAxisFlags(cimgui.PlotAxisFlags_NoDecorations)
)
