package main

import "fmt"

func f() *int {
	a := 10   // 本来栈变量
	return &a // 指针返回 → 立刻逃逸到堆
}

func main() {
	a := f()
	fmt.Println(*a)
}
