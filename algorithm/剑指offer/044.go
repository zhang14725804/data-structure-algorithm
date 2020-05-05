/*
	从上到下按层打印二叉树，同一层的结点按从左到右的顺序打印，每一层打印到一行
	宽搜，nil标识一行的结尾（思路不错）
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

	for q.size()>0 {
		t := q.front()
		q.pop()
		// nil标识一行的结尾
		if t == nil{
			// 此行为空
			if len(level) == 0{
				break
			}
			res = append(res,level)
			// 清空level
			level = make([]int,0)
			q.push(nil)
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