/*
	给定一个数组和滑动窗口的大小，找出所有滑动窗口里的最大值
	todo：单调队列问题（没懂😅）
*/
func maxInWindows(nums []int, k int) []int {
	res :=make([]int,0)
	// 支持两端插入删除的队列（存储元素下标）
	q := &queue{}
	for i:=0; i<len(nums); i++{
		// 队头
		for q.size() > 0 && q.front() <= i-k{
			q.pop_front()
		}
		for q.size()>0 && nums[q.back()] <= nums[i]{
			q.pop_back()
		}
		q.push_back(i)
		if i >= k-1{
			res = append(res,nums[q.front()])
		}
	}
	return res
}

// 队列
type queue struct{
	x []int
}

// 尾插入
func (this *queue) push_back(x int){
	this.x = append(this.x,x)
}

// 头插入 (注意这里)
func (this *queue) push_front(x int){
	this.x = append([]int{x}, this.x...)
}

// size
func (this *queue) size() int{
	return len(this.x)
}

// 返回头元素
func (this *queue) front() int{
	return this.x[0]
}

// 头删除
func (this *queue) pop_front(){
	this.x= this.x[1:]
}

// 返回头元素
func (this *queue) back() int{
	return this.x[len(this.x)-1]
}

// 尾删除
func (this *queue) pop_back(){
	this.x= this.x[:len(this.x)-1]
}
