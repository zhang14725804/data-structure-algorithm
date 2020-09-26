/*
	累加数是一个字符串，组成它的数字可以形成累加序列。
	一个有效的累加序列必须至少包含 3 个数。除了最开始的两个数以外，字符串中的其他数都等于它之前两个数相加的和。

	todo:你如何处理一个溢出的过大的整数输入?有没有更好的办法
 */

 // 方法1：回溯，代码有问题，总是返回false
 func isAdditiveNumber(nums string) bool {
    if len(nums)==0{
        return true
    }
	//i 表示第一个数字的结尾(不包括 i)
    for i:=1;i<len(nums);i++{
		// 0 开头, 并且当前数字不是 0
        if nums[0] == '0' && i>1 {
            return false
		}
		 //j 表示从 i 开始第二个数字的结尾(不包括 j)
        for j:=i+1;j<len(nums);j++{
			// 0 开头, 并且当前数字不是 0  
            if nums[i]=='0' && j-i>1{
                break
            }

            num1,_ := strconv.Atoi(nums[0:i])
            num2 ,_:= strconv.Atoi(nums[i:j])
            if helper(nums[j:],num1,num2){
                return true
            }
        }
    }
    return false
}
// 这里有问题（todo）
func helper(nums string,num1,num2 int) bool{
    sum,_ := strconv.Atoi(nums)
    if num1+num2 == sum && nums[0]!='0'{
        return true
    }
    if len(nums) == 0{
        return true
    }
    for i:=1;i<len(nums);i++{
        if nums[0]=='0' && i>1{
            return false
        }
        sum,_ := strconv.Atoi(nums[0:i])
        if num1+num2 == sum{
            return helper(nums[i:],num2,sum)
        }else if num1+num2 < sum{
            break
        }
    }
    return false
}