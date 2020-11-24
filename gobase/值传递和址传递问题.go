package main

import (
	"fmt"
)

// 数组和slice 值传递和址传递问题

func main() {
	a := [2]int{1, 2}
	alter1(a)
	fmt.Println("数组值传递hou:", a)
	b := [2]int{3, 4}
	alter2(&b)
	fmt.Println("数组址传递hou:", b)
	fmt.Println()
	c := []int{5, 6}
	alter3(c)
	fmt.Println("slice值传递hou:", c)
	d := []int{7, 8}
	alter4(&d)
	fmt.Println("slice址传递hou:", d)
}
func alter1(arr [2]int) {
	arr[0] = 2
	fmt.Println("数组值传递qian:", arr)
}
func alter2(arr *[2]int) {
	arr[0] = 2
	fmt.Println("数组址传递qian:", arr)
}

// 这个是因为他虽然表面是 值传递 其实slice底层 持有的是 数组的引用。说到底还是地址传递
func alter3(arr []int) {
	arr[0] = 33
	//arr=append(arr,11)
	fmt.Println("slice值传递qian:", arr)
}
func alter4(arr *[]int) {
	*arr = append(*arr, 12)
	fmt.Println("slice址传递qian:", arr)
}
