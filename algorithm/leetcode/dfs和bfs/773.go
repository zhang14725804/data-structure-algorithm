
/*
	迷宫问题转换为BFS问题( question )

	1、一般的 BFS 算法，是从一个起点start开始，向终点target进行寻路，但是拼图问题不是在寻路，而是在不断交换数字，这应该怎么转化成 BFS 算法问题呢？
	2、即便这个问题能够转化成 BFS 问题，如何处理起点start和终点target？它们都是数组哎，把数组放进队列，套 BFS 框架
*/
func slidingPuzzle(board [][]int) int {
	m, n := len(board), len(board[0])
	// 从start到target，将问题转化为BFS
	start := ""
	target := "123450"
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			start += fmt.Sprintf("%v", board[i][j])
		}
	}
	// 记录一维字符串的相邻索引
	neighbor := [][]int{
		{1, 3},
		{0, 4, 2},
		{1, 5},
		{0, 4},
		{3, 1, 5},
		{4, 2},
	}

	// BFS框架
	var queue []string
	// 保存已经访问过的节点
	visited := make(map[string]bool, 0)
	queue = append(queue, start)
	visited[start] = true
	step := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			// 队列，先进先出
			cur := queue[0]
			queue = queue[1:]
			// 到达目标
			if target == cur {
				return step
			}
			idx := 0
			// 找到数字 0 索引
			for cur[idx] != '0' {
				idx++
			}
			// 将0和相邻的数字交换位置
			for _, adj := range neighbor[idx] {
				newBoard := []rune(cur)
				newBoard[adj], newBoard[idx] = newBoard[idx], newBoard[adj]
				sBoard := string(newBoard)
				// 防止重复走
				if _, ok := visited[sBoard]; !ok {
					queue = append(queue, sBoard)
					visited[sBoard] = true
				}
			}
		}
		step++
	}
	return -1
}