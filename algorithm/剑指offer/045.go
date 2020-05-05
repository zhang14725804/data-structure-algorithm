/*
	之字形打印二叉树
*/
func printFromTopToBottom(root *TreeNode) [][]int {
	res:=make([][]int,0)
	if root == nil{
		return res
	}

	q := &queue{}
	q.push(root)
	// 结束标识符
	q.push(nil)

	level:= make([]int,0)
	// 奇数行偶数行
	flag:=false
	for q.size()>0 {
		t := q.front()
		q.pop()
		// nil标识一行的结尾
		if t == nil{
			// 此行为空
			if len(level) == 0{
				break
			}
			// 反转
			if flag{
				level = reverse(level)
			}
			res = append(res,level)
			// 清空level
			level = make([]int,0)
			q.push(nil)
			flag = !flag
			continue
		}
		level = append(level,t.Val)
		if t.Left != nil{
			q.push(t.Left)
		}
		if t.Right != nil{
			q.push(t.Right)
		}
	}
	return res
}
// 反转数组int
func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// 队列
type queue struct{
	x []*TreeNode
}

// 尾插入
func (this *queue) push(x *TreeNode){
	this.x = append(this.x,x)
}

// 返回头元素
func (this *queue) front() *TreeNode{
	return this.x[0]
}
// 头删除
func (this *queue) pop(){
	this.x= this.x[1:]
}
// size
func (this *queue) size() int{
	return len(this.x)
}