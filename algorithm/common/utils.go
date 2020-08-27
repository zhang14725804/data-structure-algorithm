package common

/*************************整型比较大小*******************************/
func compare(a, b int, max bool) int {
	// max 是否返回最大值
	if a > b {
		if max == true {
			return a
		}
		return b
	}
	if max == true {
		return b
	}
	return a
}

/**************************int最大值，最小值 -9223372036854775808 9223372036854775807*******************************/
const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

/**************************冒泡排序*******************************/
func sort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

/**************************栈*******************************/
type Stack struct {
	x []interface{}
}

// 入栈
func (this *Stack) push(x interface{}) {
	this.x = append(this.x, x)
}

// 出栈,返回栈顶元素
func (this *Stack) pop() interface{} {
	top := this.x[len(this.x)-1]
	this.x = this.x[:len(this.x)-1]
	return top
}

// 栈大小
func (this *Stack) size() int {
	return len(this.x)
}

// 判空
func (this *Stack) empty() bool {
	return len(this.x) == 0
}

/**************************队列*******************************/
type Queue struct {
	x []interface{}
}

// 尾插入
func (this *Queue) push_back(x interface{}) {
	this.x = append(this.x, x)
}

// 返回头元素
func (this *Queue) front() interface{} {
	return this.x[0]
}

// size
func (this *Queue) size() int {
	return len(this.x)
}

// 头删除
func (this *Queue) pop_front() {
	this.x = this.x[1:]
}
