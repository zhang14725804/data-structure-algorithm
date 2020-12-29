/*
	BFS 最短路径问题
*/
func openLock(deadends []string, target string) int {
	visited := make(map[string]bool, 0) // ps: 标记已经走过的，防止重复走
	for i := 0; i < len(deadends); i++ {
		visited[deadends[i]] = true
	}
	var queue []string
	step := 0
	queue = append(queue, "0000")
	if _, ok := visited["0000"]; ok {
		return -1
	}
	visited["0000"] = true
	for len(queue) > 0 {
		size := len(queue)
		// 将当前队列中的所有节点向周围扩散
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur == target {
				return step
			}
			// 将一个节点的相邻节点加入队列
			for j := 0; j < 4; j++ {
				up := plus(cur, j)
				if _, ok := visited[up]; !ok {
					queue = append(queue, up)
					visited[up] = true
				}
				down := minus(cur, j)
				if _, ok := visited[down]; !ok {
					queue = append(queue, down)
					visited[down] = true
				}
			}
		}
		// 增加步数
		step++
	}
	return -1
}

// 向上向下拨动
func plus(s string, i int) string {
	ch := []rune(s)
	if ch[i] == '9' {
		ch[i] = '0'
	} else {
		ch[i] += 1
	}
	return string(ch)
}
func minus(s string, i int) string {
	ch := []rune(s)
	if ch[i] == '0' {
		ch[i] = '9'
	} else {
		ch[i] -= 1
	}
	return string(ch)
}