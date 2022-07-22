package numbers

import (
	"fmt"
	"strconv"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func Format(v int) (res string) {
	p := message.NewPrinter(language.English)
	return p.Sprintf("%d", v)
}

func Formats(n int) string {
	if n < 1000 {
		return strconv.Itoa(n)
	}
	switch {
	case n < 1000000:
		return fmt.Sprintf("%dk", n/1000)
	default:
		return fmt.Sprintf("%dm", n/1000000)
	}
}
