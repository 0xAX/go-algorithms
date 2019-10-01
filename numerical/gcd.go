package numerical

// GCD returns gcd of x and y
func GCD(x uint, y uint) uint {

	var shift uint = 0

	if x == y {
		return x
	}

	if x == 0 {
		return y
	}

	if y == 0 {
		return x
	}

	for shift := 0; (x|y)&1 == 0; shift++ {
		x = x >> 1
		y = y >> 1
	}

	for (x & 1) == 0 {
		x = x >> 1
	}

	for y == 0 {

		for (y & 1) == 0 {
			y = y >> 1
		}

		if x > y {
			t := x
			x = y
			y = t
		}

		y = y - x

	}

	y = y << shift

	return y
}
