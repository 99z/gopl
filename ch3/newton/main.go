package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

var palette = [3]color.RGBA{
	color.RGBA{235, 160, 160, 255},
	color.RGBA{160, 235, 160, 255},
	color.RGBA{160, 235, 235, 255},
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
			z := complex(x, y)

			roots := []complex128{
				complex(1, 0),
				complex(-0.5, math.Sqrt(3)/2),
				complex(-0.5, -math.Sqrt(3)/2),
			}

			// Image point px, py represents complex value z
			img.Set(px, py, newton(z, roots))
		}
	}

	png.Encode(os.Stdout, img)
}

func newtonFunc(z complex128) complex128 {
	return cmplx.Pow(z, 3) - 1
}

func newtonDerivative(z complex128) complex128 {
	return 3 * cmplx.Pow(z, 2)
}

func newton(z complex128, roots []complex128) color.RGBA {
	const maxIterations = 200

	for n := uint8(0); n < maxIterations; n++ {
		z -= newtonFunc(z) / newtonDerivative(z)
		tolerance := 0.000001

		for index, root := range roots {
			diff := z - root
			if cmplx.Abs(diff) < tolerance {
				return palette[index]
			}
		}
	}

	return color.RGBA{0, 0, 0, 255}
}
