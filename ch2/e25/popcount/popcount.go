// Popcount provides the PopCount method
// that determines population size of an integer with a loop
// Clears the rightmost nonzero bit each iteration
package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	var result int
	// x will continually decreased  until 0 due to
	// rightmost nonzero bit being cleared
	for x != 0 {
		x = x&(x-1)
		result++
	}
	return result
}