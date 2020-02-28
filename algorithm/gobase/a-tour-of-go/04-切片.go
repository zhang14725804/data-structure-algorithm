/*
	切片（slice）：
	类型 []T 表示一个元素类型为 T 的切片

	切片通过两个下标来界定，即一个上界和一个下界，二者以冒号分隔：a[start:end]

	它会选择一个半开区间，包括第一个元素，但排除最后一个元素(这句话翻译有问题)
*/
func slice1(){
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// 3,5,7
	fmt.Println(primes[1:4])
}

/*
	切片的本质：切片就像数组的引用

	切片并不存储任何数据，它只是描述了底层数组中的一段。
	更改切片的元素会修改其底层数组中对应的元素。
	与它共享底层数组的切片都会观测到这些修改。
*/
func slice2(){
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

/*
	切片文法(初始化的意思吧)：：

	切片文法类似于没有长度的数组文法
*/
func slice3(){
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

/*
	切片的默认行为(上下边界问题)

	在进行切片时，你可以利用它的默认行为来忽略上下界。
	切片下界的默认值为 0，上界则是该切片的长度。
*/
func slice4(){
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

/*
	切片的长度和容量

	切片的长度就是它所包含的元素个数。
	切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。
	切片 s 的长度和容量可通过表达式 len(s) 和 cap(s) 来获取。
*/
func slice5(){
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	// 截取切片使其长度为 0
	s = s[:0]
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	// 拓展其长度
	s = s[:4]
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	// 舍弃前两个值
	s = s[2:]
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

/*
	nil 切片(todos：：这有什么意义呢)

	切片的零值是 nil。
	nil 切片的长度和容量为 0 且没有底层数组。
*/

func slice6(){
	var s []int
	// [] 0 0
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

/*
	用 make 创建切片
	切片可以用内建函数 make 来创建，这也是你创建动态数组的方式。

	make 函数会分配一个元素为零值的数组并返回一个引用了它的切片
	要指定它的容量，需向 make 传入第三个参数
*/
func slice7(){
	a := make([]int, 5)
	printSlice7("a", a)

	b := make([]int, 0, 5)
	printSlice7("b", b)

	c := b[:2]
	printSlice7("c", c)

	d := c[2:5]
	printSlice7("d", d)
}
func printSlice7(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

/*
	切片的切片

	切片可包含任何类型，甚至包括其它的切片
*/
func slice8(){
	// 创建一个井字板（经典游戏）
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// 两个玩家轮流打上 X 和 O
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

/*
	切片追加元素
*/