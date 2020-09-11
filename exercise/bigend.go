package main

import (
	"fmt"
	"unsafe"
)

func main() {
	const N int = int(unsafe.Sizeof(1))
	x :=0x08646655
	p := unsafe.Pointer(&x)
	p2 := (*[N]byte)(p)
	fmt.Println(p2, p, x, N)
	if p2[0] == 0 {
		fmt.Println("本机器：大端")
	} else {
		fmt.Println("本机器：小端")
	}


}