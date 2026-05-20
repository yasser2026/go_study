package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

/*
指针返回必逃逸
*/
func f() *int {
	a := 10   // 本来栈变量
	return &a // 指针返回 → 立刻逃逸到堆
}

/*
闭包引用必逃逸
需要持久保存状态
需要动态生成函数
写中间件、装饰器、回调
轻量封装私有变量
*/
func f1() func() {
	a := 20
	return func() {
		fmt.Println(a) // 闭包引用 → 逃逸
	}
}

/*
动态切片易逃逸
*/
func f2() {
	s := make([]int, 10) // n 不确定 → 逃逸
	s[0] = 10
}

/*
接口赋值会逃逸,给接口interface{}赋值会逃逸
*/
func f3() {
	var i interface{} //因为interface{}是动态类型，所以会逃逸
	a1 := Person{Name: "张三", Age: 18}
	i = a1 // 赋值给 interface → 逃逸

	fmt.Println(i)  //只是不会打印二次逃逸，其实已经逃逸了
	fmt.Println(10) //interface{}可以存储任何类型，所以会逃逸
}

/*
通道传指针会逃逸
*/
func f4(ch chan *int) {
	a1 := 10
	ch <- &a1 // 指针进 channel → 逃逸
}

func f5(ch chan *int) {
	a1 := <-ch
	fmt.Println(*a1)
}
func main() {
	a := f()
	fmt.Println(*a)

	f1()()
	f2()
	f3()

	ch := make(chan *int)
	f4(ch)
	f5(ch)
}
