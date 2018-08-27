package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("You must specify one arg as a file path")
		os.Exit(1)
	}
	arg := os.Args[1]
	proverbs, err := ioutil.ReadFile(arg)
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
