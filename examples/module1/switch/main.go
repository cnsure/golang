package main

import (
	"flag"
	"fmt"
)

func main() {
	num := flag.Int("num", 1, "your num")
	flag.Parse()
	switch *num {
	case 1:
	case 2:
		fmt.Println("hello ")
	case 3:
		fmt.Println("world")
	default:
		fmt.Println("default")
	}
}
