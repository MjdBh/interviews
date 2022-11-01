package leetcode

import (
	"container/list"
	"fmt"
)

//Leetcode 22. Generate Parentheses

func GenerateParentheses(input string, left int, right int, length int, result *list.List) {
	fmt.Println("solve called for : ", input, " left : ", left, " right : ", right)
	if len(input) == 2*length {
		result.PushBack(input)
	}
	if left < length {
		GenerateParentheses(input+"(", left+1, right, length, result)
	}

	if right < left {
		GenerateParentheses(input+")", left, right+1, length, result)
	}
	return
}
