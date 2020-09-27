/*
	生命游戏（使用原地算法解决本题）（todo）

	num代表当前为止周围的活细胞数量
	num <2 ,死
	num ==2 || num == 3 ,不变
	num > 3 ，死
	num == 3 活

	八个位置怎么表示
*/
func gameOfLife(board [][]int) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	n, m := len(board), len(board[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			//
			live := 0
			// 枚举周围八个格子
			for x := MaxInt(0, i-1); x <= MinInt(n-1, i+1); x++ {
				for y := MaxInt(0, j-1); y <= MinInt(m-1, j+1); y++ {
					//
					if (x != i || y != j) && (board[x][y]&1 == 1) {
						live++
					}
				}
			}
			// 当前状态
			cur := board[i][j] & 1
			// 下一时刻的状态
			next := 0
			if cur == 1 {
				if live < 2 || live > 3 {
					next = 0
				} else {
					next = 1
				}
			} else {
				if live == 3 {
					next = 1
				} else {
					next = 0
				}
			}
			// 结果存到十位
			board[i][j] |= next << 1
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			// 把十位的数移到个位
			board[i][j] >>= 1
		}
	}
}
