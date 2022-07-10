package numbers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUint64Split(t *testing.T) {
	a := assert.New(t)
	_ = a

	fmt.Println(Uint64Split(18446744071709551615))
}
