package main

import "fmt"

func main() {
	intervals := [5]int{98, 125, 232, 147, 486}
	fmt.Printf("1番目の要素は%d\n", intervals[1])
	fmt.Printf("0番目の要素は%d\n", intervals[0])
	fmt.Printf("2番目と3番目の要素を足すと%d\n", intervals[2]+intervals[3])
	intervals[1] -= 100
	fmt.Printf("今や1番目の要素は%d\n", intervals[1])
}
