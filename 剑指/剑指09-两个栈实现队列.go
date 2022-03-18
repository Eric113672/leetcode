package main

type CQueue struct {
	numSequence []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.numSequence = append(this.numSequence, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.numSequence) == 0 {
		return -1
	}
	head := this.numSequence[0]
	this.numSequence = this.numSequence[1:]
	return head
}

func main() {

}
