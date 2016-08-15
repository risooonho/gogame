package gogame

import "github.com/veandco/go-sdl2/sdl"

// Color is a RGBA representation of a color.
// All of the components should be between 0 and 1 (inclusive).
// If they are not, nevermind, I can deal with it.
type Color struct {
	R, G, B, A float64
}

// Colors defines some common colors.
var Colors = map[string]Color{
	"black":   {0.0, 0.0, 0.0, 1.0},
	"grey":    {0.5, 0.5, 0.5, 1.0},
	"white":   {1.0, 1.0, 1.0, 1.0},
	"red":     {1.0, 0.0, 0.0, 1.0},
	"green":   {0.0, 1.0, 0.0, 1.0},
	"blue":    {0.0, 0.0, 1.0, 1.0},
	"cyan":    {0.0, 1.0, 1.0, 1.0},
	"magenta": {1.0, 0.0, 1.0, 1.0},
	"yellow":  {1.0, 1.0, 0.0, 1.0},
}

func (c *Color) toSDLRGBA() (r, g, b, a byte) {
	r = byte(255 * clamp(c.R, 0, 1))
	g = byte(255 * clamp(c.G, 0, 1))
	b = byte(255 * clamp(c.B, 0, 1))
	a = byte(255 * clamp(c.A, 0, 1))
	return
}

func (c *Color) toSDL() sdl.Color {
	return sdl.Color{
		R: byte(255 * clamp(c.R, 0, 1)),
		G: byte(255 * clamp(c.G, 0, 1)),
		B: byte(255 * clamp(c.B, 0, 1)),
		A: byte(255 * clamp(c.A, 0, 1)),
	}
}

func clamp(x, low, high float64) float64 {
	if x < low {
		return low
	}
	if x > high {
		return high
	}
	return x
}
