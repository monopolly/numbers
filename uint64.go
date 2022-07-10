package numbers

//Uint64Split split to 1-2 int64 for hashing
func Uint64Split(id uint64) (res []int64) {
	if id < 9223372036854775807 {
		return []int64{int64(id)}
	}

	res = append(res, int64(9223372036854775807), int64(id-9223372036854775807))
	return
}
