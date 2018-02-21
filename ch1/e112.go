// Minimal web server that can read 'cycles' from a query string
// and produce Lissajous figures accordingly
package main

import (
	"log"
	"net/http"
	"image"
	"image/color"
	"image/gif"
	"math"
	"math/rand"
	"time"
	"io"
	"strconv"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("cycles")
		cycles, err := strconv.ParseFloat(q, 64)
		if err != nil {
			os.Stderr.WriteString("Failed to parse cycles")
			os.Exit(1)
		}
	
		// Since lissajous takes io.Writer interface, can output
		// anywhere that supports it, including web
		lissajous(w, cycles)
	})
	
	// Plan 9 won't play well if localhost:8000 is used instead of
	// just port number
	log.Fatal(http.ListenAndServe(":8000", nil))
}

var palette = []color.Color{
	color.Black,
	color.RGBA{0x00, 0xff, 0x00, 0xff},
	color.RGBA{0xaa, 0xaa, 0xaa, 0xff},
	color.RGBA{0xff, 0x00, 0x00, 0xff},
	color.RGBA{0x00, 0xff, 0xff, 0xff},
	color.RGBA{0xff, 0xff, 0x00, 0xff}}

var colorIndices = []uint8{1, 2, 3, 4, 5} // don't include black

func lissajous(out io.Writer, cycles float64) {
	const (
		res		= 0.001
		size	= 100
		nframes = 64
		delay	= 8
	)
	
	freq := rand.Float64() + 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	rand.Seed(time.Now().Unix()) // generate random seed from current time
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t+= res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndices[rand.Intn(len(colorIndices))])
		}
		
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	
	gif.EncodeAll(out, &anim)
}