package log

import (
	"testing"
)

func TestWriter(t *testing.T) {

	f, err := NewWriter()
	if err != nil {
		t.Fatal(err)
	}

	defer f.Close()

	for i := 0; i < 10000000; i++ {
		f.Write([]byte("sdfsdfsdflkjsdlkjfklsdjflkjsdfkljlkfjlsdkjflksdjfljsdlkfjlksdjflksdjfljsljfs"))
	}

}
