

/*************************ListNode*******************************/
type ListNode struct {
	Val  int
	Next *ListNode
}

/*************************TreeNode*******************************/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/*************************reverse*******************************/
func reverse(a []interface{}) []interface{} {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}

/*************************math.pow*******************************/
func pow(base, exp int) int {
	res := 1
	for i := 0; i < exp; i++ {
		res *= base
	}
	return res
}

/**************************abs*******************************/
func absInt(num int) int {
	if num > 0 {
		return num
	}
	return -num
}

/**************************int最大值，最小值 -9223372036854775808 9223372036854775807*******************************/
const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

/**************************栈*******************************/
type Stack struct {
	x []interface{}
}

// 入栈
func (this *Stack) push(x interface{}) {
	this.x = append(this.x, x)
}

// 出栈
func (this *Stack) pop() {
	this.x = this.x[:len(this.x)-1]
	// node:=this.x[len(this.x)-1]
	// this.x = this.x[:len(this.x)-1]
	// return node
}

// 返回栈顶元素
func (this *Stack) top() interface{} {
	return this.x[len(this.x)-1]
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

// 头插入
func (this *Queue) push_front(x interface{}) {
	this.x = append([]interface{}{x}, this.x...)
}

// size
func (this *Queue) size() int {
	return len(this.x)
}

// 返回头元素
func (this *Queue) front() interface{} {
	return this.x[0]
}

// 头删除
func (this *Queue) pop_front() {
	this.x = this.x[1:]
}

// 返回尾元素
func (this *Queue) back() interface{} {
	return this.x[len(this.x)-1]
}

// 尾删除
func (this *Queue) pop_back() {
	this.x = this.x[:len(this.x)-1]
}

/**************************Set*******************************/
type Set struct {
	m map[interface{}]struct{}
}

func NewSet() *Set {
	s := &Set{}
	s.m = make(map[interface{}]struct{})
	return s
}

func (s *Set) Insert(key interface{}) {
	s.m[key] = struct{}{}
}

func (s *Set) Contains(key interface{}) bool {
	_, ok := s.m[key]
	return ok
}

func (s *Set) Remove(key interface{}) {
	delete(s.m, key)
}
