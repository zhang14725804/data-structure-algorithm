/*
	用栈，我以为是双指针 😅😅
*/
func asteroidCollision(asteroids []int) (st []int) {
	for _, aster := range asteroids {
		alive := true
		// 😅😅 当前aster < 0，栈顶元素 > 0
		for alive && aster < 0 && len(st) > 0 && st[len(st)-1] > 0 {
			alive = st[len(st)-1] < -aster
			// 碰撞且嗝屁，出栈
			if st[len(st)-1] <= -aster {
				st = st[:len(st)-1]
			}
		}
		// 入栈
		if alive {
			st = append(st, aster)
		}
	}
	return
}