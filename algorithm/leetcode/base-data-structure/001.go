package leetcode

import "fmt"

var nums = []int{2, 7, 11, 15}
var target int = 9

// Letcode001 第一题
/*
	给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
	你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

	给定 nums = [2, 7, 11, 15], target = 9
	因为 nums[0] + nums[1] = 2 + 7 = 9
	所以返回 [0, 1]

*/
func Letcode001() {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == target-nums[i] {
				fmt.Println(i, j)
				return
			}
		}
	}
	panic("===Letcode001===")
}

// Letcode001Hash 哈希方法
func Letcode001Hash() {
	mp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if index, ok := mp[complement]; ok {
			fmt.Println(index, i)
			return
		}
		mp[nums[i]] = i
	}
	panic("===Letcode001Hash===")
}
