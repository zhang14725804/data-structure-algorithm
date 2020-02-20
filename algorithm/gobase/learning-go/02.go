import (
	"fmt"
	"unicode/utf8"
)

func basic(){
	// Go中的字符串是用双引号（“）括起来的一系列UTF-8字符。如果使用单引号（'）你的意思是一个字符（编码UTF-8）
	s := "Hello world!"
	//  invalid character literal (more than one character)
	// s2 := 'Hello world!'
	s2:='H'
	// Hello world! 72
	fmt.Println(s,s2)

	/*
		golang中string底层是通过byte数组实现的。
		中文字符在unicode下占2个字节，在utf-8编码下占3个字节，而golang默认编码正好是utf-8
	*/
	var str = "hello 你好"
	// len(str): 12
	fmt.Println("len(str):", len(str))

	//以下两种都可以得到str的字符串长度8
	
	//golang中的unicode/utf8包提供了用utf-8获取长度的方法
	fmt.Println("RuneCountInString:", utf8.RuneCountInString(str))

	//通过rune类型处理unicode字符
	fmt.Println("rune:", len([]rune(str)))

	/*
		golang中海油一个___byte___数据类型与rune相似，它们都是用来表示字符类型的变量类型。它们的不同在于：
		byte 等同于int8，常用来处理ascii字符
		rune 等同于int32,常用来处理unicode或utf-8字符
	*/

	/*
		一旦分配了变量，就不能更改字符串：Go中的字符串是不可变的

		cannot assign to s3[0]
		s3[0] = "H"

		要改变字符串中的单词需要如下操作，使用rune类型
	*/
	s3:="hello"
	c3:=[]rune(s3)
	// 一个字符（编码UTF-8）
	c3[0] = 'H'
	s4:=string(c3)
	// Hello
	fmt.Printf("%s\n", s4)
}