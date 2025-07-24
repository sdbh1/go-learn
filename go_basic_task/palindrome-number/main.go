package main

import (
	"fmt"
)

func main() {
	//编写测试用例
	fmt.Printf("回文数：https://leetcode.cn/problems/palindrome-number/\n")
	fmt.Printf("测试用例1：%v,原始值：%v\n", isPalindrome(121), 121)
	fmt.Printf("测试用例2：%v,原始值：%v\n", isPalindrome(-121), -121)
	fmt.Printf("测试用例3：%v,原始值：%v\n", isPalindrome(10), 10)
	fmt.Printf("测试用例4：%v,原始值：%v\n", isPalindrome(1221), 1221)
	fmt.Printf("测试用例5：%v,原始值：%v\n", isPalindrome(12321), 12321)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x == 0 {
		return true
	}

	digit := 1
	res := x
	for {
		res = res / 10
		if res == 0 {
			break
		}
		digit++
	}
	for i := 1; i <= digit; i++ {
		if i >= digit-i+1 {
			return true
		}
		left := GetValueByIndex(i, x)
		right := GetValueByIndex(digit-i+1, x)
		if left != right {
			return false
		}

	}
	return true
}

func GetValueByIndex(targetDigit, value int) int {
	res := value
	for i := 1; i < targetDigit; i++ {
		res = res / 10
	}
	return res % 10
}
