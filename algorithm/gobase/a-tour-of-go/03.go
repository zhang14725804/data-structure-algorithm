/*
	指针：：

	类型 *T 是指向 T 类型值的指针。其零值为 nil（var p *int）
	& 操作符会生成一个指向其操作数的指针。
	* 操作符表示指针指向的底层值。
*/
func point(){
	i,j:=13,9
	// 指向i
	p:=&i
	// 通过指针读取i的值
	fmt.Println(*p)
	// 通过指针设置i的值
	*p = 31
	fmt.Println(i)
	// 指向j
	p = &j
	// 通过指针对j进行运算
	*p=*p/3
	fmt.Println(j)
}

/*
	结构体：一个结构体(struct)就是一组字段(field)
	结构体字段：结构体字段使用点号来访问。
*/
func struct1Func(){
	type Struct1 struct{
		X int
		Y int
	}
	str:=Struct1{1, 2}
	str.X = 4
	fmt.Println(str.X)
}

/*
	结构体指针：结构体字段可以通过结构体指针来访问。

	有一个指向结构体的指针 p，那么可以通过 (*p).X 来访问其字段 X。不过这么写太啰嗦了，所以语言也允许我们使用隐式间接引用，直接写 p.X 就可以。
*/
func struct2Func(){
	type Struct1 struct{
		X int
		Y int
	}
	str:=Struct1{1, 2}
	p:=&str

	// (*p).X = 4
	p.X = 4
	fmt.Println(str)
}

/*
	结构体文法

	结构体文法通过直接列出字段的值来新分配一个结构体。

	使用 Name: 语法可以仅列出部分字段。（字段名的顺序无关。）

	特殊的前缀 & 返回一个指向结构体的指针。
*/
func struct3Func(){
	type Struct1 struct{
		X int
		Y int
	}
	v1 = Struct1{1, 2}  // 创建一个 Struct1 类型的结构体
	v2 = Struct1{X: 1}  // Y:0 被隐式地赋予
	v3 = Struct1{}      // X:0 Y:0
	p  = &Struct1{1, 2} // 创建一个 *Struct1 类型的结构体（指针）
}


/*
	数组：：

	类型 [n]T 表示拥有 n 个 T 类型的值的数组（var a [10]int）
	数组的长度是其类型的一部分，因此数组不能改变大小
*/
func array(){
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"	

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

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
	
}