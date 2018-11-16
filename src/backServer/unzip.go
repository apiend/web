package main

import (
	// gounzip "github.com/duncanhall/gounzip"

	"compress/flate"
	"fmt"

	"github.com/mholt/archiver"
)

func main() {

	z := archiver.Zip{
		CompressionLevel:       flate.DefaultCompression,
		MkdirAll:               true,
		SelectiveCompression:   true,
		ContinueOnError:        false,
		OverwriteExisting:      false,
		ImplicitTopLevelFolder: false,
	}

	// err := z.Walk("file/demo.zip", func(f archiver.File) error {
	// 	zfh, ok := f.Header.(zip.FileHeader)
	// 	if ok {
	// 		fmt.Println("Filename:", zfh.Name)
	// 		err2 := z.Extract("file/demo.zip", zfh.Name, "file/demo/")
	// 		fmt.Println(err2)
	// 	}
	// 	return nil
	// })

	err := z.Unarchive("file/demo.zip", "file/demo/")

	fmt.Println(err)
	// fmt.Println()
}
