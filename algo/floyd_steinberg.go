package algo

import (
	"image"
	"image/color"
)

const (
	c_B int64 = 0
	c_W int64 = 65535
	c_G int64 = c_W / 2
)

type FloydSteinberg struct {
	Error [2][]int64
}

func colorVal(c color.Color) int64 {
	v, _, _, _ := color.Gray16Model.Convert(c).RGBA()
	return int64(v)
}

func (a *FloydSteinberg) getColor(c color.Color, x, y int) color.Color {
	var v int64 = colorVal(c) + a.Error[0][x]

	y = 0

	if v >= c_G {
		v -= c_W
		c = color.White
	} else {
		v -= c_B
		c = color.Black
	}

	for _, d := range [4][3]int{
		{y, x + 1, 7}, {y + 1, x - 1, 3}, {y + 1, x, 5}, {y + 1, x + 1, 1},
	} {
		y, x, m := d[0], d[1], d[2]

		if x >= 0 && len(a.Error[y]) > x {
			a.Error[y][x] += (v / 16) * int64(m)
		}
	}

	return c
}

func (a *FloydSteinberg) Dither(i image.Image) image.Image {

	b := i.Bounds()
	w, h := b.Max.X, b.Max.Y

	for u := 0; u < 2; u++ {
		a.Error[u] = make([]int64, w)
	}

	out := image.NewGray16(b)

	for y := b.Min.Y; y < h; y++ {
		for x := b.Min.X; x < w; x++ {
			out.Set(x, y, a.getColor(i.At(x, y), x, y))
		}

		a.Error[0] = a.Error[1]
		a.Error[1] = make([]int64, w)
	}

	return out
}
