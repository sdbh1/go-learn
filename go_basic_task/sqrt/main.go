package main

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

const eps float64 = 1e-15

func Sqrt(x float64) float64 {
	if x <= 0 {
		return math.NaN()
	}
	if x == 0 {
		return 0
	}
	var result float64 = 1.0
	var itCount int32 = 0

	for math.Abs(result*result-x) > eps {
		result = (result + x/result) / 2
		itCount++
	}

	fmt.Printf("Sqrt print,result =%v itCount = %v \n", result, itCount)
	return result
}

// 求平方根
func main() {
	fmt.Printf("\n\n\n")
	fmt.Printf("这个算法的数学背景是牛顿迭代法\n")
	fmt.Printf("牛顿迭代法的公式是：x = (x + a/x) / 2\n")
	fmt.Printf("这个算法的收敛速度是非常快的，只需要很少的迭代次数就可以得到一个非常精确的结果\n")
	fmt.Printf("https://blog.csdn.net/SanyHo/article/details/106365314")
	fmt.Printf("\n\n\n")
	Sqrt(-1)
	Sqrt(0)
	Sqrt(2)
	Sqrt(4)
	Sqrt(16)
	Sqrt(25)
	Sqrt(100)
}
