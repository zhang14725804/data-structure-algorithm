/*
	47. Permutations II

	全排列问题，两种方法（这特么。。。）
	（1）枚举每个位置上放那个树
	（2）枚举每个数放在哪个位置
*/
/*
	方法二：枚举每个数放在哪个位置（这个题需要判重）

	将所有相同的数放在一起 sort
	认为规定相同数字的相对顺序：不变
	dfs增加状态

	todos::dfs递归出问题，导致程序有问题
*/ 
var (
	n int // 数组大小
	st []bool 
	ans [][]int // 所有方法
	path []int // 当前方案
)
func permuteUnique(nums []int) [][]int {
	n = len(nums)
	st = make([]bool,n)
	path = make([]int,n)
	// 排序
	for i:=0;i<n;i++{
		for j:=0;j<n;j++{
			if nums[j]<nums[i]{
				nums[i],nums[j] = nums[j],nums[i]
			}
		}
	}
	dfs(nums,0,0)
	return ans
}

func dfs(nums []int,u,start int){
	if u==n{
		ans = append(ans,path)
		return
	}
	for i:=start;i<n;i++{
		if !st[i]{
			st[i] = true
			path[i] = nums[u]
			var s int
			if u+1<n && nums[u+1]==nums[u]{
				s=i+1
			}else{
				s=0
			}
            dfs(nums,u+1,s)
			st[i] = false
		}
	}
}