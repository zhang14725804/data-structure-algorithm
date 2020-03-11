/*
	547. Friend Circles
	并查集（图的深度\宽度优先遍历）

	todos：：[[1,0,0,1],[0,1,1,0],[0,1,1,1],[1,0,1,1]] 测试不通过
	[[1,1,0],[1,1,0],[0,0,1]] 可以通过
*/

var p []int

/*
	找到父节点（每个集合的代表元素）
	路径压缩的过程
*/ 
func find(x int) int{
	if p[x] != x {
		p[x] = find(p[x])
	}
	return p[x]
}

func findCircleNum(M [][]int) int {
	n := len(M)
	for i:=0;i<n;i++{
		p = append(p,i)
	}
	res:=n
	for i:=0;i<n;i++{
		for j:=0;j<i;j++{
			if M[i][j] == 0 {
				continue
			}
			if find(i) != find(j){
				p[find(i)] = find(j)
				res -= 1
			}
		}
	}
	return res
}

/*
	并查集的概念：：

	（1）合并两个集合
	（2）判断两个点是否在同一个集合中

	优化：路径压缩，按秩合并
*/