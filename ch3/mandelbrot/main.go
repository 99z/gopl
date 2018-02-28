package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

var palette = [5]color.RGBA {
	color.RGBA{0, 100, 100, 255},
	color.RGBA{32, 107, 20, 255},
	color.RGBA{237, 255, 255, 255},
	color.RGBA{255, 170, 0, 255},
	color.RGBA{0, 2, 0, 255},
}

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py) / height * (ymax - ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px) / width * (xmax - xmin) + xmin

			subPixels := []color.RGBA{}
			for i := 0; i < 2; i++ {
				for j := 0; j < 2; j++ {
					// Now in four subpixels
					z := complex(x+float64((i/4)), y+float64((j/4)))
					subPixels = append(subPixels, mandelbrot(z))
				}
			}

			// Image point px, py represents complex value z
			// Color is the average of subpixel RGBA
			img.Set(px, py, getAverage(subPixels))
		}
	}

	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.RGBA {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n:= uint8(0); n < iterations; n++ {
		v = v * v + z
		if cmplx.Abs(v) > 2 {
			return palette[n % 5]
		}
	}

	return color.RGBA{255, 255, 255, 255}
}

func getAverage(colors []color.RGBA) color.RGBA {
	var r, g, b, a uint8
	for _, color := range colors {
		r += color.R
		g += color.G
		b += color.B
		a += color.A
	}

	return color.RGBA{r / 4, g / 4, b / 4, a / 4}
}