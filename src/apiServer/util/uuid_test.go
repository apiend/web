package util

import (
	"fmt"
	"testing"
)

func TestGetUUID(t *testing.T) {
	// assert.Len(t, GetString(2), 2)
	// fmt.Println(GetString(20))
	r := GetUUID()
	fmt.Println(r)

}
