/*
	todo:这个题不懂
	a:=[]int{1,-1,3,-1}
	b:=[]int{2,-1,-1,-1}
	fmt.Println(validateBinaryTreeNodes(4,a,b))
*/
func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	d := make([]int, n)
	// 填充入度
	for i := 0; i < n; i++ {
		if leftChild[i] != -1 {
			d[leftChild[i]]++
		}
		if rightChild[i] != -1 {
			d[rightChild[i]]++
		}
	}
	fmt.Println("arr:", d)
	// 根节点
	root := 0
	for root < n && d[root] > 0 {
		root++
	}
	fmt.Println("root:", root)
	// 没有根节点
	if root == n {
		return false
	}

	st := make([]bool, n)
	st[root] = true
	// 借助队列
	q := &queue{}
	q.push_back(root)
	// 宽度搜索
	for q.size() > 0 {
		// 对头数据获取删除
		t := q.front()
		q.pop_front()

		sons := []int{leftChild[t], rightChild[t]}
		for _, s := range sons {
			// 有子节点
			if s != -1 {
				if st[s] {
					return false
				}
				// 标记被遍历过
				st[s] = true
				q.push_back(s)
			}
		}
	}
	// 判断所有的点是否都被遍历过
	for _, state := range st {
		if !state {
			return false
		}
	}
	return true
}


