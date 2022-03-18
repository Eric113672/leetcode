package main

func reverseLeftWords(s string, n int) string {
	if n == 0 {
		return s
	}
	var byteList []byte
	for i := 0; i < n; i++ {
		byteList = append(byteList, s[i])
	}
	return s[n:] + string(byteList)
}

func main() {

}
