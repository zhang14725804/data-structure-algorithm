 

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

/* 普通的顺序查找 */
func mainNormal() {
	path := "./mock.txt"
	LINE := countLine(path)
	allData := make([]account, LINE, LINE)
	// 载入数据
	file, _ := os.Open(path) // 打开文件
	i := 0
	buf := bufio.NewReader(file) // 读取文件对象
	for {
		line, _, end := buf.ReadLine()
		if end == io.EOF { // 文件结束跳出循环
			break
		}
		lineStr := string(line)
		lines := strings.Split(lineStr, " | ")
		if len(lines) == 4 {
			allData[i].user = lines[0]
			allData[i].md5 = lines[1]
			allData[i].email = lines[2]
			allData[i].password = lines[3]
		}
		fmt.Println(allData[i])
		i++
	}
	fmt.Println("载入完成")
	// 查询数据
	for {
		fmt.Println("请输入要查询的用户名")
		var target string
		fmt.Scanln(&target) // 处理用户输入
		start := time.Now()
		for i := 0; i < LINE; i++ {
			if allData[i].user == target {
				fmt.Println(allData[i])
			}
		}
		fmt.Println("查询用时：", time.Since(start))
	}
}

//定义数据结构体
type account struct {
	user     string
	md5      string
	email    string
	password string
}

// 统计文件行数
func countLine(path string) int {
	/*
		todo:统计行数，end==io.EOF，end返回invalid argument
		mock-qq.go构造的数据mock-qq-password.txt有问题
	*/
	file, _ := os.Open(path)     // 打开文件
	i := 0                       // 统计行数
	buf := bufio.NewReader(file) // 读取文件对象

	for {
		_, _, end := buf.ReadLine()
		if end == io.EOF { // 文件结束跳出循环
			break
		}
		i++
	}
	fmt.Println("多少行：", i)
	return i
}