package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

func main() {
	TestFile := "123.txt"
	infile, err := os.Open(TestFile)
	if err == nil {
		md5h := md5.New()
		io.Copy(md5h, infile)
		fmt.Printf("%x %s\n", md5h.Sum([]byte("")), TestFile)

		sha1h := sha1.New()
		io.Copy(sha1h, infile)
		fmt.Printf("%x %s\n",sha1h.Sum([]byte("")), TestFile)
	} else {
		fmt.Println(err)
		os.Exit(1)
	}
}
