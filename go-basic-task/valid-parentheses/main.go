package main

import (
	"fmt"
)

func main() {
	fmt.Printf("20. 有效的括号：https://leetcode.cn/problems/valid-parentheses/description/\n")
	var s string = "([)]"
	fmt.Printf("isValid print,test Value = %v,result =%v \n", s, isValid(s))
	s = "()"
	fmt.Printf("isValid print,test Value = %v,result =%v \n", s, isValid(s))
	s = "()[]{}"
	fmt.Printf("isValid print,test Value = %v,result =%v \n", s, isValid(s))
}

var KeySign = []byte{'(', '[', '{'}
var ValueSign = []byte{')', ']', '}'}

type Stack []byte

func (s *Stack) Push(v byte) {
	*s = append(*s, v)
}

func (s *Stack) Pop() byte {
	if len(*s) == 0 {
		return 0
	}

	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func (s *Stack) Peek() byte {
	if len(*s) == 0 {
		return 0
	}
	v := (*s)[len(*s)-1]
	return v
}

func isValid(s string) bool {
	var stack Stack        // 声明一个 Stack 类型的变量
	stack = make(Stack, 0) // 初始化栈
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if CheckIsKeySign(ch) {
			stack.Push(ch)
		}

		if CheckIsValueSign(ch) {
			var keySign byte = stack.Peek()
			if IsRightString(keySign, ch) {
				stack.Pop()
			} else {
				return false
			}

		}
	}
	return len(stack) == 0
}

func CheckIsKeySign(ch byte) bool {
	for _, v := range KeySign {
		if v == ch {
			return true
		}
	}
	return false
}

func CheckIsValueSign(ch byte) bool {
	for _, v := range ValueSign {
		if v == ch {
			return true
		}
	}
	return false
}

func IsRightString(lastSign byte, checkSign byte) bool {
	switch checkSign {
	case ')':
		return lastSign == '('
	case ']':
		return lastSign == '['
	case '}':
		return lastSign == '{'
	default:
		return false
	}
}
