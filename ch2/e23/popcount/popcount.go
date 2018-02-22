// Popcount provides the PopCount method
// that determines population size of an integer with a loop
package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	var result int
	for i := uint64(0); i < 8; i++ {
		result += int(pc[byte(x>>(i*8))])
	}
	return result
}