package giu

import (
	"bytes"
	ctx "context"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"time"

	"github.com/AllenDang/imgui-go"
	"github.com/go-resty/resty/v2"
)

type ImageWidget interface {
	Widget

	Uv(uv0, uv1 image.Point) ImageWidget
	TintColor(tintColor color.RGBA) ImageWidget
	BorderCol(borderColor color.RGBA) ImageWidget
	OnClick(cb func()) ImageWidget
	Size(width, height float32) ImageWidget
}

func Image(data interface{}) ImageWidget {
	switch imageData := data.(type) {
	case *Texture:
		return ImageWithTexture(imageData)
	default:
		return ImageWithTexture(nil)
	}
}

type ImageState struct {
	loading bool
	failure bool
	cancel  ctx.CancelFunc
	texture *Texture
}

func (is *ImageState) Dispose() {
	is.texture = nil
	// Cancel ongoing image downloaidng
	if is.loading && is.cancel != nil {
		is.cancel()
	}
}

var _ ImageWidget = &ImageWithTextureWidget{}

type ImageWithTextureWidget struct {
	texture                *Texture
	width                  float32
	height                 float32
	uv0, uv1               image.Point
	tintColor, borderColor color.RGBA
	onClick                func()
}

func ImageWithTexture(texture *Texture) ImageWidget {
	return &ImageWithTextureWidget{
		texture:     texture,
		width:       100 * Context.platform.GetContentScale(),
		height:      100 * Context.platform.GetContentScale(),
		uv0:         image.Point{X: 0, Y: 0},
		uv1:         image.Point{X: 1, Y: 1},
		tintColor:   color.RGBA{255, 255, 255, 255},
		borderColor: color.RGBA{0, 0, 0, 0},
	}
}

func (i *ImageWithTextureWidget) Uv(uv0, uv1 image.Point) ImageWidget {
	i.uv0, i.uv1 = uv0, uv1
	return i
}

func (i *ImageWithTextureWidget) TintColor(tintColor color.RGBA) ImageWidget {
	i.tintColor = tintColor
	return i
}

func (i *ImageWithTextureWidget) BorderCol(borderColor color.RGBA) ImageWidget {
	i.borderColor = borderColor
	return i
}

func (i *ImageWithTextureWidget) OnClick(cb func()) ImageWidget {
	i.onClick = cb
	return i
}

func (i *ImageWithTextureWidget) Size(width, height float32) ImageWidget {
	scale := Context.platform.GetContentScale()
	i.width, i.height = width*scale, height*scale
	return i
}

func (i *ImageWithTextureWidget) Build() {
	size := imgui.Vec2{X: i.width, Y: i.height}
	rect := imgui.ContentRegionAvail()
	if size.X == (-1 * Context.GetPlatform().GetContentScale()) {
		size.X = rect.X
	}
	if size.Y == (-1 * Context.GetPlatform().GetContentScale()) {
		size.Y = rect.Y
	}
	if i.texture != nil && i.texture.id != 0 {
		// trick: detect click event
		if i.onClick != nil && IsMouseClicked(MouseButtonLeft) {
			cursorPos := GetCursorScreenPos()
			mousePos := GetMousePos()
			mousePos.Add(cursorPos)
			if cursorPos.X <= mousePos.X && cursorPos.Y <= mousePos.Y &&
				cursorPos.X+int(i.width) >= mousePos.X && cursorPos.Y+int(i.height) >= mousePos.Y {
				i.onClick()
			}
		}

		imgui.ImageV(i.texture.id, size, ToVec2(i.uv0), ToVec2(i.uv1), ToVec4Color(i.tintColor), ToVec4Color(i.borderColor))
	} else {
		Dummy(i.width, i.height).Build()
	}
}

type ImageWithRgbaWidget struct {
	id      string
	width   float32
	height  float32
	rgba    *image.RGBA
	onClick func()
}

func ImageWithRgba(rgba *image.RGBA) *ImageWithRgbaWidget {
	return &ImageWithRgbaWidget{
		id:     GenAutoID("ImageWithRgba_%v"),
		width:  100,
		height: 100,
		rgba:   rgba,
	}
}

func (i *ImageWithRgbaWidget) Size(width, height float32) *ImageWithRgbaWidget {
	i.width, i.height = width, height
	return i
}

func (i *ImageWithRgbaWidget) OnClick(cb func()) *ImageWithRgbaWidget {
	i.onClick = cb
	return i
}

func (i *ImageWithRgbaWidget) Build() {
	widget := ImageWithTexture(nil).Size(i.width, i.height).OnClick(i.onClick)

	if i.rgba != nil {
		state := Context.GetState(i.id)

		if state == nil {
			Context.SetState(i.id, &ImageState{})

			go func() {
				texture, err := NewTextureFromRgba(i.rgba)
				if err == nil {
					Context.SetState(i.id, &ImageState{texture: texture})
				}
			}()
		} else {
			imgState := state.(*ImageState)
			widget = ImageWithTexture(imgState.texture).Size(i.width, i.height).OnClick(i.onClick)
		}
	}

	widget.Build()
}

type ImageWithFileWidget struct {
	id      string
	width   float32
	height  float32
	imgPath string
	onClick func()
}

func ImageWithFile(imgPath string) *ImageWithFileWidget {
	return &ImageWithFileWidget{
		id:      fmt.Sprintf("ImageWithFile_%s", imgPath),
		width:   100,
		height:  100,
		imgPath: imgPath,
	}
}

func (i *ImageWithFileWidget) Size(width, height float32) *ImageWithFileWidget {
	i.width, i.height = width, height
	return i
}

func (i *ImageWithFileWidget) OnClick(cb func()) *ImageWithFileWidget {
	i.onClick = cb
	return i
}

func (i *ImageWithFileWidget) Build() {
	state := Context.GetState(i.id)

	widget := Image(nil).OnClick(i.onClick).Size(i.width, i.height)

	if state == nil {
		// Prevent multiple invocation to LoadImage.
		Context.SetState(i.id, &ImageState{})

		img, err := LoadImage(i.imgPath)
		if err == nil {
			go func() {
				texture, err := NewTextureFromRgba(img)
				if err == nil {
					Context.SetState(i.id, &ImageState{texture: texture})
				}
			}()
		}
	} else {
		imgState := state.(*ImageState)
		widget = Image(imgState.texture).OnClick(i.onClick).Size(i.width, i.height)
	}

	widget.Build()
}

type ImageWithUrlWidget struct {
	id              string
	imgUrl          string
	downloadTimeout time.Duration
	width           float32
	height          float32
	whenLoading     Layout
	whenFailure     Layout
	onReady         func()
	onFailure       func(error)
	onClick         func()
}

func ImageWithUrl(url string) *ImageWithUrlWidget {
	return &ImageWithUrlWidget{
		id:              fmt.Sprintf("ImageWithUrl_%s", url),
		imgUrl:          url,
		downloadTimeout: 10 * time.Second,
		width:           100,
		height:          100,
		whenLoading:     Layout{Dummy(100, 100)},
		whenFailure:     Layout{Dummy(100, 100)},
	}
}

// Event trigger when image is downloaded and ready to display.
func (i *ImageWithUrlWidget) OnReady(onReady func()) *ImageWithUrlWidget {
	i.onReady = onReady
	return i
}

func (i *ImageWithUrlWidget) OnFailure(onFailure func(error)) *ImageWithUrlWidget {
	i.onFailure = onFailure
	return i
}

func (i *ImageWithUrlWidget) OnClick(cb func()) *ImageWithUrlWidget {
	i.onClick = cb
	return i
}

func (i *ImageWithUrlWidget) Timeout(downloadTimeout time.Duration) *ImageWithUrlWidget {
	i.downloadTimeout = downloadTimeout
	return i
}

func (i *ImageWithUrlWidget) Size(width, height float32) *ImageWithUrlWidget {
	i.width, i.height = width, height
	return i
}

func (i *ImageWithUrlWidget) LayoutForLoading(widgets ...Widget) *ImageWithUrlWidget {
	i.whenLoading = Layout(widgets)
	return i
}

func (i *ImageWithUrlWidget) LayoutForFailure(widgets ...Widget) *ImageWithUrlWidget {
	i.whenFailure = Layout(widgets)
	return i
}

func (i *ImageWithUrlWidget) Build() {
	state := Context.GetState(i.id)

	widget := Image(nil).OnClick(i.onClick).Size(i.width, i.height)

	if state == nil {
		// Prevent multiple invocation to download image.
		downloadContext, cancalFunc := ctx.WithCancel(ctx.Background())
		Context.SetState(i.id, &ImageState{loading: true, cancel: cancalFunc})

		go func() {
			// Load image from url
			client := resty.New()
			client.SetTimeout(i.downloadTimeout)
			resp, err := client.R().SetContext(downloadContext).Get(i.imgUrl)
			if err != nil {
				Context.SetState(i.id, &ImageState{failure: true})

				// Trigger onFailure event
				if i.onFailure != nil {
					i.onFailure(err)
				}
				return
			}

			img, _, err := image.Decode(bytes.NewReader(resp.Body()))
			if err != nil {
				Context.SetState(i.id, &ImageState{failure: true})

				// Trigger onFailure event
				if i.onFailure != nil {
					i.onFailure(err)
				}
				return
			}

			rgba := image.NewRGBA(img.Bounds())
			draw.Draw(rgba, img.Bounds(), img, image.Point{}, draw.Src)

			texture, err := NewTextureFromRgba(rgba)
			if err != nil {
				Context.SetState(i.id, &ImageState{failure: true})

				// Trigger onFailure event
				if i.onFailure != nil {
					i.onFailure(err)
				}
				return
			}
			Context.SetState(i.id, &ImageState{loading: false, texture: texture})

			// Trigger onReady event
			if i.onReady != nil {
				i.onReady()
			}
		}()
	} else {
		imgState := state.(*ImageState)
		if imgState.failure {
			i.whenFailure.Build()
			return
		}

		if imgState.loading {
			i.whenLoading.Build()
			return
		}

		widget = Image(imgState.texture).OnClick(i.onClick).Size(i.width, i.height)
	}

	widget.Build()
}
