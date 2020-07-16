/*

	将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

	比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：

	L   C   I   R
	E T O E S I I G
	E   D   H   N
	之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。

	请你实现这个将字符串进行指定行数变换的函数：

	string convert(string s, int numRows);

	思路：找规律。最后一行和第一行是 公差为2n-2的等差数列，中间行是奇偶间隔的等差数列，公差也是2n-2
*/
func convert(s string, n int) string {
	res:=""
	if n == 1{
		return s
	}
	for i := 0; i < n; i++ {
		// 第一行和最后一行
		if i==0 || i == n-1{
			for j:=i;j<len(s);j += 2*n-2{
				// s[i]是byte类型
				res+=string(s[j])
			}
		}else{
			// 中间行，奇偶交替的等差数列
			for j,k:=i,2*n-2-i;j<len(s)||k<len(s);j,k = j+2*n-2,k + 2*n-2{
				// 注意顺序
				if j<len(s){
					res += string(s[j])
				}
				if k<len(s){
					res += string(s[k])
				}
			}
		}	
	}
	return res
}