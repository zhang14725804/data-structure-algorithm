## 深度优先遍历（DFS（Depth-First-Search））和广度优先遍历（BFS（Breadth-First Search））


### BFS

广度优先遍历时需要一个队列，还需要存储所有方案；适用于**最短最小**之类的问题。

BFS 找到的路径一定是最短的，但代价就是空间复杂度比 DFS 大很多


```golang
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
			// 将0和相邻的数字交换位置
			for _, adj := range neighbor[idx] {
				newBoard := []rune(cur)
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
```

### 双向BFS

传统的 BFS 框架就是从起点开始向四周扩散，遇到终点时停止；而双向 BFS 则是从起点和终点同时开始扩散，当两边有交集的时候停止

双向 BFS 也有局限，因为你必须知道终点在哪里

双向 BFS 还是遵循 BFS 算法框架的，只是不再使用队列，而是使用 HashSet 方便快速判断两个集合是否有交集。