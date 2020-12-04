package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sorter.go/algorithm/bubblesort"
	"sorter.go/algorithm/qsort"
	"strconv"
	"time"
)

var infile *string = flag.String("i", "infile", "File contains values for sorting")
var outfile *string = flag.String("o", "outfile", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

func readValues(infile string) (values []int, err error) {
	/*
	1. 根据infile 读取文件
	2. 创建一个具有默认大小缓冲、从r读取的*Reader
	3. 初始化values 设置切片类型slice
	4.
	 */
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("打开输入文件失败", infile)
		return
	}
	defer file.Close()
	br := bufio.NewReader(file)

	values = make([]int, 0)
	// 创建循环 读取文件的内容
	for {
		// 一行行读取缓存文件内容
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}
		if isPrefix {
			fmt.Println("排长队，似乎出乎意料.")
			return
		}

		str := string(line) // 转换字符数组为字符串
		// Atoi是ParseInt(s, 10, 0)的简写 字符串表示的整数值
		value, err1 := strconv.Atoi(str)

		if err1 != nil {
			err = err1
			return
		}
		// 添加读取的内容到values
		values = append(values, value)
	}
	return
}

// 写到输出文件
func writeValues(values []int, outfile string) error {
	// 创建文件
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("创建输出文件失败", outfile)
		return err
	}
	defer file.Close()
	// 循环内容，写入文件
	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}

func main() {
	//从arguments中解析注册的flag。必须在所有flag都注册好而未访问其值时执行
	flag.Parse()

	if infile != nil {
		fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =", *algorithm)
	}

	values, err := readValues(*infile)
	if err != nil {
		fmt.Println(err)
	} else {
		t1 := time.Now()
		switch *algorithm {
			case "qsort":
				qsort.QuickSort(values)
			case "bubblesort":
				bubblesort.BubbleSort(values)
			default:
			fmt.Println("Sorting algorithm", *algorithm, "is either unknown orunsupported.")
		}
		t2 := time.Now()
		fmt.Println("排序过程时间", t2.Sub(t1), "to complete.")
		writeValues(values, *outfile)
		fmt.Println("Read values:", values)
	}
}
