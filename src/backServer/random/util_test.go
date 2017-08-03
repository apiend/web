package random

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Len(t, String(2), 2)
	// fmt.Println(String(20))
	// fmt.Println(String(2))
	r := New()
	// fmt.Println(r.String(6, Hex))
	assert.Regexp(t, regexp.MustCompile("[0-9]+$"), r.String(8, Numeric))
}
