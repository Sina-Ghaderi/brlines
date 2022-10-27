package lines

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

const sizesimpix = 20

const (
	white = iota
	darks
	bered
)

type BreLines struct {
	img     *image.RGBA
	col     map[byte]color.Color
	pngfile *os.File
	imgsize int
	pixsize int
}

func (p *BreLines) GetIMG() *image.RGBA { return p.img }

func NewBreLine(size int, path string) (*BreLines, error) {
	imageRGBA := image.NewRGBA(image.Rect(0, 0, size, size))
	cols := make(map[byte]color.Color)
	cols[darks] = color.Gray16{0xdada}
	cols[white] = color.White
	cols[bered] = color.RGBA{255, 0, 0, 255}

	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	return &BreLines{
		img: imageRGBA, col: cols, pngfile: file,
		pixsize: int(size / sizesimpix), imgsize: size}, err
}

func (p *BreLines) DrawMesh() {
	draw.Draw(
		p.img, p.img.Bounds(),
		&image.Uniform{p.col[white]}, image.Point{}, draw.Src,
	)

	var locx int
	var icor byte
	for currx := 0; currx < p.img.Bounds().Max.X; currx++ {
		var locy int
		for curr_y := 0; curr_y < sizesimpix; curr_y++ {
			fg := image.Rect(locx, locy, locx+p.pixsize, locy+p.pixsize)
			un := &image.Uniform{p.col[icor]}
			draw.Draw(p.img, fg, un, image.Point{}, draw.Src)
			locy += p.pixsize
			icor = 1 - icor
		}
		locx += p.pixsize
		icor = 1 - icor
	}
}

func (p *BreLines) setPixel(x, y int) {
	x = x * p.pixsize
	y = (sizesimpix - 1 - y) * p.pixsize
	draw.Draw(
		p.img, image.Rect(x, y, x+p.pixsize, y+p.pixsize),
		&image.Uniform{p.col[bered]}, image.Point{}, draw.Src,
	)

}

func (p *BreLines) WriteToFile() error {
	defer p.pngfile.Close()
	return png.Encode(p.pngfile, newXAndYAxis(p.imgsize).drawXandY(p.img))
}

func (p *BreLines) lessThanOrEqOne(x1, y1 int, x2, y2 int, dx, dy int) {
	dx = abs(dx)
	dy = abs(dy)

	Po := (2 * dy) - dx
	p.setPixel(x1, y1)
	xk, yk := x1, y1
	for k := x1; k < x2; k++ {
		if Po < 0 {
			xk++
			p.setPixel(xk, yk)
			Po = Po + (2 * dy)
		} else {
			xk++
			yk++
			p.setPixel(xk, yk)
			Po = Po + (2 * dy) - (2 * dx)
		}
	}
}

func (p *BreLines) greaterThanOne(x1, y1 int, x2, y2 int, dx, dy int) {
	dx = abs(dx)
	dy = abs(dy)
	Po := (2 * dx) - dy
	p.setPixel(x1, y1)
	xk, yk := x1, y1

	for k := y1; k < y2; k++ {
		if Po < 0 {
			yk++
			p.setPixel(xk, yk)
			Po = Po + (2 * dx)
		} else {
			xk++
			yk++
			p.setPixel(xk, yk)
			Po = Po + (2 * dx) - (2 * dy)
		}
	}
}

func (p *BreLines) lessThanOrEqNegOne(x1, y1 int, x2, y2 int, dx, dy int) {
	dx = abs(dx)
	dy = abs(dy)
	Po := (2 * dy) - dx
	p.setPixel(x1, y1)
	xk, yk := x1, y1

	for k := x1; k < x2; k++ {
		if Po < 0 {
			xk++
			p.setPixel(xk, yk)
			Po = Po + (2 * dy)
		} else {
			xk++
			yk--
			p.setPixel(xk, yk)
			Po = Po + (2 * dy) - (2 * dx)
		}
	}
}

func (p *BreLines) greaterThanNegOne(x1, y1 int, x2, y2 int, dx, dy int) {
	dx = abs(dx)
	dy = abs(dy)
	Po := (2 * dy) - dx
	p.setPixel(x1, y1)

	xk, yk := x1, y1

	for k := y1; k > y2; k-- {
		if Po < 0 {
			yk--
			p.setPixel(xk, yk)
			Po = Po + (2 * dx)
		} else {
			xk++
			yk--
			p.setPixel(xk, yk)
			Po = Po + (2 * dx) - (2 * dy)
		}
	}

}

func (p *BreLines) BresenhamLine(x1, y1 int, x2, y2 int) {

	// swap if point start is greater than point end
	if x1 > x2 {
		tempx := x2
		tempy := y2
		x2 = x1
		y2 = y1
		x1 = tempx
		y1 = tempy
	}

	dx := x2 - x1 // delat X
	dy := y2 - y1 // delta Y

	// for 0 < M <= 1, constant increment on x axis
	if dy <= dx && dy > 0 {
		p.lessThanOrEqOne(x1, y1, x2, y2, dx, dy)
		return

	}

	// for M > 1, constant increment on y axis
	if dy > dx && dy > 0 {
		p.greaterThanOne(x1, y1, x2, y2, dx, dy)
		return

	}

	// for 0 > M > -1 constant increment on x axis
	// decrement on y axis
	if dy >= -dx {
		p.lessThanOrEqNegOne(x1, y1, x2, y2, dx, dy)
		return

	}

	// for M < -1 constant decrement on y axis
	// increment on x axis
	if dy < -dx {
		p.greaterThanNegOne(x1, y1, x2, y2, dx, dy)
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
