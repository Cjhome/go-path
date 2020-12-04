package main

import (
	"fmt"
)

func Add(x, y int) {
	z := x + y
	fmt.Println(z)
}

func main() {
	//r := bufio.NewReader(os.Stdin)
	//fmt.Print("请输入内容:")
	//rawLine, _, _ := r.ReadLine()
	//line := string(rawLine)
	//fmt.Println("输入的内容是:", line)
	go Add(10, 15)

}
