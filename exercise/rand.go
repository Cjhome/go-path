package main

import (
	"fmt"
	"math/rand"
	"os"
	"text/tabwriter"
	"time"
)

func main() {
	rand.Seed(int64(time.Second))
	answers := []string{
		"It is certain",
		"It is decidedly so",
		"Without a doubt",
		"Yes definitely",
		"You may rely on it",
		"As I see it yes",
		"Most likely",
		"Outlook good",
		"Yes",
		"Signs point to yes",
		"Reply hazy try again",
		"Ask again later",
		"Better not tell you now",
		"Cannot predict now",
		"Concentrate and ask again",
		"Don't count on it",
		"My reply is no",
		"My sources say no",
		"Outlook not so good",
		"Very doubtful",
	}
	fmt.Println(rand.Int())
	fmt.Println("Magic 8-Ball says:", answers[rand.Intn(len(answers))])
	//创建生成器并播下种子。
	//通常应该使用一个非固定的种子，例如time.Now(). unixnano()。
	//使用固定的种子将在每次运行时产生相同的输出。
	r := rand.New(rand.NewSource(99))
	//这里的制表符帮助我们生成对齐的输出。
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	defer w.Flush()
	show := func(name string, v1, v2, v3 interface{}) {
		fmt.Fprintf(w, "%s\t%v\t%v\n", name, v1, v2, v3)
	}
	// Float32 and Float64 values are in [0, 1).
	show("Float32", r.Float32(), r.Float32(), r.Float32())
	show("Float64", r.Float64(), r.Float64(), r.Float64())
}
