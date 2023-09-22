package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	file := os.Args[1]
	key := os.Args[2]
	dat, _ := os.ReadFile(string(file))
	v := bytes.Split(dat, []byte{'\n'})
	for _, l := range v {
		if bytes.Contains(l, []byte(key)) {
			fmt.Println(string(l))
		}
	}
}
