package giu

import (
	ctx "context"
	"fmt"
	"image"
	"image/color"
	"net/http"
	"reflect"
	"time"

	"github.com/AllenDang/cimgui-go/imgui"
)

var _ Widget = &ImageWidget{}

// ImageWidget adds an image.
// The default size is the size of the image,
// to set a specific size, use .Size(width, height).
// NOTE: ImageWidget is going to be deprecated. ImageWithRGBAWidget
// should be used instead, however, because it is a native
// imgui's solution it is still there.
type ImageWidget struct {
	texture                *Texture
	width                  float32
	height                 float32
	scale                  imgui.Vec2
	uv0, uv1               imgui.Vec2
	tintColor, borderColor color.Color
	onClick                func()
}

// Image adds an image from giu.Texture.
func Image(texture *Texture) *ImageWidget {
	return &ImageWidget{
		texture:     texture,
		width:       0,
		height:      0,
		scale:       imgui.Vec2{X: 1, Y: 1},
		uv0:         imgui.Vec2{X: 0, Y: 0},
		uv1:         imgui.Vec2{X: 1, Y: 1},
		tintColor:   color.RGBA{255, 255, 255, 255},
		borderColor: color.RGBA{0, 0, 0, 0},
	}
}

// Uv allows to specify uv parameters.
func (i *ImageWidget) Uv(uv0X, uv0Y, uv1X, uv1Y float32) *ImageWidget {
	i.uv0.X, i.uv0.Y, i.uv1.X, i.uv1.Y = uv0X, uv0Y, uv1X, uv1Y
	return i
}

// TintColor sets image's tint color.
func (i *ImageWidget) TintColor(tintColor color.Color) *ImageWidget {
	i.tintColor = tintColor
	return i
}

// BorderCol sets color of the border.
func (i *ImageWidget) BorderCol(borderColor color.Color) *ImageWidget {
	i.borderColor = borderColor
	return i
}

// OnClick adds on-click-callback.
func (i *ImageWidget) OnClick(cb func()) *ImageWidget {
	i.onClick = cb
	return i
}

// Size sets image size.
func (i *ImageWidget) Size(width, height float32) *ImageWidget {
	// Size image with DPI scaling
	i.width, i.height = width, height

	return i
}

// Scale multiply dimensions after size.
func (i *ImageWidget) Scale(scaleX, scaleY float32) *ImageWidget {
	// Size image with DPI scaling
	i.scale = imgui.Vec2{X: scaleX, Y: scaleY}

	return i
}

// Build implements Widget interface.
func (i *ImageWidget) Build() {
	if i.width == 0 && i.height == 0 {
		if i.texture != nil {
			i.width, i.height = float32(i.texture.tex.Width), float32(i.texture.tex.Height)
		} else {
			i.width, i.height = 100, 100
		}
	}

	size := imgui.Vec2{X: i.width, Y: i.height}

	if size.X == -1 {
		rect := imgui.ContentRegionAvail()
		size.X = rect.X
	}

	if size.Y == -1 {
		rect := imgui.ContentRegionAvail()
		size.Y = rect.Y
	}

	size.X *= i.scale.X
	size.Y *= i.scale.Y

	if i.texture == nil || i.texture.tex == nil {
		Dummy(size.X, size.Y).Build()
		return
	}

	// trick: detect click event
	if i.onClick != nil && IsMouseClicked(MouseButtonLeft) && IsWindowHovered(0) {
		cursorPos := GetCursorScreenPos()
		mousePos := GetMousePos()

		if cursorPos.X <= mousePos.X && cursorPos.Y <= mousePos.Y &&
			cursorPos.X+int(size.X) >= mousePos.X && cursorPos.Y+int(size.Y) >= mousePos.Y {
			i.onClick()
		}
	}

	imgui.ImageWithBgV(i.texture.tex.ID, size, i.uv0, i.uv1, ToVec4Color(i.borderColor), ToVec4Color(i.tintColor))
}

type imageState struct {
	loading bool
	failure bool
	cancel  ctx.CancelFunc
	texture *Texture
	img     *image.RGBA
}

// Dispose cleans imageState (implements Disposable interface).
func (is *imageState) Dispose() {
	is.texture = nil
	// Cancel ongoing image downloading
	if is.loading && is.cancel != nil {
		is.cancel()
	}
}

var _ Widget = &ImageWithRgbaWidget{}

// ImageWithRgbaWidget wraps ImageWidget.
// It is more useful because it doesn't make you to care about
// imgui textures. You can just pass golang-native image.Image and
// display it in giu.
type ImageWithRgbaWidget struct {
	id   ID
	rgba *image.RGBA
	img  *ImageWidget
}

// ImageWithRgba creates ImageWithRgbaWidget.
// The default size is the size of the image,
// to set a specific size, use .Size(width, height).
func ImageWithRgba(rgba image.Image) *ImageWithRgbaWidget {
	return &ImageWithRgbaWidget{
		id:   GenAutoID("ImageWithRgba"),
		rgba: ImageToRgba(rgba),
		img:  Image(nil),
	}
}

// ID sets the interval id of ImageWithRgba widgets.
func (i *ImageWithRgbaWidget) ID(id ID) *ImageWithRgbaWidget {
	i.id = id
	return i
}

// Size sets image's size.
func (i *ImageWithRgbaWidget) Size(width, height float32) *ImageWithRgbaWidget {
	i.img.Size(width, height)
	return i
}

// OnClick sets click callback.
func (i *ImageWithRgbaWidget) OnClick(cb func()) *ImageWithRgbaWidget {
	i.img.OnClick(cb)
	return i
}

// Build implements Widget interface.
func (i *ImageWithRgbaWidget) Build() {
	if i.rgba != nil {
		var imgState *imageState
		if imgState = GetState[imageState](Context, i.id); imgState == nil || !reflect.DeepEqual(i.rgba, imgState.img) {
			imgState = &imageState{}
			imgState.img = i.rgba
			SetState(Context, i.id, imgState)

			NewTextureFromRgba(i.rgba, func(tex *Texture) {
				imgState.texture = tex
			})
		}

		i.img.texture = imgState.texture
	}

	i.img.Build()
}

var _ Widget = &ImageWithFileWidget{}

// ImageWithFileWidget allows to display an image directly
// from .png file.
// NOTE: Be aware that project using this solution may not be portable
// because files are not included in executable binaries!
// You may want to use "embed" package and ImageWithRgba instead.
type ImageWithFileWidget struct {
	id      ID
	imgPath string
	img     *ImageWidget
}

// ImageWithFile constructs a new ImageWithFileWidget.
// The default size is the size of the image,
// to set a specific size, use .Size(width, height).
func ImageWithFile(imgPath string) *ImageWithFileWidget {
	return &ImageWithFileWidget{
		id:      ID(fmt.Sprintf("ImageWithFile_%s", imgPath)),
		imgPath: imgPath,
		img:     Image(nil),
	}
}

// ID sets the interval id of ImageWithFile widgets.
func (i *ImageWithFileWidget) ID(id ID) *ImageWithFileWidget {
	i.id = id
	return i
}

// Size sets image's size.
func (i *ImageWithFileWidget) Size(width, height float32) *ImageWithFileWidget {
	i.img.Size(width, height)
	return i
}

// OnClick sets click callback.
func (i *ImageWithFileWidget) OnClick(cb func()) *ImageWithFileWidget {
	i.img.OnClick(cb)
	return i
}

// Build implements Widget interface.
func (i *ImageWithFileWidget) Build() {
	var imgState *imageState
	if imgState = GetState[imageState](Context, i.id); imgState == nil {
		// Prevent multiple invocation to LoadImage.
		imgState = &imageState{}
		SetState(Context, i.id, imgState)

		img, err := LoadImage(i.imgPath)
		if err == nil {
			NewTextureFromRgba(img, func(tex *Texture) {
				imgState.texture = tex
			})
		}
	}

	i.img.texture = imgState.texture
	i.img.Build()
}

var _ Widget = &ImageWithURLWidget{}

// ImageWithURLWidget allows to display an image using
// an URL as image source.
type ImageWithURLWidget struct {
	id              ID
	imgURL          string
	downloadTimeout time.Duration
	whenLoading     Layout
	whenFailure     Layout
	onReady         func()
	onFailure       func(error)
	img             *ImageWidget
}

// ImageWithURL creates ImageWithURLWidget.
// The default size is the size of the image,
// to set a specific size, use .Size(width, height).
func ImageWithURL(url string) *ImageWithURLWidget {
	return &ImageWithURLWidget{
		id:              ID(fmt.Sprintf("ImageWithURL_%s", url)),
		imgURL:          url,
		downloadTimeout: 10 * time.Second,
		whenLoading:     Layout{Dummy(100, 100)},
		whenFailure:     Layout{Dummy(100, 100)},
		img:             Image(nil),
	}
}

// OnReady sets event trigger when image is downloaded and ready to display.
func (i *ImageWithURLWidget) OnReady(onReady func()) *ImageWithURLWidget {
	i.onReady = onReady
	return i
}

// OnFailure sets event trigger when image failed to download/load.
func (i *ImageWithURLWidget) OnFailure(onFailure func(error)) *ImageWithURLWidget {
	i.onFailure = onFailure
	return i
}

// OnClick sets click callback.
func (i *ImageWithURLWidget) OnClick(cb func()) *ImageWithURLWidget {
	i.img.OnClick(cb)
	return i
}

// Timeout sets download timeout.
func (i *ImageWithURLWidget) Timeout(downloadTimeout time.Duration) *ImageWithURLWidget {
	i.downloadTimeout = downloadTimeout
	return i
}

// Size sets image's size.
func (i *ImageWithURLWidget) Size(width, height float32) *ImageWithURLWidget {
	i.img.Size(width, height)
	return i
}

// LayoutForLoading allows to set layout rendered while loading an image.
func (i *ImageWithURLWidget) LayoutForLoading(widgets ...Widget) *ImageWithURLWidget {
	i.whenLoading = widgets
	return i
}

// LayoutForFailure allows to specify layout when image failed to download.
func (i *ImageWithURLWidget) LayoutForFailure(widgets ...Widget) *ImageWithURLWidget {
	i.whenFailure = widgets
	return i
}

// Build implements Widget interface.
func (i *ImageWithURLWidget) Build() {
	var imgState *imageState
	if imgState = GetState[imageState](Context, i.id); imgState == nil {
		imgState = &imageState{}
		SetState(Context, i.id, imgState)

		// Prevent multiple invocation to download image.
		downloadContext, cancelFunc := ctx.WithCancel(ctx.Background())

		SetState(Context, i.id, &imageState{loading: true, cancel: cancelFunc})

		errorFn := func(err error) {
			SetState(Context, i.id, &imageState{failure: true})

			// Trigger onFailure event
			if i.onFailure != nil {
				i.onFailure(err)
			}
		}

		go func() {
			// Load image from url
			client := &http.Client{Timeout: i.downloadTimeout}

			req, err := http.NewRequestWithContext(downloadContext, http.MethodGet, i.imgURL, http.NoBody)
			if err != nil {
				errorFn(err)
				return
			}

			resp, err := client.Do(req)
			if err != nil {
				errorFn(err)
				return
			}

			defer func() {
				if closeErr := resp.Body.Close(); closeErr != nil {
					errorFn(closeErr)
				}
			}()

			img, _, err := image.Decode(resp.Body)
			if err != nil {
				errorFn(err)
				return
			}

			rgba := ImageToRgba(img)

			EnqueueNewTextureFromRgba(rgba, func(tex *Texture) {
				SetState(Context, i.id, &imageState{
					loading: false,
					failure: false,
					texture: tex,
				})
			})

			// Trigger onReady event
			if i.onReady != nil {
				i.onReady()
			}
		}()
	}

	imgState = GetState[imageState](Context, i.id)

	switch {
	case imgState.failure:
		i.whenFailure.Build()
	case imgState.loading:
		i.whenLoading.Build()
	default:
		i.img.texture = imgState.texture
		i.img.Build()
	}
}
