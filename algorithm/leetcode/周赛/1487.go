package main

import (
	"fmt"
	"strconv"
)

/*
	给你一个长度为 n 的字符串数组 names 。你将会在文件系统中创建 n 个文件夹：
	在第 i 分钟，新建名为 names[i] 的文件夹。

	由于两个文件 不能 共享相同的文件名，因此如果新建文件夹使用的文件名已经被占用，
	系统会以 (k) 的形式为新文件夹的文件名添加后缀，其中 k 是能保证文件名唯一的 最小正整数 。

	返回长度为 n 的字符串数组，其中 ans[i] 是创建第 i 个文件夹时系统分配给该文件夹的实际名称。

	todo:(根据数据量推算时间复杂度空间复杂度)

	1 <= names.length <= 5 * 10^4
	1 <= names[i].length <= 20
	names[i] 由小写英文字母、数字和/或圆括号组成。

*/
func getFolderNames(names []string) []string {
	// 存储name和k的hash
	name_k := make(map[string]int, 0)
	// 存储文件名的set（用hash模拟set）
	set := make(map[string]int, 0)
	res := make([]string, 0)

	for i := 0; i < len(names); i++ {
		end := ""
		k := 0
		// 如果set中存在当前name
		if _, ok := set[names[i]]; ok {
			k = name_k[names[i]]
		}
		if _, ok := set[names[i]+end]; ok {
			k++
			end = "(" + strconv.Itoa(k) + ")"
		}
		name_k[names[i]] = k
		set[names[i]+end] = k
		res = append(res, names[i]+end)
	}
	fmt.Println(res)
	return res
}

func main() {
	// tode：测试不通过！！
	str := []string{"onepiece", "onepiece(1)", "onepiece(2)", "onepiece(3)", "onepiece"}
	getFolderNames(str)
}
