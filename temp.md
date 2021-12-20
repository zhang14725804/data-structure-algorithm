```sh

package main

import "fmt"

func main() {
	// arr1 := []int{1, 2, 3}
	// arr2 := []int{3, 4, 5}
	// fmt.Println(diff(arr1, arr2))
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
}

/*
	第一题：
	1. 全局常量
	2. 全局变量
	3. init函数
	4. main函数
*/

/*
	第三题：
	return 语句并不是原子指令
	先给返回值赋值，然后调用defer表达式，最后才是返回到调用函数中。
*/
func f1() (r int) {
	/*
		r = 0
		defer r++
		return
		所以返回：1
	*/
	defer func() {
		r++
	}()
	return 0
}
func f2() (r int) {
	/*
		t := 5
		r = t
		defer t=t+5
		return
		所以返回：5
	*/
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}
func f3() (r int) {
	/*
		r = 1
		defer r=r+5 // 这里的r是值传递
		return
		返回 1
	*/
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

/*
	第二题：
*/
func diff(target []int, other []int) []int {
	res := make([]int, 0)
	hash := make(map[int]struct{}, 0)
	for _, v := range other {
		hash[v] = struct{}{}
	}
	for _, v := range target {
		if _, ok := hash[v]; !ok {
			res = append(res, v)
		}
	}
	return res
}

```