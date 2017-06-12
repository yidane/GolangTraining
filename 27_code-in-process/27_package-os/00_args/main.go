package main

import (
	"fmt"
	"os"
)

func main() {
	entry := os.Args[1]
	fmt.Println(entry)

	for _, v := range os.Args {
		fmt.Println(v)
	}
}

/*
step 1:
go install

step 2 - from terminal:
programName arguments

OR
go run main.go '123'
*/
