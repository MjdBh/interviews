package main

import (
	"container/list"
	"fmt"
	"mjdbh.net/solutions/leetcode"
)

func main() {
	//runParentheses()
}

func runParentheses() {
	parentheses := list.New()
	leetcode.GenerateParentheses("", 0, 0, 4, parentheses)
	for e := parentheses.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value) // print out the elements
	}
}
