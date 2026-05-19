package main

import "fmt"

/*
切片扩容
len == cap 时进行扩容
< 1024 时，扩容 2 倍
> 1024 时，扩容 1.25 倍
考察：cap扩容机制
*/
func Test_slice_1() {
	s := make([]int, 10, 11)
	s = append(s, 10)
	fmt.Println("s: %v, len of s: %d, cap of s: %d", s, len(s), cap(s))
	s = append(s, 10)
	fmt.Println("s: %v, len of s: %d, cap of s: %d", s, len(s), cap(s))
}

/*

初始化切片 s，长度为 10，容量为 12
截取切片 s 从索引 8 往后的内容，赋值给 s1
修改 s1[0] 的值
问题：这个修改是否会影响到 s？此时 s 的内容是什么？
考察：共享内存数组
*/

func Test_slice_2() {
	s := make([]int, 10, 12)
	s1 := s[8:]
	s1[0] = -1
	fmt.Println("s: %v", s)
}

/*
初始化切片 s，长度为 10，容量为 12
问题：访问 s[10] 是否会越界？
考察：数组越界
*/
func Test_slice_3() {
	s := make([]int, 10, 12)
	v := s[10]
	fmt.Println("v: %v", v)

}

/*
初始化切片 s，长度为 10，容量为 12
截取切片 s 从索引 8 往后的内容，赋值给 s1
修改 s1[0] 的值
问题：这个修改是否会影响到 s？此时 s 的内容是什么？
考察：共享内存数组,是否越界，扩容了，不再是共享底层数组了，slice 地址改变
*/
func Test_slice_4() {
	s := make([]int, 10, 12)
	s1 := s[8:]
	fmt.Println("扩容前")
	fmt.Printf("s  底层数组地址：%p\n", s)
	fmt.Printf("s1 底层数组地址：%p\n", s1)
	s1 = append(s1, []int{10, 11, 12}...)
	fmt.Println("扩容后")
	fmt.Printf("s  底层数组地址：%p\n", s)
	fmt.Printf("s1 底层数组地址：%p\n", s1)
	fmt.Println("s: %v")
	v := s[10]
	fmt.Println("v: %v", v)
	// 求问，此时数组访问是否会越界
}

/*
初始化切片 s，长度为 10，容量为 12
截取切片 s 从索引 8 往后的内容，赋值给 s1
修改 s1[0] 的值
问题：这个修改是否会影响到 s？此时 s 的内容是什么？
考察：共享内存数组,是否越界，扩容了，不再是共享底层数组了，slice 地址改变；如果未扩容，则会影响到原切片

扩容了 s1地址变了，
changeSlice 函数内s1的地址与外面的一样的，都是同一个底层数组地址，
不扩容能修改值，扩容后就不行了， 扩容后s1地址变了，不再是共享底层数组了，所以修改不了
你用fmt.Printf("%p\n", s1)打印的，不是切片变量s1自己的内存地址，而是它指向的底层数组的首地址，所以扩容后就会变


*/
func Test_slice_5() {
	s := make([]int, 10, 12)
	s1 := s[8:]
	fmt.Printf("s1 地址0：%p\n", &s1)
	fmt.Printf("s1 len----------: %d, cap: %d\n", len(s1), cap(s1))
	changeSlice(s1)
	fmt.Printf("s1 底层数组地址0：%p\n", s1)
	fmt.Println("s: %v", s)
	fmt.Println("s1: %v", s1)
}

func changeSlice(s1 []int) {
	fmt.Printf("s1 地址1：%p\n", &s1)
	fmt.Printf("s1 底层数组地址1：%p\n", s1)
	s1 = append(s1, 10)
	s1 = append(s1, 11)
	s1 = append(s1, 12)
	fmt.Printf("s1 len-----------: %d, cap: %d\n", len(s1), cap(s1))
	fmt.Printf("s1 底层数组地址2：%p\n", s1)
	fmt.Println("s1 cap: %d", cap(s1))
	s1[0] = -1
}

/*
Go 函数传参机制：切片作为参数传递时，传递的是结构体副本，函数内对切片变量本身的修改（如 len）仅影响副本，不影响外部变量。
*/
func Test_slice_6() {
	fmt.Println("Test_slice_6--------------------------------")
	s := make([]int, 10, 12)
	s1 := s[8:]
	changeSlice2(s1)
	fmt.Println("s: %v, len of s: %d, cap of s: %d", s, len(s), cap(s))
	fmt.Println("s1: %v, len of s1: %d, cap of s1: %d", s1, len(s1), cap(s1))
}

func changeSlice2(s1 []int) {
	s1 = append(s1, 10)
}

func Test_slice_7() {
	s := []int{0, 1, 2, 3, 4}
	s = append(s[:2], s[3:]...)
	fmt.Println("s: %v, len: %d, cap: %d", s, len(s), cap(s))
	v := s[4]
	fmt.Println("v: %v", v)
	// 是否会数组访问越界
}
func main() {
	Test_slice_1()
	Test_slice_2()
	//Test_slice_3()  //越界 ，多大长度就只能访问多大长度
	//Test_slice_4() //扩容了，不再是共享底层数组了，slice 地址改变
	Test_slice_5() //传入切片，修改切片，会影响到原切片，切片是引用传递，值会被修改
	Test_slice_6()
}
