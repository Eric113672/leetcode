/*
* @Author: lishuang
* @Date:   2022/3/6 11:06
 */

package main

import (
	"bytes"
	"fmt"
	"io"
)

const debug = true

func main() {
	var buf io.Writer
	if debug {
		buf = new(bytes.Buffer) // enable collection of output
	}
	f(buf) // NOTE: subtly incorrect!
	fmt.Println(buf)
}

// If out is non-nil, output will be written to it.
func f(out io.Writer) {
	// ...do something...
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}
