package main

import (
	"fmt"
	"math/rand"
	"sort"

	"github.com/deckarep/golang-set"
)

/**
二分查找方法
a 要查找的数组
b 要查找的int
 */
func BinarySearch(a []int, b int) int {
	var (
		low int = 0
		hight int = len(a) - 1
		counter = 0 // 统计查找的次数
	)

	for i := 0; i < len(a); i++ {
		middle := (low + hight) / 2 // 中间的值
		counter++
		if a[middle] > b {
			hight = middle - 1
		} else if a[middle] < b {
			low = middle + 1
		} else {
			fmt.Println("次数：", counter)
			return middle
		}
		if hight < low {
			fmt.Println("次数", counter)
			return -1
		}
	}
	fmt.Println("次数：", counter)
	return  -1
}

func main() {
	set := mapset.NewSet()

	for i := 0; i < 24; i++ {
		set.Add(rand.Intn(10000000))
	}

	setarr := set.ToSlice()
	bucket := make([]int, len(setarr))
	for i := 0; i < len(setarr); i++ {
		bucket[i] = setarr[i].(int)
	}
	n := bucket[1]
	sort.Ints(bucket)
	fmt.Println("数组：", bucket, n)
	fmt.Println("查找到的位置：", BinarySearch(bucket, 128162))
}
