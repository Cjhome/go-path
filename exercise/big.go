package main

import (
	"fmt"
	"math/big"
	"math/cmplx"
)

func main() {
	fmt.Println(big.MaxBase)
	const x int64 = 150000000
	fmt.Println(big.NewInt(x))
	fmt.Println(big.Float{})
	fmt.Println(big.Above)
	fmt.Println(big.Below)
	fmt.Println(big.ErrNaN{})
	fmt.Println(big.MaxExp)
	fmt.Println(big.Word(5))

	fmt.Println("cmplx", cmplx.NaN())
	fmt.Println("cmplx", cmplx.IsNaN(0 + 4))
}
