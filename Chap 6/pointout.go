package main

import "fmt"

func main() {

	inta := 5
	adinta := &inta

	bdinta := &inta
	*bdinta = 9

	fmt.Printf("adinta のアドレスは %p\n", &adinta)
	fmt.Printf("bdinta のアドレスは %p\n", &bdinta)
}
