package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	var wordlist string
	var keyword string
	flag.StringVar(&wordlist, "w", "", "Specify wordlist.")
	flag.StringVar(&keyword, "k", "", "Specify keyword.")
	flag.Parse()
	if wordlist == "" || keyword == "" {
		log.Fatal("Please enter a wordlist(-w) and a keyword(-k).")
	}
	start := time.Now()
	dat, _ := ioutil.ReadFile(string(wordlist))
	sdat := bytes.Split(dat, []byte{'\n'})
	line := 0
	for _, l := range sdat {
		line += 1
		if bytes.Contains(l, []byte(keyword)) {
			fmt.Println(strconv.Itoa(line) + " | " + string(l))
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("\nFinished in %s", elapsed)
}
