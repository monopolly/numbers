package numbers

import "math/big"

func HexBigInt(h string) (res *big.Int) {
	res = new(big.Int)
	res.SetString(h, 16)
	return
}
