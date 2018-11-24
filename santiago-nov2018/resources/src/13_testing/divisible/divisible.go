package divisible

// Sum returns the sum of numbers divisible by "by"
func Sum(top int, by ...int) int {
	var total int
	for i := 1; i <= top; i++ {
		for _, v := range by {
			if i%v == 0 {
				total += i
			}
		}
	}
	return total
}
