package main

/*
请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。


示例1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。


func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	maxLength, tempLength := 0, 0

	tempMap := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		fmt.Println("currMap", tempMap, tempLength, maxLength)
		_, ok := tempMap[s[i]]
		if !ok {
			tempMap[s[i]] = i
			tempLength++
		} else {
			if tempLength > maxLength {
				maxLength = tempLength
			}
			tempLength = 1
		}
	}
	if tempLength > maxLength {
		return tempLength
	} else {
		return maxLength
	}
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
*/

// new
/*
解题思路
状态定义 dp[i]表示以i为结尾的最长子字符串的长度.
状态转移方程为 有点类似于剑指 Offer 42. 连续子数组的最大和 。但是当以i为结尾的字符不能从dp[i-1]转移过来的时候,更新方式略有不同。

以"tmmzuxt"为例，但遍历到第二个‘m’的时候。但是我们仔细观察"tmm"，显然以't'开头的子串，包含'm'。那么要把dp[i]更新为1。
以"abba"为例，遍历到第二个'b'时，和上述一致，遍历到第二个'a'的时候，以‘b’为开头的新子串长度为1，表示为dp[i-1] = 1， 两个'a'的相对距离超过了dp[i-1]表示的子串的长度。
意味着前一个'a'和新的子串(以第二个'b'开头的子串)没有关系,虽然当前'a'与前一个'a'冲突，但是对于第二'b'开头的子串来说可以计入答案。可以从dp[i-1]这个状态转移而来。
综上的意思就是当前字符出现在dp[i-1]表示的子串中的时候，转移方程为dp[i] = i - m[s[i]]//记录长度
当前字符没有出现在dp[i-1]表示的子串中 dp[i] = dp[i-1] + 1
淦。。还是双指针写吧

*/

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	dp := make([]int, n+1)
	// for i := range dp {dp[i] = 1}
	m := map[byte]int{}
	ans := 1
	dp[0] = 1
	m[s[0]] = 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 1; i < n; i++ {
		//i-m[s[i]]  > dp[i-1] 判断 s[i]在不在dp[i-1]表示的子串中
		if _, ok := m[s[i]]; !ok || i-m[s[i]] > dp[i-1] {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = i - m[s[i]]
		}
		m[s[i]] = i
		ans = max(ans, dp[i])
	}
	// fmt.Println(dp)
	return ans
}
