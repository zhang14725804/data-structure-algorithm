/*
	深度优先遍历或者广度优先遍历
	我没想到 😅 😅 😅
*/
func getImportance(employees []*Employee, id int) (total int) {
	mp := map[int]*Employee{}
	for _, emploee := range employees {
		mp[emploee.Id] = emploee
	}

	var dfs func(int)
	dfs = func(id int) {
		emploee := mp[id]
		total += emploee.Importance
		for _, subId := range emploee.Subordinates {
			dfs(subId)
		}
	}
	dfs(id)
	return
}

func getImportance(employees []*Employee, id int) (total int) {
	mp := map[int]*Employee{}
	for _, employee := range employees {
		mp[employee.Id] = employee
	}

	// id入栈
	queue := []int{id}
	for len(queue) > 0 {
		// 出栈
		employee := mp[queue[0]]
		queue = queue[1:]
		total += employee.Importance
		// 下属入栈
		for _, subId := range employee.Subordinates {
			queue = append(queue, subId)
		}
	}
	return
}