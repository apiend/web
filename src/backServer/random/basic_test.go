/*
	## Basic 一些简单随机数据
*/
package random

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var AnError = errors.New("assert.AnError general error for testing")

func TestGetNum(t *testing.T) {
	t.Parallel()

	// given
	// _, err := Natural(100)
	// val := 10

	// when
	_, err := Natural(0)
	if assert.Error(t, err) {
		// fmt.Println(err)
		assert.Equal(t, WrongTypeError, err)
	}
	_, err1 := Natural(-1)
	// then
	if assert.Error(t, err1) {
		// fmt.Println(err)
		assert.Equal(t, WrongTypeError, err1)
	}

	_, err2 := Natural(100)

	assert.NoError(t, err2)

}
