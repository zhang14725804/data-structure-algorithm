/*
	方法

	Go 没有类。不过你可以为结构体类型定义方法。
	方法就是一类带特殊的 接收者 参数的函数。	
	方法接收者在它自己的参数列表内，位于 func 关键字和方法名之间。
*/
type ABS struct{
	X,Y float64
}
func (abs ABS) method() float64{
	return math.Sqrt(abs.X*abs.X+abs.Y*abs.Y)
}
func callMethod(){
	abs:=ABS{3,4}
	fmt.Println(abs.method())
}

/*
	方法即函数：：方法只是个带接收者参数的函数。
*/
type ABS struct{
	X,Y float64
}
func method(abs ABS) float64{
	return math.Sqrt(abs.X*abs.X+abs.Y*abs.Y)
}
func callMethod(){
	abs:=ABS{3,4}
	fmt.Println(method(abs))
}

/*
	也可以为非结构体类型声明方法。
	接收者的类型定义和方法声明必须在同一包内
*/
type Myfloat64 float64

func (f Myfloat64) Abs() float64{
	if f<0{
		return float64(-f)
	}
	return float64(f)
}
func call(){
	f:=Myfloat64(math.Sqrt2)
	fmt.Println(f.Abs())
}

/*
	指针接收者

	对于某类型 T，接收者的类型可以用 *T 的文法
	指针接收者的方法可以修改接收者指向的值。由于方法经常需要修改它的接收者，指针接收者比值接收者更常用。
*/