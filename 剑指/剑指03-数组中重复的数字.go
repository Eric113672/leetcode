package main

func findRepeatNumber(nums []int) int {
	numsMap := map[int]bool{}
	for _, num := range nums {
		if _, ok := numsMap[num]; !ok {
			numsMap[num] = true
		} else {
			return num
		}

	}
	return 0
}

func main() {

}
