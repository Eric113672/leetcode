package main

type MinStack struct {
	numStack []int
	minStack []int
}

func (this *MinStack) Push(x int) {
	this.numStack = append(this.numStack, x)
	top := this.minStack[len(this.minStack)-1]
	this.minStack = append(this.minStack, min(top, x))
}

func (this *MinStack) Pop() {
	this.numStack = this.numStack[:len(this.numStack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.numStack[len(this.numStack)-1]
}

func (this *MinStack) Min() int {
	return this.minStack[len(this.minStack)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {

}
