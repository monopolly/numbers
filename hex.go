package numbers

import (
	"math/big"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
)

func Hex2BigInt(h string) (res *big.Int) {
	res = new(big.Int)
	res.SetString(h, 16)
	return
}

func Text2BigInt(num string) (res *big.Int) {
	res = new(big.Int)
	res.SetString(num, 10)
	return
}

func Text2ModNScalar(num string) (res *secp256k1.ModNScalar) {
	i := Text2BigInt(num)
	b := make([]byte, 32)
	s := i.FillBytes(b)
	res = new(secp256k1.ModNScalar)
	res.SetByteSlice(s)
	return
}
