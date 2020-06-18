/*
	todo：bfs解法，思路不懂
*/ 
func numSquares(n int) int {
	q := &queue{}
	dist := make([]int,n+1)
	// 赋初始值
	for i := 0; i < len(dist); i++ {
		dist[i] = 1000000
	}
	q.push_back(0)
	dist[0] = 0
	// 他么什么意思
	for q.size()>0{
		t := q.pop_front()
		if t == n {
			return dist[t]
		}
		for i:=1;i*i+t <= n ;i++{
			j:=t+i*i
			if dist[j] > dist[t]+1{
				dist[j] = dist[t]+1
				q.push_back(j)
			}
		}
	}
	return 0
}


// 队列
type queue struct{
	Items []int
}

// 尾插入
func (this *queue) push_back(x int){
	this.Items = append(this.Items,x)
}

// 返回头元素
func (this *queue) front() int{
	return this.Items[0]
}
// size
func (this *queue) size() int{
	return len(this.Items)
}
// 头删除
func (this *queue) pop_front() int{
	res := this.Items[0]
	this.Items= this.Items[1:]
	return res
}