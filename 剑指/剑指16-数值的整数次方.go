/*
* @Author: lishuang
* @Date:   2022/3/11 21:30
 */

package main

/*
实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn）。不得使用库函数，同时不需要考虑大数问题。
输入：x = 2.00000, n = 10
输出：1024.00000
方法一：分治+二分法
分治拆解过程：
也是没懂。。
x^n --> 2^10: 2^5 --> 2^2 * 2
pow(x, n): subproblem: pow(x, n/2)
merge: if n%2==1奇数: result=subresultsubresultx else偶数: result=subresult*subresult
时间复杂度： O(logN) 由于每次递归都会使得指数减少一半，因此递归的层数为log(n)
空间复杂度： O(logN) 递归的层数，递归函数调用使用的栈空间

*/

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	} else if n < 0 {
		return 1.0 / myPow(x, -n)
	}
	temp := myPow(x, n/2)
	if n%2 == 0 {
		return temp * temp
	}
	return x * temp * temp
}
