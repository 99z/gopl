package main

import (
	"fmt"
	"math"
)

const (
	width, height	= 600, 320				// canvas size in pixels
	cells			= 100					// number of grid cells
	xyrange			= 30.0					// axis ranges
	xyscale			= width / 2 / xyrange	// pixels per x or y unit
	zscale			= height * 0.4			// pixels per z unit
	angle			= math.Pi / 6			// angle of x, y axes
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, z := corner(i + 1, j)
			bx, by, _ := corner(i, j)
			cx, cy, _ := corner(i, j + 1)
			dx, dy, _ := corner(i + 1, j + 1)

			if (ax == 0 && ay == 0) ||
			   (bx == 0 && by == 0) ||
			   (cx == 0 && cy == 0) ||
			   (dx == 0 && dy == 0) {
				continue
			}

			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' " +
				"style='fill: #%06x'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, 0xFFFFFF ^ int(math.Max(0, (z * 50000))))
		}
	}

	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, float64) {
	// Find x,y at corner of cell i,j
	x := xyrange * (float64(i) / cells - 0.5)
	y := xyrange * (float64(j) / cells - 0.5)

	// Compute surface height z
	z := f(x, y)

	// It might be better to return NaN instead of 0 here
	// but would rather avoid using NaN at all
	if (math.IsInf(z, 0)) {
		return float64(0), float64(0), float64(0)
	}

	// Project x,y,z isometrically onto 2-D svg canvas sx,sy
	sx := width / 2 + (x - y) * cos30 * xyscale
	sy := height / 2 + (x + y) * sin30 * xyscale - z * zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}