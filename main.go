package main

import (
	"github.com/disintegration/gift"
	farbfeld "github.com/hullerob/go.farbfeld"
	"image"
	"log"
	"os"
)

func main() {
	var src image.Image
	src, err := farbfeld.Decode(os.Stdin)
	if err != nil {
		log.Panic(err)
	}
	size := src.Bounds().Size()
	width := size.X
	height := size.Y
	g := gift.New(
		gift.Resize(width/4, height/4, gift.NearestNeighborResampling),
		gift.GaussianBlur(2),
		gift.Resize(width, height, gift.LinearResampling),
	)
	dst := image.NewRGBA(g.Bounds(src.Bounds()))
	g.Draw(dst, src)
	farbfeld.Encode(os.Stdout, dst)
}
