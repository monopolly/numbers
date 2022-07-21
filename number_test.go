package numbers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	a := assert.New(t)
	_ = a

	fmt.Println(Format(1253523423423552345))

	v := []uint64{17086447805432677, 17086447805432677}
	b := UintsToBytes64(v)
	d := BytesToUints64(b)
	fmt.Println(d)
}

func BenchmarkUintsToBytes64(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		UintsToBytes64([]uint64{1, 2, 3, 4, 5})
	}
}

func TestUnique(t *testing.T) {
	a := assert.New(t)
	_ = a

	slice := []uint64{1, 2, 2, 3, 3, 4, 3, 3, 4, 4, 4, 4, 4, 4, 5}
	UniqueUint(&slice)
	fmt.Println(slice)

	slices := []int64{1, 2, 2, 3, 3, 4, 3, 3, 4, 4, 4, 4, 4, 4, 5}
	UniqueInt(&slices)
	fmt.Println(slices)
}

func BenchmarkUnique(b *testing.B) {
	slice := []uint64{1, 2, 3, 4, 5, 5, 6, 7, 8}
	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		UniqueUint(&slice)
	}
}
