package util

import (
	"errors"
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

var AnError = errors.New("assert.AnError general error for testing")

func Test(t *testing.T) {
	assert.Len(t, GetString(2), 2)
	fmt.Println(GetString(20))
	// fmt.Println(String(2))
	r := NewRandom()
	// fmt.Println(r.String(6, Hex))
	assert.Regexp(t, regexp.MustCompile("[0-9]+$"), r.String(8, Numeric))
}
