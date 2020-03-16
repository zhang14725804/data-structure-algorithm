package ArrayList

/*
	自定义数组实现stack
*/
type StackArray interface{
	Clear() // 清空
	Size() int // 大小
	Pop() interface{} // 出战
	Push(data interface{}) // 入栈
	IsFull() bool // 是否满了
	IsEmpty() bool // 是否为空
}

// 用自定义ArrayList实现栈
type Stack struct{
	myArray *ArrayList
	capSize int // 最大范围
}

func NewArrayListStack() *Stack{
	stack := new(Stack)
	// 初始化dataSource
	stack.myArray = NewArrayList()
	stack.capSize = 10
	return stack
}

// 因为golang的垃圾回收机制，所以clear简单一些
func (stack *Stack) Clear(){
	stack.myArray.Clear()
	stack.capSize = 10
}

// 返回当前大小（不是len(stack.dataSource)）
func (stack *Stack) Size() int{
	return stack.myArray.TheSize
}

// 入栈
func (stack *Stack) Push(data interface{})  {
	if !stack.IsFull(){
		stack.myArray.Append(data)
	}
}

// 出战
func (stack *Stack) Pop() interface{} {
	if !stack.IsEmpty(){
		// 最后一个数据
		last := stack.myArray.dataStore[stack.myArray.TheSize-1]
		stack.myArray.Delete(stack.myArray.TheSize-1)
		return last
	}
	return nil
}

//判断是否满了
func (stack *Stack)IsFull()  bool {
	if stack.myArray.TheSize == stack.capSize{
		return true
	}
	return false
}

// 判断是否为空
func (stack *Stack) IsEmpty() bool  {
	if stack.myArray.TheSize == 0 {
		return true
	}
	return false
}

