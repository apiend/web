/*
	## Basic 一些简单随机数据
*/
package random

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var WrongTypeError = errors.New("wrong type")

// 统一的变量
var R *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

/**
	生成 bool 值
**/
// 返回一个随机的布尔值。
func Boolean() bool {
	return R.Float32() >= 0.5
}

/**
	生成 数字
**/
// 返回一个随机的自然数（大于等于 0 的整数）。
func Natural(max int) (int, error) {
	if max > 0 {
		return R.Intn(max), nil
	}
	return 0, WrongTypeError
}

// 返回一个随机的整数。 min 最小值  max 最大值
func Integer(min, max int) int {
	randNum := R.Intn(max-min) + min
	return randNum
}

/**
	生成 随机字符串 英文
**/
func RandString(length int) string {
	rand.Seed(time.Now().UnixNano())
	rs := make([]string, length)
	for start := 0; start < length; start++ {
		t := rand.Intn(3)
		if t == 0 {
			rs = append(rs, strconv.Itoa(rand.Intn(10)))
		} else if t == 1 {
			rs = append(rs, string(rand.Intn(26)+65))
		} else {
			rs = append(rs, string(rand.Intn(26)+97))
		}
	}
	return strings.Join(rs, "")
}
