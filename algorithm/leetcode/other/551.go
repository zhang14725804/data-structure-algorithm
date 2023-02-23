/*
	1. 从前向后遍历
		遇到A，count++，判断是否大于2，满足条件return false
		遇到L，for下一个是否是L，count++，判断是否大于等于3，满足条件return false
		否则return true
*/