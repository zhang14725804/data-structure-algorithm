/*
	todo：宽搜问题：bfs宽度优先遍历（也可以使用dfs深度优先遍历，容易溢出）
*/
// 返回坐标数位之和
func getSingleNum(x int) int {
	s:=0
	for x > 0{
		s += x%10
		x /= 10
	}
	return s
}
// 横纵坐标
type Pair struct{
	X,Y int
}

func getSum(p Pair) int {
	return getSingleNum(p.X) + getSingleNum(p.Y)
}

func movingCount(threshold, rows, cols int) int {
	res := 0
	if rows == 0 || cols == 0{
		return res
	}
	st := make([][]bool,rows)
	for i:=0;i<rows;i++  {
		st[i] = make([]bool,cols)
	}

	q := &queue{x:nil}
	// 起始点
	q.x = append(q.x,Pair{0,0})

	// 四个方向
	dx:=[]int{-1,0,1,0}
	dy:=[]int{0,1,0,-1}

	for len(q.x) > 0  {
		t := q.front()
		q.pop_front()
		// 已经走过或者个位数之和等于threshold
		if getSum(t) > threshold || st[t.X][t.Y]{
			continue
		}
		res++
		st[t.X][t.Y] = true
		// 遍历四个方向
		for i:=0;i<4 ;i++  {
			x := t.X+dx[i]
			y := t.Y+dy[i]
			if x>=0 && x<rows && y>=0 && y<cols{
				q.x = append(q.x,Pair{x,y})
			}
		}
	}
	return res
}

// 队列
type queue struct{
	x []Pair
}

// 尾插入
func (this *queue) push_back(x Pair){
	this.x = append(this.x,x)
}

// 返回头元素
func (this *queue) front() Pair{
	return this.x[0]
}
// size
func (this *queue) size() int{
	return len(this.x)
}
// 头删除
func (this *queue) pop_front(){
	this.x= this.x[1:]
}