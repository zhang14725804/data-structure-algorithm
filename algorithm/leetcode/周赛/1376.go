/*
	图论问题：题意都不好懂😅
	todo：求最长路径问题（dfs、bfs）
*/
var son [][]int
func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	son = make([][]int,n)
	for i := 0; i < n; i++ {
		son[i]=make([]int,n)
	}

	for i := 0; i < n; i++ {
		if i!=headID{
			son[manager[i]] = append(son[manager[i]],i)
		}
	}
	return dfs(headID,informTime)
}
// todo:goroutine stack exceeds 1000000000-byte limit
func dfs(u int,informTime []int) int{
	res:=0
	for _, s := range son[u] {
		res = max(res,dfs(s,informTime))
	}
	return res+informTime[u]
}
func max(x, y int) int{
	if x > y{
		return x
	}
	return y
}