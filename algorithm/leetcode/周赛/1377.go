/*
	图论问题：dfs、bfs
	todo：题意如何判断边界，四种边界
	答案有问题
*/
var e [][]int

func frogPosition(n int, edges [][]int, t int, target int) float64 {
	e = make([][]int,n+1)
	for i := 0; i < n+1; i++ {
		e[i]=make([]int,n+1)
	}
	for _, edge := range edges {
		a:=edge[0]
		b:=edge[1]
		e[a] = append(e[a],b)
		e[b] = append(e[b],a)
	}
	return dfs(1,-1,t,target,1)
}

func dfs(u, father, t, target int, p float64) float64 {
	if t==0{
		if u == target{
			return p
		}
		return 0
	}
	k := len(e[u])
	if u != 1{
		k--
	}
	if k == 0{
		if u == target{
			return p
		}
		return 0
	}
	var res float64 = 0
	for _, s := range e[u] {
		if s != father{
			res=max(res, dfs(s, u, t-1, target, p/float64(k)))
		}
	}
	return res
}
func max(x, y float64) float64{
	if x > y{
		return x
	}
	return y
}