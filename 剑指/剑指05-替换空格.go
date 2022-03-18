package main

/*
输入：s = "We are happy."
输出："We%20are%20happy."
*/

func replaceSpace(s string) string {
	var res []byte
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			res = append(res, s[i])
		} else {
			res = append(res, []byte("%20")...)
		}
	}
	return string(res)
}

func main() {

}
