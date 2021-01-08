 

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
	数据载入map内存，然后在进行搜索，相比数组存储结构更快
*/
func mainMap() {
	const N = 3
	// 初始化map,如果是map，自己构造的模拟数据键值对重复概率比较大
	allStrs := make(map[int]string, N)
	i := 0 // 统计多少行
	// 打开文件
	path := "./mock1.txt"
	file, _ := os.Open(path)
	defer file.Close()
	buf := bufio.NewReader(file)
	// 数据载入内存
	for {
		// 逐行读取
		line, _, end := buf.ReadLine()
		if end == io.EOF {
			break
		}
		userPassword := strings.Split(string(line), "----")
		Password := userPassword[1]
		User, _ := strconv.Atoi(userPassword[0])
		// 映射账户密码到map
		allStrs[User] = Password
		i++
	}
	// 根据QQ搜索密码
	for {
		fmt.Println("请输入要查询的数据")
		var keywords int
		// 获取键盘输入
		fmt.Scanf("%d", &keywords)
		// 内存搜索
		pwd, err := allStrs[keywords]
		if err {
			fmt.Println(keywords, pwd)
		} else {
			fmt.Println("buncunzai")
		}
	}
}