/*
	90. Subsets II

	todos::dfs递归出问题，导致程序有问题
*/
var ans [][]int
var path []int
var n int
func subsetsWithDup(nums []int) [][]int {
    // 排序
	for i:=0;i<n;i++{
		for j:=0;j<n;j++{
			if nums[j]<nums[i]{
				nums[i],nums[j] = nums[j],nums[i]
			}
		}
	}
	dfs(nums,0)
	return ans
}
func dfs(nums []int,u int){
	if u == n{
		ans = append(ans, path)
		return 
	}
	// 计算当前数字的个数
	k:=0
	for u+k<n && nums[u+k]==nums[u]{
		k=k+1
	}
	for i:=0;i<=k;i++{
		dfs(nums,u+k)
		path = append(path,nums[u])
	}
	// 恢复现场
	path = make([]int)
}