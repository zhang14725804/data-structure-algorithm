package main

import "fmt"
import "./ArrayList"
import "./StackArray"

func main1(){
	// todos：：定义接口对象，复制的对象比u实现接口的所有方法（这他么怎么理解）
	// list := NewArrayList()
	var list ArrayList.List = ArrayList.NewArrayList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Delete(2)
	fmt.Println(list)
}

func main2(){
	// todos：：定义接口对象，复制的对象比u实现接口的所有方法（这他么怎么理解）
	// list := NewArrayList()
	var list ArrayList.List = ArrayList.NewArrayList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Insert(2, 4)
	/*
		超出10就会出问题
		[<nil> 6 <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil>]
	*/
	for i:=0;i<11;i++{
		list.Insert(1, 6)
		fmt.Println(list)
	}
	// [1 2 3]
	fmt.Println(list)
}


func main3(){
	// todos：：定义接口对象，复制的对象比u实现接口的所有方法（这他么怎么理解）
	// list := NewArrayList()
	var list ArrayList.List = ArrayList.NewArrayList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	fmt.Println(list)
	// 使用迭代器
	for it:=list.Iterator();it.HasNext();{
		item,_:=it.Next()
		if item == 3{
			it.Remove()
		}
		fmt.Println(item)
	}
	fmt.Println(list)
}

func main4()  {
	stack := StackArray.NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

	/*
		注意：入栈之后立即出栈，可以实现先进先出（就是一个队列）
	*/
	stack.Push(5)
	fmt.Println(stack.Pop())
	stack.Push(6)
	fmt.Println(stack.Pop())
}

func main5(){
	// 使用自定义数组实现的栈
	stack := ArrayList.NewArrayListStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

	/*
		注意：入栈之后立即出栈，可以实现先进先出（就是一个队列）
	*/
	stack.Push(5)
	fmt.Println(stack.Pop())
	stack.Push(6)
	fmt.Println(stack.Pop())
}

func main6(){
	// 使用自定义数组实现的栈 迭代器版本
	stack := ArrayList.NewArrayListStackX()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	// 迭代器循环()
	for it:=stack.MyIt;it.HasNext();{
		item,_ := it.Next()
		fmt.Println(item)
	}
}

// 简单的递归
func Add(num int) int{
	if num == 0{
		return 0
	}
	return num + Add(num-1)
}

/*
	如何用栈模拟递归（所有的递归都可以转成 循环+栈）
	遍历文件目录（那不就是递归么）
*/
func main(){
	fmt.Println(Add(5))
	stack := StackArray.NewStack()
	stack.Push(5)
	last:=0
	for !stack.IsEmpty(){
		data:=stack.Pop()
		if data.(int) > 0 {
			last += data.(int)
			stack.Push(data.(int) - 1)
		}else{
			last+=0
		}
	}
	fmt.Println(last)
}