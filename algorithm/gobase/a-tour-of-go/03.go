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

