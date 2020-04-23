/*
	go中只有一种循环结构，for循环
	初始化语句通常为一句短变量声明，该变量声明仅在 for 语句的作用域中可见。
	Go 的 for 语句后面的三个构成部分外没有小括号， 大括号 { } 则是必须的

	for循环由三部分组成
*/

func forfunc(){
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

/*
	for循环，初始化语句和后置语句是可选的。
*/

func forfunc2(){
	sum := 1
	for ; sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

/*
	go中的"while"
*/
func whilefunc(){
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

/*
	无限循环
*/
func unlimited(){
	for{

	}
}

/*
	条件语句：
		if：表达式外无需小括号 ( ) ，而大括号 { } 则是必须的

		if的简短语句：
		同 for 一样， if 语句可以在条件表达式前执行一个简单的语句。
		该语句声明的变量作用域仅在 if 之内。

		if和else：在 if 的简短语句中声明的变量同样可以在任何对应的 else 块中使用。
*/
func ifFunc(){
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
func pow(x, n, lim float64) float64 {
	// if的简短语句
	if v := math.Pow(x, n); v < lim {
		return v
	}else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// 这里开始就不能使用 v 了
	return lim
}

/*
	switch语句
*/
func switchFunc(){
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func switchFunc2(){
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

/*
	没有条件的switch
*/
func switchFunc3(){
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

/*
	defer:defer 语句会将函数推迟到外层函数返回之后执行。
	
	defer栈：推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用
*/
func deferFunc(){
	defer fmt.Println("world")
	fmt.Println("hello")
}