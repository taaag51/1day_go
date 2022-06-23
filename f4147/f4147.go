package main

import "fmt"

func main() {
	thenum := 4147

	fmt.Printf("%dは...\n", thenum)
	for n := 0; n < thenum; n++ {
		switch {
		case n == 0 || n == 1:
		case thenum%n == 0:
			thenum /= n
			fmt.Printf("%dで割ると、%d\n", n, thenum)
		default:

		}
	}
}
