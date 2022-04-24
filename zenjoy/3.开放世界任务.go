/*
* @Author: lishuang
* @Date:   2022/4/22 20:59
 */

package main

/*
思路解析， 遍历每一个当前任务是否可以完成 可以完成则进行递归 relateTask也去掉当前这个任务 当任务数为0 时 便是所有可以完成 反之
*/

func checkCanFinishTask(relatedTasks [][]int) bool {
	var res bool
	record := make(map[int]bool)
	backTrack(relatedTasks, record, 0, &res)
	return res

}

func backTrack(relateTask [][]int, record map[int]bool, startIndex int, res *bool) {
	if len(relateTask) == 0 {
		*res = true
		return
	}
	for i := startIndex; i < len(relateTask); i++ {
		if check(record, relateTask[i]) {
			record[i] = true
			curr := relateTask[i]
			relateTask = append(relateTask[:i], relateTask[i+1:]...)
			backTrack(relateTask, record, i+1, res)
			relateTask = append(relateTask[:i], append([][]int{curr}, relateTask[i:]...)...)
		}
		continue
	}
}

func check(record map[int]bool, currTask []int) bool {
	for i := 0; i < len(currTask); i++ {
		if res, ok := record[currTask[i]]; !res || !ok {
			return false
		}
	}
	return true
}
