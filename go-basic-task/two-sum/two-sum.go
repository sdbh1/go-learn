package main

func main() {

}

func twoSumBySelf(nums []int, target int) []int {
	var result []int = make([]int, 2)
	for i := 0; i < len(nums); i++ {
		findNum := target - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == findNum {
				result[0] = i
				result[1] = j
				return result
			}
		}
	}
	return result
}

// 官方题解
func twoSumByLeetCode(nums []int, target int) []int {
	for i, x := range nums {
		for j := i + 1; j < len(nums); j++ {
			if x+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// HashTable
func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}
