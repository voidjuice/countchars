package main

import (
	"fmt"
	"strings"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	proverbs, err := ioutil.ReadFile("/tmp/dat")
	sproverbs := string(proverbs)
	check(err)

  lines := strings.Split(sproverbs, "\n")
  for _, l := range lines {
	  fmt.Printf("%s\n", l)
	  for k, v := range charCount(l) {
		  fmt.Printf("'%c'=%d, ", k, v)
	  }
	  fmt.Print("\n\n")
  }
}
func charCount(line string) map[rune]int {
	m := make(map[rune]int, 0)
	for _, c := range line {
		m[c] = m[c] + 1
	}
	return m
}
