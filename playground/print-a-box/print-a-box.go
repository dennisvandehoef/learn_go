package main

import (
	"flag"
	"fmt"
)

func main() {
	x := flag.Int("x", 5, "Width of the box")
	y := flag.Int("y", 3, "Height of the box")
	flag.Parse()

	for y_i := 1; y_i <= *y; y_i++ {
		for x_i := 1; x_i <= *x; x_i++ {
			if y_i == 1 || y_i == *y || x_i == 1 || x_i == *x {
				fmt.Printf("X")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}
