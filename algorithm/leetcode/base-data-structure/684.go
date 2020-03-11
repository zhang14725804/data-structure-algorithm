/*
	684. Redundant Connection（树中找环）

	并查集思路：所有已经联通的点加到集合

	todos:这个不好理解
	[[1,2],[1,3],[2,3]] 测试通过
	[[1,2],[2,3],[3,4],[1,4],[1,5]]测试不通过

	提到个新算法：强连通分量算法（也叫Tarjan算法）（分为：有向图，无向图）
*/
var p []int
func find(x int) int{
	if p[x] != x{
		p[x] = find(p[x])
	}
	return p[x]
}
func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	for i:=0;i<=n;i++{
		p = append(p,i)
	}
	for _,e := range edges{
		a,b:=e[0],e[1]
		if find(a) == find(b){
			return []int{a,b}
		}
		p[find(a)] = find(b)
	}
	return []int{-1,-1}
}