package main

import "fmt"

func main() {
	var x, y, z, temp float64
	fmt.Println("请输入要比较的三个数：")
	fmt.Scan(&x, &y, &z)
	if x < y {
		temp = x
		x = y
		y = temp
	}
	if x < z {
		temp = z
		z = x
		x = temp
	}
	if y < z {
		temp = z
		z = y
		y = temp
	}
	fmt.Println("这三个数的从大到小的顺序是：", x, y, z)
}
