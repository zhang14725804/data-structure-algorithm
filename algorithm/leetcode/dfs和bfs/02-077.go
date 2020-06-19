/*
	todo:测试不通过，有问题
	输入2，4；返回[[1,4],[1,4],[1,4],[2,4],[2,4],[3,4]]
*/
var ans [][]int

func combine(n int, k int) [][]int {
	path := make([]int,0)
	dfs(path, 1, n, k)
	return ans
}
// 
func dfs(path []int, start,n,k int)  {
	if k == 0{
		ans = append(ans,path)
		return
	}
	for i:=start; i<=n; i++{
		path = append(path,i)
		dfs(path, i+1, n, k-1)
		// 
		path = path[:len(path)-1]
	}
}