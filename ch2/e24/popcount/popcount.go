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
	for i := uint64(0); i < 64; i++ {
		// Test the value of the rightmost bit
		if x&1 == 1{
			result++
		}

		// Shift bit to the right by 1
		x >>= 1
	}
	return result
}