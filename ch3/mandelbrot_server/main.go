package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

var palette = [5]color.RGBA{
	color.RGBA{0, 100, 100, 255},
	color.RGBA{32, 107, 20, 255},
	color.RGBA{237, 255, 255, 255},
	color.RGBA{255, 170, 0, 255},
	color.RGBA{0, 2, 0, 255},
}

var params = map[string]float64{
	"xmin": -2,
	"ymin": -2,
	"xmax": +2,
	"ymax": +2,
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for param := range params {
			arg := r.FormValue(param)
			if arg == "" {
				continue
			}
			argFloat, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				http.Error(w, "Unable to parse query string argument", http.StatusBadRequest)
				return
			}
			params[param] = argFloat
		}

		drawCanvas(w)
	})

	// Plan 9 won't play well if localhost:8000 is used instead of
	// just port number
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func drawCanvas(out io.Writer) {
	const (
		width, height = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(params["ymax"]-params["ymin"]) + params["ymin"]
		for px := 0; px < width; px++ {
			x := float64(px)/width*(params["xmax"]-params["xmin"]) + params["xmin"]

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

	png.Encode(out, img)
}

func mandelbrot(z complex128) color.RGBA {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return palette[n%5]
		}
	}

	return color.RGBA{255, 255, 255, 255}
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
