/*
	从上往下打印出二叉树的每个结点，同一层的结点按照从左到右的顺序打印
	思路：宽搜（维护一个queue队列）
*/
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
func printFromTopToBottom(root *TreeNode) []int {
	res:=make([]int,0)
	if root == nil{
		return res
	}
	q := &queue{}
	q.push(root)
	for q.size() > 0{
		t := q.front()
		q.pop()
		res = append(res,t.Val)
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