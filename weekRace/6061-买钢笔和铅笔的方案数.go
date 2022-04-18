/*
* @Author: lishuang
* @Date:   2022/4/16 22:42
 */

package main

import "fmt"

/*
给你一个整数 total ，表示你拥有的总钱数。同时给你两个整数 cost1 和 cost2 ，分别表示一支钢笔和一支铅笔的价格。你可以花费你部分或者全部的钱，去买任意数目的两种笔。

请你返回购买钢笔和铅笔的 不同方案数目 。

示例 1：
输入：total = 20, cost1 = 10, cost2 = 5
输出：9
解释：一支钢笔的价格为 10 ，一支铅笔的价格为 5 。
- 如果你买 0 支钢笔，那么你可以买 0 ，1 ，2 ，3 或者 4 支铅笔。
- 如果你买 1 支钢笔，那么你可以买 0 ，1 或者 2 支铅笔。
- 如果你买 2 支钢笔，那么你没法买任何铅笔。
所以买钢笔和铅笔的总方案数为 5 + 3 + 1 = 9 种。
*/
//超时了。。
/*
func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {
	var allWays int64
	for i := 0; cost1*i <= total; i++ {
		extra := total - cost1*i
		if extra == 0 {
			allWays++
			continue
		}
		for j := 0; j*cost2 <= extra; j++ {
			allWays++
		}
	}
	return allWays
}*/
// 我是垃圾
func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {
	var allWays int64
	for i := 0; cost1*i <= total; i++ {
		extra := total - cost1*i
		fmt.Println("before extra", extra, allWays)
		if extra == 0 {
			allWays++
			continue
		}
		allWays = allWays + int64(extra/cost2+1)
		fmt.Println("after", allWays, int64(extra/cost2+1))
	}
	return allWays
}

func main() {
	fmt.Println(waysToBuyPensPencils(20, 10, 5))
}
