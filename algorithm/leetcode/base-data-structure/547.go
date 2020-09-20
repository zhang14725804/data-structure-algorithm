/*
	547. Friend Circles
	1 <= N <= 200
	M[i][i] == 1
	M[i][j] == M[j][i]

	并查集（图的深度\宽度优先遍历）
*/

var p []int

/*
	找到父节点（每个集合的代表元素）
	路径压缩的过程

	todo：[[1,0,0,1],[0,1,1,0],[0,1,1,1],[1,0,1,1]] 测试未通过
*/
func find(x int) int {
	if p[x] != x {
		p[x] = find(p[x])
	}
	return p[x]
}

func findCircleNum(M [][]int) int {
	n := len(M)
	for i := 0; i < n; i++ {
		p = append(p, i)
	}
	res := n
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if M[i][j] == 0 {
				continue
			}
			if find(i) != find(j) {
				p[find(i)] = find(j)
				res -= 1
			}
		}
	}
	return res
}

/*
	思路：dfs

	选定一个节点，开始深度优先搜索，将遍历到的节点标记为 visited，直到遍历结束，连通图数目加一
	选取另外一个未遍历的节点，重复上述过程
	直到所有节点都被遍历

*/
var visited []bool // M[i][j] == M[j][i]，所以一维数组搞定
var l int
var M [][]int

// 目标数组长宽相等，所以循环用一个边界条件（l）
func findCircleNum(_M [][]int) int {
	M = _M
	res := 0
	l = len(M)
	visited = make([]bool, l)
	// 每行
	for i := 0; i < l; i++ {
		if !visited[i] {
			visited[i] = true
			dfs(i)
			res++ // M[i][i] == 1
		}
	}
	return res
}
func dfs(i int) {
	// 每列
	for j := 0; j < l; j++ {
		if i != j && !visited[j] && M[i][j] == 1 {
			visited[j] = true
			dfs(j)
		}
	}
}