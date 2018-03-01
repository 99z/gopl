package main

import (
	"image"
	"image/color"
	"image/png"
	"math/big"
	"os"
)

var palette = [5]color.RGBA{
	color.RGBA{0, 100, 100, 255},
	color.RGBA{32, 107, 20, 255},
	color.RGBA{237, 255, 255, 255},
	color.RGBA{255, 170, 0, 255},
	color.RGBA{0, 2, 0, 255},
}

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin

			subPixels := []color.RGBA{}
			for i := 0; i < 2; i++ {
				for j := 0; j < 2; j++ {
					// Now in four subpixels
					xs, ys := x+float64((i/4)), y+float64((j/4))
					subPixels = append(subPixels, mandelbrot(xs, ys))
				}
			}

			// Image point px, py represents complex value z
			// Color is the average of subpixel RGBA
			img.Set(px, py, getAverage(subPixels))
		}
	}

	png.Encode(os.Stdout, img)
}

func mandelbrot(x, y float64) color.RGBA {
	// From Wikipedia's pseudocode for Mandelbrot set
	// Calculates without the use of complex numbers
	// Extremely slow compared to complex128
	const iterations = 20
	const contrast = 15
	var n uint8 = 0

	bigX, _ := big.NewFloat(x).Rat(nil)
	bigY, _ := big.NewFloat(y).Rat(nil)
	x0 := big.NewRat(0.0, 1)
	y0 := big.NewRat(0.0, 1)

	for ; (&big.Rat{}).Add((&big.Rat{}).Mul(x0, x0), (&big.Rat{}).Mul(y0, y0)).Cmp(big.NewRat(4, 1)) == -1 && n < iterations; n++ {
		xtemp := (&big.Rat{}).Add((&big.Rat{}).Sub((&big.Rat{}).Mul(x0, x0), (&big.Rat{}).Mul(y0, y0)), bigX)
		y0 = (&big.Rat{}).Add((&big.Rat{}).Mul(big.NewRat(2, 1), (&big.Rat{}).Mul(x0, y0)), bigY)
		x0 = xtemp
		n++
	}

	return palette[n%5]
}

func getAverage(colors []color.RGBA) color.RGBA {
	var r, g, b uint8
	for _, color := range colors {
		r += color.R / 5
		g += color.G / 5
		b += color.B / 5
	}

	return color.RGBA{r, g, b, 255}
}
