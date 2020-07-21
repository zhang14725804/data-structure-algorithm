package main

/*
	在一排树中，第 i 棵树产生 tree[i] 型的水果。
	你可以从你选择的任何树开始，然后重复执行以下步骤：

	把这棵树上的水果放进你的篮子里。如果你做不到，就停下来。
	移动到当前树右侧的下一棵树。如果右边没有树，就停下来。
	请注意，在选择一颗树后，你没有任何选择：你必须执行步骤 1，然后执行步骤 2，然后返回步骤 1，然后执行步骤 2，依此类推，直至停止。

	你有两个篮子，每个篮子可以携带任何数量的水果，但你希望每个篮子只携带一种类型的水果。
	用这个程序你能收集的水果总量是多少？

	思路：滑动窗口
	[3,3,3,1,2,1,1,2,3,3,4]
*/
func totalFruit(tree []int) int {
	i, j := 0, 0
	hash := make(map[int]int, 0)
	max := 1
	// 遍历树
	for ; i < len(tree); i++ {
		// 当前元素（水果）作为键，当前序号作为value
		hash[tree[i]] = i
		hashSize := 0
		min := len(tree)
		for key, _ := range hash {
			hashSize++
			min = compare(min, hash[key], false)
		}
		// 当前篮子中的水果种类大于2的时候
		if hashSize > 2 {
			// 删除序号较小的那个水果
			delete(hash, tree[min])
			// 移动j到tree[min]的下一位
			j = min + 1
		}
		max = compare(max, i-j+1, true)
	}
	return max
}
