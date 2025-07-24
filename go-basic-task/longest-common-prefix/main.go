package main

import (
	"fmt"
)

func main() {

	fmt.Printf("14. 最长公共前缀https://leetcode.cn/problems/longest-common-prefix/description/\n")
	var test1 []string = []string{"flower", "flow", "flight"}
	result := longestCommonPrefix(test1)
	fmt.Printf("longest-common-prefix,go,result =%v \n", result)
}

func longestCommonPrefix(strs []string) string {
	var result string = ""
	i := 0
	for {
		var ch byte = 0
		for j := 0; j < len(strs); j++ {
			if i >= len(strs[j]) {
				return result
			}

			thisch := strs[j][i]

			if ch != 0 && thisch != ch {
				return result
			}

			ch = strs[j][i]
		}
		if ch != 0 {
			result = result + string(ch)
			ch = 0
		}
		i++
	}
	return result
}
