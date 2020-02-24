// 参数类型
func add1(a int,b int) int{
	return a+b
}

// 省略相同的参数类型
func add2(a,b int) int{
	return a+b
}

// 函数可以返回任意数量的返回值
func swap(x,y string)(string,string){
	return y,x
}

/*	
	命名返回值（第一次见这种用法）
	split(17) 返回 7 10
*/ 
func split(sum int)(x,y int){
	x = sum*4/9
	y = sum-x
	return
}

/*
	var 声明变量：
	var 语句用于声明一个变量列表，跟函数的参数列表一样，类型在最后
	var 语句可以出现在包或函数级别
	声明之后会给默认值
	如果初始化值已存在，则可以省略类型；变量会从初始值中获得类型

	":="短变量声明
	函数外的每个语句都必须以关键字开始（var, func 等等），因此 := 结构不能在函数外使用
*/

/*
	基本数据类型
	bool，string，int（各种int），byte（unit8别名），rune（int32别名），float，complex

	int, uint 和 uintptr 在 32 位系统上通常为 32 位宽，在 64 位系统上则为 64 位宽
*/
func t(){
	var char rune = '3'
	// 输出数据类型
	fmt.Printf("Type: %T",char)
	fmt.Printf(reflect.TypeOf(char))
}