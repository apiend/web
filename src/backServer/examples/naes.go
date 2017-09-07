package main

import (
	"fmt"
	"strconv"
)

var XorKey []byte = []byte{0xB2, 0x09, 0xBB, 0x55, 0x93, 0x6D, 0x44, 0x47}

type Xor struct {
}

// type m interface {
// 	enc(src string) string
// 	dec(src string) string
// }

func (a *Xor) enc(src string) string {
	var result string
	j := 0
	s := ""
	bt := []rune(src)
	fmt.Println(len(bt))
	for i := 0; i < len(bt); i++ {
		s = strconv.FormatInt(int64(byte(bt[i])^XorKey[j]), 16)

		if len(s) == 1 {
			s = "0" + s
		}
		result = result + (s)
		j = (j + 1) % 8
	}
	return result
}

func (a *Xor) dec(src string) string {
	var result string
	var s int64
	j := 0
	bt := []rune(src)
	fmt.Println(bt)
	for i := 0; i < len(src)/2; i++ {
		s, _ = strconv.ParseInt(string(bt[i*2:i*2+2]), 16, 0) // 取字符串每2位的16进制转换成10进制后异或运算
		result = result + string(byte(s)^XorKey[j])
		j = (j + 1) % 8
	}
	return result
}
func main() {
	xor := Xor{}
	fmt.Println(xor.enc("456465465465465465465cde"))
	fmt.Println(xor.dec("c17a957bbd0920699c27"))
	// 不支持中文加密
}
