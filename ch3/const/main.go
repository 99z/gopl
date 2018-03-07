package main

import "fmt"

const (
	KB = 1000
	MB = KB * KB
	GB = MB * KB
	TB = GB * KB
	PB = TB * KB
	EB = PB * KB
	ZB = EB * KB
	YB = ZB * KB
)

func main() {
	fmt.Printf(`KB: %d
MB: %d
GB: %d
TB: %d
PB: %d
EB: %d
ZB: Too big!
YB: Too big!`+"\n",
		KB, MB, GB, TB, PB, EB)
}
