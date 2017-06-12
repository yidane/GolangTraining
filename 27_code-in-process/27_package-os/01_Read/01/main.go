package main

import (
	"fmt"
	"os"
)

func main() {
	src, err := os.Open("src.txt")
	if err != nil {
		panic(err)
	}
	defer src.Close()

	dst, err := os.Create("dst.txt")
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	bs := make([]byte, 5)
	n, err := src.Read(bs)
	for n > 0 {
		if err != nil {
			break
		}
		dst.Write(bs[:n])
		n, err = src.Read(bs)
	}

	fmt.Println(string(bs))
	fmt.Println("Complete")
}

// this is a limit reader
// we limit what is read
// see io.ReadFull for func similiar to (*File)Read
