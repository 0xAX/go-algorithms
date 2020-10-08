package numerical

//O(log n) function for pow(x, y)
func FastPow(n uint, power uint) uint {
	var res uint = 1
	for power > 0 {

		if (power & 1) != 0 {
			res = res * n
		}

		power = power >> 1
		n = n * n
	}
	return res
}
