package gogame

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
)

// LoadPicture loads a picture from a file stored at the specified path.
// If the loading fails, an error is returned.
func LoadPicture(path string) (*Picture, error) {
	var (
		pic Picture
		err error
	)
	pic.surface, err = img.Load(path)
	if err != nil {
		return nil, fmt.Errorf("failed to load picture: %s", path)
	}
	pic.surface.Flags |= staticSurface
	pic.rect = sdl.Rect{X: 0, Y: 0, W: pic.surface.W, H: pic.surface.H}
	return &pic, nil
}

// Picture is a static raster image, usually loaded from a file.
type Picture struct {
	surface *sdl.Surface
	rect    sdl.Rect
}

// Size returns the width and height of a picture in pixels.
func (p *Picture) Size() (w, h int) {
	return int(p.surface.W), int(p.surface.H)
}

// Slice cuts a rectangle (x, y, w, h) from a picture.
func (p *Picture) Slice(x, y, w, h int) *Picture {
	return &Picture{
		surface: p.surface,
		rect: sdl.Rect{
			X: p.rect.X + int32(x),
			Y: p.rect.Y + int32(y),
			W: int32(w),
			H: int32(h),
		},
	}
}

const (
	staticSurface = 1 << iota
)
