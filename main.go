package main

import (
	"flag"
	"fmt"
	"image"
	"os"
	"strings"

	"github.com/nk2ge5k/go-dither/algo"

	"image/gif"
	"image/jpeg"
	"image/png"
)

func fileFormat(filename string) string {
	i := strings.LastIndex(filename, ".")
	if i < 0 {
		return ""
	}

	i++

	return filename[i:]
}

func encode(img image.Image, filename string) error {

	w, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0655)
	if nil != err {
		panic(err)
	}
	defer w.Close()

	f := fileFormat(filename)

	switch f {
	case "png":
		return png.Encode(w, img)
	case "jpg", "jpeg":
		return jpeg.Encode(w, img, nil)
	case "gif":
		return gif.Encode(w, img, nil)
	default:
		return fmt.Errorf("Unknown file format %s", f)
	}
}

func main() {
	var src_file, dst_file string

	flag.StringVar(&src_file, "s", "", "source file")
	flag.StringVar(&dst_file, "d", "", "destination file")
	flag.Parse()

	r, err := os.Open(src_file)
	if nil != err {
		panic(err)
	}
	defer r.Close()

	src, _, err := image.Decode(r)
	if nil != err {
		panic(err)
	}

	a := new(algo.FloydSteinberg)

	if err := encode(a.Dither(src), dst_file); nil != err {
		panic(err)
	}
}
