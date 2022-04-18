/*
* @Author: lishuang
* @Date:   2022/4/16 23:05
 */

package main

type ATM struct {
	cashCnt  []int
	cashType []int
}

func Constructor() ATM {
	cashCnt := []int{0, 0, 0, 0, 0}
	cashType := []int{20, 50, 100, 200, 500}
	return ATM{cashCnt, cashType}
}

func (this *ATM) Deposit(banknotesCount []int) {
	for i := range this.cashCnt {
		this.cashCnt[i] += banknotesCount[i]
	}
}

func (this *ATM) Withdraw(amount int) []int {
	res := make([]int, 5)
	for i := 4; i >= 0 && amount > 0; i-- {
		t := this.cashType[i]
		c := this.cashCnt[i]
		if c == 0 {
			continue
		}
		res[i] = min(c, amount/t)
		amount -= t * res[i]
	}
	if amount != 0 {
		return []int{-1}
	}
	for i := range res {
		this.cashCnt[i] -= res[i]
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
 * Your ATM object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Deposit(banknotesCount);
 * param_2 := obj.Withdraw(amount);
 */
