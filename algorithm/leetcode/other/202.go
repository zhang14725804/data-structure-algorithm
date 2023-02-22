/*
	解法1：
	如果过程中得到了 1 直接返回 true 。
	什么时候得不到 1 呢？产生了循环，也就是出现的数字在之前出现过，那么 1 一定不会得到了，此时返回 false。
	在代码中，我们只需要用 HashSet 去记录已经得到的数字即可
*/
func isHappy(n int) bool {
	set := make(map[int]int, 0)
	set[n] = 1

	for {
		next := getNext(n)
		if next == 1 {
			return true
		}
		if set[next] == 1 {
			return false
		} else {
			set[next] = 1
			n = next
		}
	}
}

func getNext(n int) int {
	next := 0
	for n > 0 {
		// 取余
		t := n % 10
		next += t * t
		// 取整
		n /= 10
	}
	return next
}

/*
	解法2：（妙啊，question）
	这道题，其实本质上就是判断链表是否有环，当出现重复的数字也就是产生了环
*/
func isHappy(n int) bool {
	slow, fast := n, n
	// do while
	for {
		slow = getNext(slow)
		fast = getNext(getNext(fast))
		if slow == fast {
			break
		}
	}
	return slow == 1
}