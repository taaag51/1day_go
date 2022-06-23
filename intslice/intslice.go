package main

import "fmt"

func main() {
	intervals3 := []int{98, 125, 232, 147, 486}
	intervals3[1] -= 100

	fmt.Printf("このスライスが見ている配列の要素数は%d\n", len(intervals3))
	fmt.Printf("最後の要素は%d\n", intervals3[len(intervals3)-1])
	fmt.Printf("このスライスが見ている配列のキャパシティは%d\n", cap(intervals3))
	fmt.Printf("今や1番目の要素は%d\n", intervals3[1])
}
