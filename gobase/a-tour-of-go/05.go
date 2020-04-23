/*
	range

	for 循环的 range 形式可遍历切片或映射
	第一个值为当前元素的下标，第二个值为该下标所对应元素的一份副本。

	可以将下标或值赋予 _ 来忽略它
*/
func rangeFunc(){
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("%d = %d\n", i, v)
	}
}

/*
	映射

	映射的零值为 nil 。nil 映射既没有键，也不能添加键

	删除元素：delete(m, key)
	通过双赋值检测某个键是否存在：elem, ok = m[key]

*/


/*
	函数值

	函数也是值。它们可以像其它值一样传递。
	函数值可以用作函数的参数或返回值
*/

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func funcFunc() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

/*
	函数闭包

	Go 函数可以是一个闭包。闭包是一个函数值，它引用了其函数体之外的变量。该函数可以访问并赋予其引用的变量的值，换句话说，该函数被这些变量“绑定”在一起
*/
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func funcClosure() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}