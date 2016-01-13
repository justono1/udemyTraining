package main

import "fmt"

func main() {
	for i := 100000000; i < 100000100; i++ {
		fmt.Printf("%d \t %b \t %x \n", i, i, i)
	}
}
