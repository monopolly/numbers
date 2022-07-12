package numbers

func Unit64Int64(u uint64) (a, b int64) {
	if u > 9223372036854775807 {
		a = 9223372036854775807
		c := u - 9223372036854775807
		if c > 0 {
			b = int64(c - 1)
		}

		return
	}
	return int64(u), 0
}
