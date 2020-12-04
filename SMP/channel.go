package main

import "fmt"

func Count1(ch chan int, index int) {
	ch <- index * index
	fmt.Println("Counting")
}

func main() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count1(chs[i], i)
	}

	for _, ch := range chs {
		a := <- ch
		fmt.Println(a)
	}
}
