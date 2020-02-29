/*
	216. Combination Sum III

	枚举所有选法，判断总和

	todos：：dfs递归出问题，导致程序有问题
*/
var ans [][]int 
var path []int
func combinationSum3(k int, n int) [][]int {
	dfs(k,1,n)
	return ans
}

func dfs(k,start,n int){
	if k==0{
		if n==0{
			ans = append(ans,path)
		}
        return
	}
	for i:=start;i<=10-k;i++{
		path = append(path,i)
		dfs(k-1,i+1,n-1)
		// 
		path = path[:len(path)-1]
	}
}
