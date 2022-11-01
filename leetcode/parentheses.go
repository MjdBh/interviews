package leetcode

import (
	"container/list"
	"fmt"
)

//Leetcode 22. Generate Parentheses

func main() {
	parentheses := list.New()
	generateParentheses("", 0, 0, 4, parentheses)
	for e := parentheses.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value) // print out the elements
	}
}

func generateParentheses(input string, left int, right int, length int, result *list.List) {
	fmt.Println("solve called for : ", input, " left : ", left, " right : ", right)
	if len(input) == 2*length {
		result.PushBack(input)
	}
	if left < length {
		generateParentheses(input+"(", left+1, right, length, result)
	}

	if right < left {
		generateParentheses(input+")", left, right+1, length, result)
	}
	return
}
