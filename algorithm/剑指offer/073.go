/*
	数组中只出现一次的两个数字
	思路1：for循环+hash（笨办法）
	todo：思路2：知识点：异或运算（不好理解）
*/
func findNumsAppearOnce(nums []int) []int {
	sum := 0 // x^y
	// 所有数异或的结果就是x^y
	for _,x := range nums{
		sum ^= x
	}
	// 😅
	k := 0
	for (sum >> k & 1) == 0{
		k++
	}
	// 😅
	first := 0
	for _,x := range nums{
		if (x >> k & 1) == 1{
			first ^= x
		}
	}
	return []int{first,sum^first}
}