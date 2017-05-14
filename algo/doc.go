package algo

import (
	"image"
	"io"
)

type Algo interface {
	Dither(r io.Reader) image.Image
}
