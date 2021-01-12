package StackArray

// 数组大小
const CAP_SIZE int = 1000
type StackArray interface{
	Clear() // 清空
	Size() int // 大小
	Pop() interface{} // 出战
	Push(data interface{}) // 入栈
	IsFull() bool // 是否满了
	IsEmpty() bool // 是否为空
}

type Stack struct{
	dataSource []interface{}
	capSize int // 最大范围
	currentSize int // 实际使用大小
}

func NewStack() *Stack{
	stack := new(Stack)
	// 初始化dataSource
	stack.dataSource = make([]interface{},0,CAP_SIZE)
	stack.capSize = CAP_SIZE
	stack.currentSize = 0
	return stack
}

// 因为golang的垃圾回收机制，所以clear简单一些
func (stack *Stack) Clear(){
	stack.dataSource = make([]interface{},0,CAP_SIZE)
	stack.currentSize = 0
	stack.capSize = CAP_SIZE
}

// 返回当前大小（不是len(stack.dataSource)）
func (stack *Stack) Size() int{
	return stack.currentSize
}

// 入栈
func (stack *Stack) Push(data interface{})  {
	if !stack.IsFull(){
		stack.dataSource = append(stack.dataSource,data)
		stack.currentSize++
	}
}

// 出战
func (stack *Stack) Pop() interface{} {
	if !stack.IsEmpty(){
		// 最后一个数据
		last := stack.dataSource[stack.currentSize - 1]
		// 删除最后一个数据
		stack.dataSource = stack.dataSource[:stack.currentSize-1]
		stack.currentSize--
		return last
	}
	return nil
}

//判断是否满了
func (stack *Stack)IsFull()  bool {
	if stack.currentSize == stack.capSize{
		return true
	}
	return false
}

// 判断是否为空
func (stack *Stack) IsEmpty() bool  {
	if stack.currentSize == 0 {
		return true
	}
	return false
}
