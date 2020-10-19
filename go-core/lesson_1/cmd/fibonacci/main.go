package main

import (
	"flag"
	"fmt"
	fibo "go-core/lesson_1/pkg/fibonacci"
)

var nFlag = flag.Int("n", 1, "positive numbers up to 20")

func main() {
	flag.Parse()
	n := *nFlag
	if n < 0 || n > 20 {
		fmt.Println("A program works with numbers only from 0 to 20.")
		return
	}
	fmt.Println("your n = ", n)

	result := fibo.Num(n)
	fmt.Println("the number of fibonacci = ", result)
}
