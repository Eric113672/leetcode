/*
* @Author: lishuang
* @Date:   2022/3/21 15:50
 */

package main

/*
输入一个字符串，打印出该字符串中字符的所有排列。

你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。

示例:

输入：s = "abc"
输出：["abc","acb","bac","bca","cab","cba"]

*/
// 思路解析： 采取hash表计数,回溯递归

func permutation(s string) []string {
	res := []string{} //返回值列表
	Hmap := make(map[byte]int)
	ele := "" //待构造的字符串
	for i := 0; i < len(s); i++ {
		Hmap[s[i]]++ //计数
	}
	var dfs func(start int)
	dfs = func(start int) {
		if start == len(s) {
			res = append(res, ele) //字符串构造完毕 添加进返回值列表
			return
		}
		for k, _ := range Hmap {
			if Hmap[k] != 0 { //次数不为0说明可用
				ele += string(k)
				Hmap[k]--
				dfs(start + 1)         //从新的点继续构造字符串
				ele = ele[:len(ele)-1] //回溯
				Hmap[k]++
			}
		}
	}
	dfs(0)
	return res
}
