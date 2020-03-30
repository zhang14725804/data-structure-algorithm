package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
	用户名密码组成的数据,如下所示：
	235345356----!@#$%trewq
*/
type Account struct{
	Password string
	User int
}
/*
	内存搜索，全部读到内存，数据大了岂不是要跪
*/
func mainStruct(){
	const N = 3
	// 初始化数组
	allStrs:=make([]Account,N,N)
	i:=0// 统计多少行
	// 打开文件
	path:="./mock1.txt"
	file,_:=os.Open(path)
	defer file.Close()
	buf:=bufio.NewReader(file)
	// 数据载入内存
	for {
		// 逐行读取
		line,_,end:=buf.ReadLine()
		if end==io.EOF{
			break
		}
		userPassword:=strings.Split(string(line),"----")
		allStrs[i].Password = userPassword[1]
		allStrs[i].User,_ =strconv.Atoi(userPassword[0])
		i++
	}
	// 根据QQ搜索密码
	for ; ;  {
		fmt.Println("请输入要查询的数据")
		var keywords int
		// 获取键盘输入
		fmt.Scanf("%d",&keywords)
		// 内存搜索
		for j:=0;j<N ;j++  {
			if allStrs[j].User == keywords{
				fmt.Println(j,allStrs[j].User,allStrs[j].Password)
			}
		}
	}
}
