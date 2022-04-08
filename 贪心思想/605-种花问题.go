/*
* @Author: lishuang
* @Date:   2022/4/8 11:10
 */

package main

import "fmt"

/*
假设有一个很长的花坛，一部分地块种植了花，另一部分却没有。可是，花不能种植在相邻的地块上，它们会争夺水源，两者都会死去。

给你一个整数数组flowerbed 表示花坛，由若干 0 和 1 组成，其中 0 表示没种植花，1 表示种植了花。
另有一个数n ，能否在不打破种植规则的情况下种入n朵花？能则返回 true ，不能则返回 false。


示例 1：

输入：flowerbed = [1,0,0,0,1], n = 1
输出：true

大概思路就是统计该值是0 然后pre是0 next也是0 则cnt +1

[1,0,0,0,0,0,1]
*/

func canPlaceFlowers(flowerbed []int, n int) bool {
	cnt := 0
	for i := 0; i < len(flowerbed); {
		if flowerbed[i] == 0 && ((i-1) < 0 || flowerbed[i-1] == 0) && ((i+1) >= len(flowerbed) || flowerbed[i+1] == 0) {
			//fmt.Println(i, cnt)
			cnt++
			i += 2
			continue
		}
		//fmt.Println("CanPlace", i, cnt)
		i++
	}
	return cnt >= n
}

func main() {
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 0, 0, 1}, 2))

}
