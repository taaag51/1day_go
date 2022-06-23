package main

import "fmt"

func getfactord(thenum int) {
	fmt.Println(thenum)
	for n := 0; n < thenum; n++ {
		switch {
		case n == 0 || n == 1:
		case thenum%n == 0:
			thenum /= n
			fmt.Printf("を%dで割ると、%d\n", n, thenum)
		default:

		}

	}
	fmt.Println()
}

func main() {
	getfactord(2310)
	getfactord(37789)
	getfactord(1256697)
}
