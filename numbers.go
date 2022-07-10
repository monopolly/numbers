package numbers

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
	"time"
	"unicode"
)

/*
func Uint32Bytes(i uint32) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, i)
	return b
}

func BytesUint32(b []byte) uint32 {
	return binary.LittleEndian.Uint32(b)
}
*/
func Uint64Bytes(i uint64) (r []byte) {
	r = make([]byte, 8)
	binary.LittleEndian.PutUint64(r, i)
	return
}

func BytesUint64(b []byte) (i uint64) {
	if b == nil {
		return 0
	}
	return binary.LittleEndian.Uint64(b)
}

//int64
func Int64Bytes(i int64) (r []byte) {
	r = make([]byte, 8)
	binary.LittleEndian.PutUint64(r, uint64(i))
	return
}

func BytesInt64(b []byte) (i int64) {
	if b == nil {
		return 0
	}
	return int64(binary.LittleEndian.Uint64(b))
}

//int
func IntBytes(i int) (r []byte) {
	r = make([]byte, 8)
	binary.LittleEndian.PutUint64(r, uint64(i))
	return
}

func BytesInt(b []byte) (i int) {
	if b == nil {
		return 0
	}
	return int(binary.LittleEndian.Uint64(b))
}

func Duration(d time.Duration) string {
	d = d.Round(time.Minute)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	return fmt.Sprintf("%02d:%02d", h, m)
}

func Format(n interface{}, delim ...byte) string {
	var fo byte = ' '
	if len(delim) > 0 {
		fo = delim[0]
	}
	in := fmt.Sprint(n)
	out := make([]byte, len(in)+(len(in)-2+int(in[0]/'0'))/3)
	if in[0] == '-' {
		in, out[0] = in[1:], '-'
	}

	for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
		out[j] = in[i]
		if i == 0 {
			return string(out)
		}
		if k++; k == 3 {
			j, k = j-1, 0
			out[j] = fo
		}
	}
}

func Formats(n int) string {
	if n < 1000 {
		return strconv.Itoa(n)
	}

	switch {
	case n < 1000000:
		return fmt.Sprintf("%dk", n/1000)
	case n < 1000000000:
		return fmt.Sprintf("%dm", n/1000000)
	case n < 1000000000000:
		return fmt.Sprintf("%db", n/1000000000)
	}

	return ""
}

//четное
func Even(number interface{}) (res bool) {
	switch number.(type) {
	case int:
		return number.(int)%2 == 0
	case int8:
		return number.(int8)%2 == 0
	case int16:
		return number.(int16)%2 == 0
	case int32:
		return number.(int32)%2 == 0
	case int64:
		return number.(int64)%2 == 0
	case uint:
		return number.(uint)%2 == 0
	case uint8:
		return number.(uint8)%2 == 0
	case uint16:
		return number.(uint16)%2 == 0
	case uint32:
		return number.(uint32)%2 == 0
	case uint64:
		return number.(uint64)%2 == 0
	default:
		return
	}
}

//нечетное
func Odd(number int) bool {
	return !Even(number)
}

func Int64FromAnyString(s interface{}) int64 {
	return int64(intfromstring(s))
}

func Int32FromAnyString(s interface{}) int32 {
	return int32(intfromstring(s))
}

func IntFromAnyString(s interface{}) int {
	return intfromstring(s)
}

func UintFromAnyString(s interface{}) uint {
	return uint(intfromstring(s))
}

func Uint32FromAnyString(s interface{}) uint32 {
	return uint32(intfromstring(s))
}

func Uint64FromAnyString(s interface{}) uint64 {
	return uint64(intfromstring(s))
}

//берет строку и выбирает из нее все цифры и превращает в число
// [[a,1234,s5],"some"] > 12345
func intfromstring(s interface{}) int {
	var b bytes.Buffer
	for _, letter := range fmt.Sprint(s) {
		switch {
		case unicode.IsDigit(letter):
			b.WriteRune(letter)
		default:
		}
	}
	num, err := strconv.Atoi(b.String())
	switch err {
	case nil:
		return num
	}
	return 0
}
func Uint32toBytes(v uint32) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, v)
	return buf.Bytes()
}

func BytesToUint32(in []byte) (r uint32) {
	if in == nil {
		return 0
	}
	rbuf := bytes.NewReader(in)
	binary.Read(rbuf, binary.LittleEndian, &r)
	return
}

func UintsToBytes32(vs []uint32) []byte {
	buf := make([]byte, len(vs)*4)
	for i, v := range vs {
		binary.LittleEndian.PutUint32(buf[i*4:], v)
	}
	return buf
}

func UintsToBytes64a(vs ...uint64) []byte {
	buf := make([]byte, len(vs)*8)
	for i, v := range vs {
		binary.LittleEndian.PutUint64(buf[i*8:], v)
	}
	return buf
}

func UintsToBytes64(vs []uint64) []byte {
	buf := make([]byte, len(vs)*8)
	for i, v := range vs {
		binary.LittleEndian.PutUint64(buf[i*8:], v)
	}
	return buf
}

func Ints64Bytes(vs []int64) []byte {
	buf := make([]byte, len(vs)*8)
	for i, v := range vs {
		binary.LittleEndian.PutUint64(buf[i*8:], uint64(v))
	}
	return buf
}

func Ints64Bytesa(vs ...int64) []byte {
	buf := make([]byte, len(vs)*8)
	for i, v := range vs {
		binary.LittleEndian.PutUint64(buf[i*8:], uint64(v))
	}
	return buf
}

func BytesToInts64(vs []byte) []int64 {
	if vs == nil {
		return nil
	}
	out := make([]int64, len(vs)/8)
	for i := range out {
		out[i] = int64(binary.LittleEndian.Uint64(vs[i*8:]))
	}
	return out
}

func BytesToUints64(vs []byte) []uint64 {
	if vs == nil {
		return nil
	}
	out := make([]uint64, len(vs)/8)
	for i := range out {
		out[i] = binary.LittleEndian.Uint64(vs[i*8:])
	}
	return out
}

func BytesToUints32(vs []byte) []uint32 {
	if vs == nil {
		return nil
	}
	out := make([]uint32, len(vs)/4)
	for i := range out {
		out[i] = binary.LittleEndian.Uint32(vs[i*4:])
	}
	return out
}

func UniqueInts(s *[]int) {
	un := make(map[interface{}]bool)
	var i int
	for _, x := range *s {
		if !un[x] {
			(*s)[i] = x
			i++
		}
		un[x] = true
	}

	(*s) = (*s)[:i]
	return
}

func UniqueUint(s *[]uint64) {
	un := make(map[uint64]bool)
	var i int
	for _, x := range *s {
		if !un[x] {
			(*s)[i] = x
			i++
		}
		un[x] = true
	}

	(*s) = (*s)[:i]
	return
}

func UniqueInt(s *[]int64) {
	un := make(map[interface{}]bool)
	var i int
	for _, x := range *s {
		if !un[x] {
			(*s)[i] = x
			i++
		}
		un[x] = true
	}

	(*s) = (*s)[:i]
	return
}

func UniqueInt32(s *[]int32) {
	un := make(map[interface{}]bool)
	var i int
	for _, x := range *s {
		if !un[x] {
			(*s)[i] = x
			i++
		}
		un[x] = true
	}

	(*s) = (*s)[:i]
	return
}

func UniqueUint32(s *[]uint32) {
	un := make(map[interface{}]bool)
	var i int
	for _, x := range *s {
		if !un[x] {
			(*s)[i] = x
			i++
		}
		un[x] = true
	}

	(*s) = (*s)[:i]
	return
}
