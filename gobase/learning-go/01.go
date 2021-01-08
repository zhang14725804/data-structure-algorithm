 
import "fmt"

/*

 */
func variable() {
	// int类型具有适合您计算机的长度，这意味着在32位计算机上为32位，在64位计算机上为64位
	var a int
	var b bool
	a = 1
	b = false
	// 此形式只能在函数内部使用
	c := 2
	d := true
	// cannot use a (type int) as type string in argument to fmt.Printf
	// fmt.Printf(a,b,c,d)
	fmt.Println(a, b, c, d)

	var (
		e int
		f bool
	)
	e = 3
	f = true
	fmt.Println(e, f)

	/*
		如果要明确说明长度，也可以使用 int32或设置uint32。对于（符号和无符号）整数的完整列表 int8，int16，int32，int64和byte，uint8，uint16，uint32， uint64，与byte作为一个别名uint8。
		对于浮点值，有float32和float64（没有float类型）。在32位架构上，64位整数或浮点值始终为 64位
	*/
	var g int
	var h int32
	// cannot use g + g (type int) as type int32 in assignment
	h = g + g
	h = h + 5
	fmt.Println(h)

	/*
		常量
		todos：：下面是什么骚操作
	*/
	const (
		a1 = iota
		a2
	)
	// 0 1
	fmt.Println(a1, a2)

	/*  字符串  */

}