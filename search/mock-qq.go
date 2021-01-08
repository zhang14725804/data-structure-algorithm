 

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

//生成count个[start,end)结束的不重复的随机数
func GenerateRandomNumber(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}

	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start
		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}
		if !exist {
			nums = append(nums, num)
		}
	}
	return nums
}

/*
	mock一些数据，写入文件
	269279025----269279025i
	数字太大不行
*/
func main() {
	const LEN = 666666
	nums := GenerateRandomNumber(1000000, 999999999, LEN)
	letter := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	// 统计计算时间
	t1 := time.Now()
	// 新建文件并写入
	path := "./mock-qq-password.txt"
	saveFile, _ := os.Create(path)
	defer saveFile.Close()
	save := bufio.NewWriter(saveFile)
	for i := 0; i < LEN; i++ {
		// 写入文件
		fmt.Fprintln(save, fmt.Sprintf("%d", nums[i])+"----"+fmt.Sprintf("%d", nums[i])+string(letter[rand.Intn(10)]))
	}
	save.Flush() // 刷新
	fmt.Println(time.Since(t1))
}
