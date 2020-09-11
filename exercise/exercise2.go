package main

import "fmt"

func main() {
	var a rune = 176
	var b rune = 219
	fmt.Printf("%c%c%c%c%c\n", b, a, a, a, b)
	fmt.Printf("%c%c%c%c%c\n", a, b, a, b, a)
	fmt.Printf("%c%c%c%c%c\n", a, a, b, a, a)
	fmt.Printf("%c%c%c%c%c\n", a, b, a, b, a)
	fmt.Printf("%c%c%c%c%c\n", b, a, a, a, b)
}
