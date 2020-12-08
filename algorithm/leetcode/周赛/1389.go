/*
	1389，看不懂
*/
func createTargetArray(nums []int, index []int) []int {
	target := make([]int,0,0)
	for i:=0;i<len(nums);i++{
		// todo:看不懂
		target = append(target,0)
		fmt.Println("第",i+1,target,len(target)-1,index[i])
		for j:=len(target) - 1;j>index[i];j--{
			target[j] = target[j-1]
		}
		fmt.Println(target)
		target[index[i]] = nums[i]
	}
	return target
}