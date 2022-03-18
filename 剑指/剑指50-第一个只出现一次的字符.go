package main

/*
示例 1:

输入：s = "abaccdeff"
输出：'b

*/

func firstUniqChar(s string) byte {
	if len(s) == 0 {
		return ' '
	}
	recordMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		recordMap[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		if recordMap[s[i]] == 1 {
			return s[i]
		}
	}
	return ' '
}
