/*
* @Author: lishuang
* @Date:   2022/3/21 16:09
 */

package main

/*
请实现一个函数用来匹配包含'. '和'*'的正则表达式。模式中的字符'.'表示任意一个字符，而'*'表示它前面的字符可以出现任意次（含0次）。
在本题中，匹配是指字符串的所有字符匹配整个模式。例如，字符串"aaa"与模式"a.a"和"ab*ac*a"匹配，但与"aa.a"和"ab*a"均不匹配。

示例 1:
输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。
*/
// 无赖解法
/*

func isMatch(s, p string) bool {
	return regexp.MustCompile("^" + p + "$").MatchString(s)
}
*/

// dp  这可读性太差了
func isMatch(s string, p string) bool {
	l1, l2 := len(s), len(p)
	f := make([][]bool, l1+1)
	for i := range f {
		f[i] = make([]bool, l2+1)
	}

	f[0][0] = true
	for i := 0; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if p[j-1] == '*' {
				if f[i][j-2] {
					f[i][j] = true
				} else if i != 0 && f[i-1][j] && (s[i-1] == p[j-2] || p[j-2] == '.') {
					f[i][j] = true
				}
			} else if i != 0 && f[i-1][j-1] && (s[i-1] == p[j-1] || p[j-1] == '.') {
				f[i][j] = true
			}
		}
	}

	return f[l1][l2]
}
