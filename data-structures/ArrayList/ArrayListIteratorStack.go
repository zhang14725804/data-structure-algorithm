

package ArrayList

/*
	迭代器版本的stack
*/
type StackArrayX interface{
	Clear() // 清空
	Size() int // 大小
	Pop() interface{} // 出战
	Push(data interface{}) // 入栈
	IsFull() bool // 是否满了
	IsEmpty() bool // 是否为空
}

// 用自定义ArrayList实现栈
type StackX struct{
	myArray *ArrayList
	// 添加迭代器（公共方法，属性）
	MyIt Iterator
}

func NewArrayListStackX() *StackX{
	stack := new(StackX)
	// 初始化dataSource
	stack.myArray = NewArrayList()
	stack.MyIt = stack.myArray.Iterator() // 迭代
	return stack
}

// 因为golang的垃圾回收机制，所以clear简单一些
func (stack *StackX) Clear(){
	stack.myArray.Clear()
	stack.myArray.TheSize = 0
}

// 返回当前大小（不是len(stack.dataSource)）
func (stack *StackX) Size() int{
	return stack.myArray.TheSize
}

// 入栈
func (stack *StackX) Push(data interface{})  {
	if !stack.IsFull(){
		stack.myArray.Append(data)
	}
}

// 出战
func (stack *StackX) Pop() interface{} {
	if !stack.IsEmpty(){
		// 最后一个数据
		last := stack.myArray.dataStore[stack.myArray.TheSize-1]
		stack.myArray.Delete(stack.myArray.TheSize-1)
		return last
	}
	return nil
}

//判断是否满了
func (stack *StackX)IsFull()  bool {
	if stack.myArray.TheSize >= 10{
		return true
	}
	return false
}

// 判断是否为空
func (stack *StackX) IsEmpty() bool  {
	if stack.myArray.TheSize == 0 {
		return true
	}
	return false
}
