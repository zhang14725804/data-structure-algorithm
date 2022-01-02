/*
	给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
*/

/*
	方法1： 
		（1）层序遍历
		（2）根据奇数偶数翻转
	0102  根据奇数偶数选择队头对尾出队，想歪了
*/
func zigzagLevelOrder(root *TreeNode) [][]int {
    var res [][]int
    if root ==nil{
        return res
    }
    var queue []*TreeNode
	queue = append(queue,root)
	// 记录奇数偶数
	flag:=true
	// （1）遍历每一层
    for len(queue)>0{
        cLen:=len(queue)
		var level []int
		// （2）遍历每一层的每个节点
        for i:=0;i<cLen;i++{
            cnode:=queue[0]
            queue = queue[1:]
            level = append(level,cnode.Val)
			// 左右子树
            if cnode.Left!=nil{
                queue = append(queue,cnode.Left)
            }
            if cnode.Right!=nil{
                queue = append(queue,cnode.Right)
            }
		}
		// 偶数层翻转
        if !flag{
            for i, j := 0, len(level)-1; i < j; i, j = i+1, j-1 {
				level[i], level[j] = level[j], level[i]
			}
		}
		// 填充数据
		res = append(res,level)
		// 改变奇偶
        flag = !flag
    }
    return res
}