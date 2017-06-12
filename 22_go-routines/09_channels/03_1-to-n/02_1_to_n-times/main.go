package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	n := 10
	c := make(chan int)
	done := make(chan bool)

	outMap := make(map[int]bool)

	go func() {
		for i := 0; i < 100000; i++ {
			c <- i
		}
		close(c)
	}()

	for i := 0; i < n; i++ {
		go func() {
			for n := range c {
				if _, ok := outMap[i]; !ok {
					outMap[i] = true
				}
				fmt.Println(i, ":", n)
			}
			done <- true
		}()
	}

	for i := 0; i < n; i++ {
		<-done
	}

	fmt.Println(outMap)
}
