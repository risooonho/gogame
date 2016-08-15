package gogame

// Output combines all of the output methods together.
type Output interface {
	WindowOutput
	VideoOutput
	AudioOutput
}

// WindowOutput lets you resize a window, set it fullscreen or windowed.
type WindowOutput interface {
	// WindowSetFullscreen sets window fullscreen if fullscreen is true. Otherwise
	// sets it windowed.
	WindowSetFullscreen(fullscreen bool)

	// WindowResize changes the size of a window (or resolution, if the window is fullscreen).
	WindowResize(w, h int)
}

// VideoOutput combines all of the video output methods together and adds a Clear method
// to clear the screen.
type VideoOutput interface {
	PrimitiveVideoOutput
	PictureVideoOutput

	// Clear fill whole video output with one color.
	Clear(color Color)
}

// PrimitiveVideoOutput lets you draw primitive geometric shapes.
type PrimitiveVideoOutput interface {
	// DrawLine draws a line of specified thickness from point (x1, y1) to point (x2, y2)
	// using the provided color.
	DrawLine(x1, y1, x2, y2, thickness float64, color Color)

	// DrawPolygon draws a closed polygon with min(len(x), len(y)) points.
	// i-th point of the polygon has coordinates (x[i], y[i]).
	// If the thickness is 0, the polygon will be filled.
	DrawPolygon(x, y []float64, thickness float64, color Color)
}

// PictureVideoOutput lets you draw pictures.
type PictureVideoOutput interface {
	// DrawPicture draws a picture onto a rectangle (x, y, w, h). The picture will be
	// stretched to fit the rectangle.
	DrawPicture(x, y, w, h float64, pic Picture)
}

// AudioOutput lets you play sounds and music.
type AudioOutput interface {
	//TODO
}
