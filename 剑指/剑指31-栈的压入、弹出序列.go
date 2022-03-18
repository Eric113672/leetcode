/*
* @Author: lishuang
* @Date:   2022/3/16 11:15
 */

package main

/*
输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈的弹出顺序。
假设压入栈的所有数字均不相等。例如，序列 {1,2,3,4,5} 是某栈的压栈序列，序列 {4,5,3,2,1} 是该压栈序列对应的一个弹出序列，但 {4,3,5,1,2} 就不可能是该压栈序列的弹出序列。

示例 1：

输入：pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
输出：true
解释：我们可以按以下顺序执行：
push(1), push(2), push(3), push(4), pop() -> 4,
push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1

思路：
借助辅助栈 stack, 弹出序列的索引为 i。遍历 pushed并入栈，如果 stack 元素==pushed元素，则执行出栈与 i++
*/

func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0)
	var idx int

	for _, v := range pushed {
		stack = append(stack, v)

		for len(stack) > 0 && stack[len(stack)-1] == popped[idx] {
			stack = stack[:len(stack)-1] // poped the last element
			idx++
		}
	}

	return len(stack) == 0

}
