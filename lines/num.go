package lines

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

const padding = 30

type xAndYAxis struct {
	img  *image.RGBA
	size int
}

func newXAndYAxis(size int) *xAndYAxis {
	imageRGBA := image.NewRGBA(image.Rect(0, 0, size+padding, size+padding))
	return &xAndYAxis{img: imageRGBA, size: size}
}

func (p *xAndYAxis) drawXandY(s *image.RGBA) *image.RGBA {
	draw.Draw(p.img, p.img.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)
	r := image.Rectangle{image.Point{padding, 0}, p.img.Rect.Max}
	draw.Draw(p.img, r, s, image.Point{0, 0}, draw.Src)

	p.vLine(padding-1, 0, s.Bounds().Max.Y)
	p.hLine(padding-1, s.Bounds().Max.Y+1, p.img.Bounds().Max.X)

	px := int(p.size / sizesimpix)
	ds := sizesimpix - 1
	for cx := 0; cx < sizesimpix; cx++ {
		p.addTextTo(int(padding/3), int(((cx+1)*px)-(px/3)), fmt.Sprintf("%02d", ds))
		ds--
	}

	for cx := 0; cx < sizesimpix; cx++ {
		p.addTextTo(int(cx*px+padding+(px/4)), int((p.img.Bounds().Max.Y)-padding/2), fmt.Sprintf("%02d", cx))
	}

	return p.img

}

func (p *xAndYAxis) hLine(x1, y, x2 int) {
	for ; x1 <= x2; x1++ {
		p.img.Set(x1, y, color.Gray16{0x0})
	}
}

func (p *xAndYAxis) vLine(x, y1, y2 int) {
	for ; y1 <= y2; y1++ {
		p.img.Set(x, y1, color.Gray16{0x0})
	}
}

func (p *xAndYAxis) addTextTo(x, y int, label string) {
	d := &font.Drawer{Dst: p.img,
		Src:  image.NewUniform(color.Gray16{0x0}),
		Face: basicfont.Face7x13,
		Dot:  fixed.Point26_6{X: fixed.I(x), Y: fixed.I(y)},
	}

	d.DrawString(label)
}
