package giu

import (
	"image"
	"image/color"
	"runtime"
	"time"

	"github.com/AllenDang/cimgui-go"
	"github.com/faiface/mainthread"
	"github.com/go-gl/glfw/v3.3/glfw"
	"gopkg.in/eapache/queue.v1"
)

// MasterWindowFlags wraps cimgui.GLFWWindowFlags.
type MasterWindowFlags cimgui.GLFWWindowFlags

// master window flags.
const (
	// Specifies the window will be fixed size.
	MasterWindowFlagsNotResizable MasterWindowFlags = MasterWindowFlags(cimgui.GLFWWindowFlagsNotResizable)
	// Specifies whether the window is maximized.
	MasterWindowFlagsMaximized MasterWindowFlags = MasterWindowFlags(cimgui.GLFWWindowFlagsMaximized)
	// Specifies whether the window will be always-on-top.
	MasterWindowFlagsFloating MasterWindowFlags = MasterWindowFlags(cimgui.GLFWWindowFlagsFloating)
	// Specifies whether the window will be frameless.
	MasterWindowFlagsFrameless MasterWindowFlags = MasterWindowFlags(cimgui.GLFWWindowFlagsFrameless)
	// Specifies whether the window will be transparent.
	MasterWindowFlagsTransparent MasterWindowFlags = MasterWindowFlags(cimgui.GLFWWindowFlagsTransparent)
)

// DontCare could be used as an argument to (*MasterWindow).SetSizeLimits.
// var DontCare int = cimgui.GlfwDontCare

// MasterWindow represents a glfw master window
// It is a base for a windows (see Window.go).
type MasterWindow struct {
	width      int
	height     int
	clearColor [4]float32
	title      string
	platform   cimgui.Platform
	renderer   cimgui.Renderer
	context    *cimgui.ImGuiContext
	io         *cimgui.ImGuiIO
	updateFunc func()

	// possibility to expend InputHandler's stuff
	// See SetAdditionalInputHandler
	additionalInputCallback InputHandlerHandleCallback
}

// NewMasterWindow creates a new master window and initializes GLFW.
// it should be called in main function. For more details and use cases,
// see examples/helloworld/.
func NewMasterWindow(title string, width, height int, flags MasterWindowFlags) *MasterWindow {
	context := cimgui.CreateContext(nil)
	cimgui.ImPlotCreateContext()
	cimgui.ImNodesCreateContext()

	io := cimgui.GetIO()

	io.SetConfigFlags(cimgui.ConfigFlagEnablePowerSavingMode | cimgui.BackendFlagsRendererHasVtxOffset)

	// Disable cimgui.ini
	io.SetIniFilename("")

	p, err := cimgui.NewGLFW(io, title, width, height, cimgui.GLFWWindowFlags(flags))
	if err != nil {
		panic(err)
	}

	r, err := cimgui.NewOpenGL3(io, 1.0)
	if err != nil {
		panic(err)
	}

	Context = CreateContext(p, r)

	// init texture loading queue
	Context.textureLoadingQueue = queue.New()

	mw := &MasterWindow{
		clearColor: [4]float32{0, 0, 0, 1},
		width:      width,
		height:     height,
		title:      title,
		io:         &io,
		context:    context,
		platform:   p,
		renderer:   r,
	}

	mw.SetInputHandler(newInputHandler())

	p.SetSizeChangeCallback(mw.sizeChange)

	mw.setTheme()

	return mw
}

func (w *MasterWindow) setTheme() {
	style := cimgui.GetStyle()

	// Scale DPI in windows
	if runtime.GOOS == "windows" {
		style.ScaleAllSizes(Context.GetPlatform().GetContentScale())
	}

	cimgui.PushStyleVar_Float(cimgui.StyleVar_WindowRounding, 2)
	cimgui.PushStyleVar_Float(cimgui.StyleVar_FrameRounding, 4)
	cimgui.PushStyleVar_Float(cimgui.StyleVar_GrabRounding, 4)
	cimgui.PushStyleVar_Float(cimgui.StyleVar_FrameBorderSize, 1)

	style.SetColor(cimgui.Col_Text, cimgui.ImVec4{X: 0.95, Y: 0.96, Z: 0.98, W: 1.00})
	style.SetColor(cimgui.Col_TextDisabled, cimgui.ImVec4{X: 0.36, Y: 0.42, Z: 0.47, W: 1.00})
	style.SetColor(cimgui.Col_WindowBg, cimgui.ImVec4{X: 0.11, Y: 0.15, Z: 0.17, W: 1.00})
	style.SetColor(cimgui.Col_ChildBg, cimgui.ImVec4{X: 0.15, Y: 0.18, Z: 0.22, W: 1.00})
	style.SetColor(cimgui.Col_PopupBg, cimgui.ImVec4{X: 0.08, Y: 0.08, Z: 0.08, W: 0.94})
	style.SetColor(cimgui.Col_Border, cimgui.ImVec4{X: 0.08, Y: 0.10, Z: 0.12, W: 1.00})
	style.SetColor(cimgui.Col_BorderShadow, cimgui.ImVec4{X: 0.00, Y: 0.00, Z: 0.00, W: 0.00})
	style.SetColor(cimgui.Col_FrameBg, cimgui.ImVec4{X: 0.20, Y: 0.25, Z: 0.29, W: 1.00})
	style.SetColor(cimgui.Col_FrameBgHovered, cimgui.ImVec4{X: 0.12, Y: 0.20, Z: 0.28, W: 1.00})
	style.SetColor(cimgui.Col_FrameBgActive, cimgui.ImVec4{X: 0.09, Y: 0.12, Z: 0.14, W: 1.00})
	style.SetColor(cimgui.Col_TitleBg, cimgui.ImVec4{X: 0.09, Y: 0.12, Z: 0.14, W: 0.65})
	style.SetColor(cimgui.Col_TitleBgActive, cimgui.ImVec4{X: 0.08, Y: 0.10, Z: 0.12, W: 1.00})
	style.SetColor(cimgui.Col_TitleBgCollapsed, cimgui.ImVec4{X: 0.00, Y: 0.00, Z: 0.00, W: 0.51})
	style.SetColor(cimgui.Col_MenuBarBg, cimgui.ImVec4{X: 0.15, Y: 0.18, Z: 0.22, W: 1.00})
	style.SetColor(cimgui.Col_ScrollbarBg, cimgui.ImVec4{X: 0.02, Y: 0.02, Z: 0.02, W: 0.39})
	style.SetColor(cimgui.Col_ScrollbarGrab, cimgui.ImVec4{X: 0.20, Y: 0.25, Z: 0.29, W: 1.00})
	style.SetColor(cimgui.Col_ScrollbarGrabHovered, cimgui.ImVec4{X: 0.18, Y: 0.22, Z: 0.25, W: 1.00})
	style.SetColor(cimgui.Col_ScrollbarGrabActive, cimgui.ImVec4{X: 0.09, Y: 0.21, Z: 0.31, W: 1.00})
	style.SetColor(cimgui.Col_CheckMark, cimgui.ImVec4{X: 0.28, Y: 0.56, Z: 1.00, W: 1.00})
	style.SetColor(cimgui.Col_SliderGrab, cimgui.ImVec4{X: 0.28, Y: 0.56, Z: 1.00, W: 1.00})
	style.SetColor(cimgui.Col_SliderGrabActive, cimgui.ImVec4{X: 0.37, Y: 0.61, Z: 1.00, W: 1.00})
	style.SetColor(cimgui.Col_Button, cimgui.ImVec4{X: 0.20, Y: 0.25, Z: 0.29, W: 1.00})
	style.SetColor(cimgui.Col_ButtonHovered, cimgui.ImVec4{X: 0.28, Y: 0.56, Z: 1.00, W: 1.00})
	style.SetColor(cimgui.Col_ButtonActive, cimgui.ImVec4{X: 0.06, Y: 0.53, Z: 0.98, W: 1.00})
	style.SetColor(cimgui.Col_Header, cimgui.ImVec4{X: 0.20, Y: 0.25, Z: 0.29, W: 0.55})
	style.SetColor(cimgui.Col_HeaderHovered, cimgui.ImVec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.80})
	style.SetColor(cimgui.Col_HeaderActive, cimgui.ImVec4{X: 0.26, Y: 0.59, Z: 0.98, W: 1.00})
	style.SetColor(cimgui.Col_Separator, cimgui.ImVec4{X: 0.20, Y: 0.25, Z: 0.29, W: 1.00})
	style.SetColor(cimgui.Col_SeparatorHovered, cimgui.ImVec4{X: 0.10, Y: 0.40, Z: 0.75, W: 0.78})
	style.SetColor(cimgui.Col_SeparatorActive, cimgui.ImVec4{X: 0.10, Y: 0.40, Z: 0.75, W: 1.00})
	style.SetColor(cimgui.Col_ResizeGrip, cimgui.ImVec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.25})
	style.SetColor(cimgui.Col_ResizeGripHovered, cimgui.ImVec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.67})
	style.SetColor(cimgui.Col_ResizeGripActive, cimgui.ImVec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.95})
	style.SetColor(cimgui.Col_Tab, cimgui.ImVec4{X: 0.11, Y: 0.15, Z: 0.17, W: 1.00})
	style.SetColor(cimgui.Col_TabHovered, cimgui.ImVec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.80})
	style.SetColor(cimgui.Col_TabActive, cimgui.ImVec4{X: 0.20, Y: 0.25, Z: 0.29, W: 1.00})
	style.SetColor(cimgui.Col_TabUnfocused, cimgui.ImVec4{X: 0.11, Y: 0.15, Z: 0.17, W: 1.00})
	style.SetColor(cimgui.Col_TabUnfocusedActive, cimgui.ImVec4{X: 0.11, Y: 0.15, Z: 0.17, W: 1.00})
	style.SetColor(cimgui.Col_PlotLines, cimgui.ImVec4{X: 0.61, Y: 0.61, Z: 0.61, W: 1.00})
	style.SetColor(cimgui.Col_PlotLinesHovered, cimgui.ImVec4{X: 1.00, Y: 0.43, Z: 0.35, W: 1.00})
	style.SetColor(cimgui.Col_PlotHistogram, cimgui.ImVec4{X: 0.90, Y: 0.70, Z: 0.00, W: 1.00})
	style.SetColor(cimgui.Col_PlotHistogramHovered, cimgui.ImVec4{X: 1.00, Y: 0.60, Z: 0.00, W: 1.00})
	style.SetColor(cimgui.Col_TextSelectedBg, cimgui.ImVec4{X: 0.26, Y: 0.59, Z: 0.98, W: 0.35})
	style.SetColor(cimgui.Col_DragDropTarget, cimgui.ImVec4{X: 1.00, Y: 1.00, Z: 0.00, W: 0.90})
	style.SetColor(cimgui.Col_NavHighlight, cimgui.ImVec4{X: 0.26, Y: 0.59, Z: 0.98, W: 1.00})
	style.SetColor(cimgui.Col_NavWindowingHighlight, cimgui.ImVec4{X: 1.00, Y: 1.00, Z: 1.00, W: 0.70})
	style.SetColor(cimgui.Col_TableHeaderBg, cimgui.ImVec4{X: 0.12, Y: 0.20, Z: 0.28, W: 1.00})
	style.SetColor(cimgui.Col_TableBorderStrong, cimgui.ImVec4{X: 0.20, Y: 0.25, Z: 0.29, W: 1.00})
	style.SetColor(cimgui.Col_TableBorderLight, cimgui.ImVec4{X: 0.20, Y: 0.25, Z: 0.29, W: 0.70})
}

// SetBgColor sets background color of master window.
func (w *MasterWindow) SetBgColor(bgColor color.Color) {
	const mask = 0xffff

	r, g, b, a := bgColor.RGBA()
	w.clearColor = [4]float32{
		float32(r) / mask,
		float32(g) / mask,
		float32(b) / mask,
		float32(a) / mask,
	}
}

func (w *MasterWindow) sizeChange(width, height int) {
	w.render()
}

func (w *MasterWindow) render() {
	if !w.platform.IsVisible() || w.platform.IsMinimized() {
		return
	}

	Context.invalidAllState()
	defer Context.cleanState()

	Context.FontAtlas.rebuildFontAtlas()

	p := w.platform
	r := w.renderer

	mainStylesheet := Style()
	if s, found := Context.cssStylesheet["main"]; found {
		mainStylesheet = s
	}

	p.NewFrame()
	r.PreRender(w.clearColor)

	cimgui.NewFrame()
	mainStylesheet.Push()
	w.updateFunc()
	mainStylesheet.Pop()
	cimgui.Render()

	r.Render(p.DisplaySize(), p.FramebufferSize(), cimgui.RenderedDrawData())
	p.PostRender()
}

// Run the main loop to create new frame, process events and call update ui func.
func (w *MasterWindow) run() {
	p := w.platform

	ticker := time.NewTicker(time.Second / time.Duration(p.GetTPS()))
	shouldQuit := false

	for !shouldQuit {
		mainthread.Call(func() {
			// process texture load requests
			if Context.textureLoadingQueue != nil && Context.textureLoadingQueue.Length() > 0 {
				for Context.textureLoadingQueue.Length() > 0 {
					request, ok := Context.textureLoadingQueue.Remove().(textureLoadRequest)
					Assert(ok, "MasterWindow", "Run", "processing texture requests: wrong type of texture request")
					NewTextureFromRgba(request.img, request.cb)
				}
			}

			p.ProcessEvents()
			w.render()

			shouldQuit = p.ShouldStop()
		})

		<-ticker.C
	}
}

// GetSize return size of master window.
func (w *MasterWindow) GetSize() (width, height int) {
	if w.platform != nil {
		if glfwPlatform, ok := w.platform.(*cimgui.GLFW); ok {
			return glfwPlatform.GetWindow().GetSize()
		}
	}

	return w.width, w.height
}

// GetPos return position of master window.
func (w *MasterWindow) GetPos() (x, y int) {
	if w.platform != nil {
		if glfwPlatform, ok := w.platform.(*cimgui.GLFW); ok {
			x, y = glfwPlatform.GetWindow().GetPos()
		}
	}

	return
}

// SetPos sets position of master window.
func (w *MasterWindow) SetPos(x, y int) {
	if w.platform != nil {
		if glfwPlatform, ok := w.platform.(*cimgui.GLFW); ok {
			glfwPlatform.GetWindow().SetPos(x, y)
		}
	}
}

// SetSize sets size of master window.
func (w *MasterWindow) SetSize(x, y int) {
	if w.platform != nil {
		if glfwPlatform, ok := w.platform.(*cimgui.GLFW); ok {
			mainthread.CallNonBlock(func() {
				glfwPlatform.GetWindow().SetSize(x, y)
			})
		}
	}
}

// SetCloseCallback sets the close callback of the window, which is called when
// the user attempts to close the window, for example by clicking the close
// widget in the title bar.
//
// The close flag is set before this callback is called, but you can modify it at
// any time with returned value of callback function.
//
// Mac OS X: Selecting Quit from the application menu will trigger the close
// callback for all windows.
func (w *MasterWindow) SetCloseCallback(cb func() bool) {
	w.platform.SetCloseCallback(cb)
}

// SetDropCallback sets callback when file was dropped into the window.
func (w *MasterWindow) SetDropCallback(cb func([]string)) {
	w.platform.SetDropCallback(cb)
}

// Run runs the main loop.
// loopFunc will be used to construct the ui.
// Run should be called at the end of main function, after setting
// up the master window.
func (w *MasterWindow) Run(loopFunc func()) {
	mainthread.Run(func() {
		Context.isRunning = true
		w.updateFunc = loopFunc

		Context.isAlive = true

		w.run()

		Context.isAlive = false

		mainthread.Call(func() {
			w.renderer.Dispose()
			w.platform.Dispose()

			cimgui.ImNodesDestroyContext()
			cimgui.ImPlotDestroyContext()
			w.context.Destroy()
		})

		Context.isRunning = false
	})
}

// RegisterKeyboardShortcuts registers a global - master window - keyboard shortcuts.
func (w *MasterWindow) RegisterKeyboardShortcuts(s ...WindowShortcut) *MasterWindow {
	for _, shortcut := range s {
		Context.InputHandler.RegisterKeyboardShortcuts(Shortcut{
			Key:      shortcut.Key,
			Modifier: shortcut.Modifier,
			Callback: shortcut.Callback,
			IsGlobal: GlobalShortcut,
		})
	}

	return w
}

// SetIcon sets the icon of the specified window. If passed an array of candidate images,
// those of or closest to the sizes desired by the system are selected. If no images are
// specified, the window reverts to its default icon.
//
// The image is ideally provided in the form of *image.NRGBA.
// The pixels are 32-bit, little-endian, non-premultiplied RGBA, i.e. eight
// bits per channel with the red channel first. They are arranged canonically
// as packed sequential rows, starting from the top-left corner. If the image
// type is not *image.NRGBA, it will be converted to it.
//
// The desired image sizes varies depending on platform and system settings. The selected
// images will be rescaled as needed. Good sizes include 16x16, 32x32 and 48x48.
func (w *MasterWindow) SetIcon(icons []image.Image) {
	w.platform.SetIcon(icons)
}

// SetSizeLimits sets the size limits of the client area of the specified window.
// If the window is full screen or not resizable, this function does nothing.
//
// The size limits are applied immediately and may cause the window to be resized.
// To specify only a minimum size or only a maximum one, set the other pair to giu.DontCare.
// To disable size limits for a window, set them all to giu.DontCare.
func (w *MasterWindow) SetSizeLimits(minw, minh, maxw, maxh int) {
	w.platform.SetSizeLimits(minw, minh, maxw, maxh)
}

// SetTitle updates master window's title.
func (w *MasterWindow) SetTitle(title string) {
	w.platform.SetTitle(title)
}

// Close will safely close the master window.
func (w *MasterWindow) Close() {
	w.SetShouldClose(true)
}

// SetShouldClose sets whether master window should be closed.
func (w *MasterWindow) SetShouldClose(v bool) {
	w.platform.SetShouldStop(v)
}

// SetInputHandler allows to change default input handler.
// see InputHandler.go.
func (w *MasterWindow) SetInputHandler(handler InputHandler) {
	Context.InputHandler = handler

	w.platform.SetInputCallback(func(key glfw.Key, modifier glfw.ModifierKey, action glfw.Action) {
		k, m, a := Key(key), Modifier(modifier), Action(action)
		handler.Handle(k, m, a)
		if w.additionalInputCallback != nil {
			w.additionalInputCallback(k, m, a)
		}
	})
}

// SetAdditionalInputHandlerCallback allows to set an input callback to handle more events (not only these from giu.inputHandler).
// See examples/issue-501.
func (w *MasterWindow) SetAdditionalInputHandlerCallback(cb InputHandlerHandleCallback) {
	w.additionalInputCallback = cb
}
