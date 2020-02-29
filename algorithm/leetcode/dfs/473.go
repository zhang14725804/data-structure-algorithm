/*
	473. Matchsticks to Square（经典的剪枝优化问题）

	难！！！

	依次构造正方形的每条边
	（1）从大到小枚举所有边
	（2）每条边内部的木棒长度规定从大到小
	（3）若当前木棒拼接失败，则跳过接下来所有长度相同的木棒
	（4）若当前木棒拼接失败，且是当前边的第一个，则直接剪掉当前分支
	（5）若当前木棒拼接失败，且是当前边的最后一个，则直接剪掉当前分支
*/
func makesquare(nums []int) bool {
    
}