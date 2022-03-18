/*
* @Author: lishuang
* @Date:   2022/3/18 11:21
 */

package main

/*
请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。
若队列为空，pop_front 和 max_value需要返回 -1

示例 1：

输入:
["MaxQueue","push_back","push_back","max_value","pop_front","max_value"]
[[],[1],[2],[],[],[]]
输出:[null,null,null,2,1,2]
思路：双端队列
从队列尾部插入元素时，我们可以提前取出队列中所有比这个元素小的元素，使得队列中只保留对结果有影响的数字。
这样的方法等价于要求维持队列单调递减，即要保证每个元素的前面都没有比它小的元素。

*/

type MaxQueue struct {
	q []int
	p []int
}

func Constructor3() MaxQueue {
	return MaxQueue{}
}

func (this *MaxQueue) Max_value() int {
	if len(this.q) == 0 {
		return -1
	}
	return this.p[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.q = append(this.q, value)
	for len(this.p) > 0 && value > this.p[len(this.p)-1] {
		this.p = this.p[:len(this.p)-1]
	}
	this.p = append(this.p, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.q) == 0 {
		return -1
	}
	if this.p[0] == this.q[0] {
		this.p = this.p[1:]
	}
	value := this.q[0]
	this.q = this.q[1:]
	return value
}
