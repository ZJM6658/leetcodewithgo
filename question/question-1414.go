package question

import "fmt"

func FindMinFibonacciNumbers(k int) int {
	var numbers []int

	visibleFibonacciNumbers := fetchFibonacciNum(k)
	fmt.Printf("查找后得到：%d\n", visibleFibonacciNumbers)
	i := len(visibleFibonacciNumbers) - 1
	for k > 0 {
		a := visibleFibonacciNumbers[i]
		if k >= a {
			numbers = append(numbers, a)
			k = k - a
		}
		i -= 1
	}
	fmt.Printf("数组：%v", numbers)
	return len(numbers)
}

func fetchFibonacciNum(k int) []int {
	resultSlice := []int{1, 1}
	var f1, f2 = 1, 1

	for true {
		var b = f2
		f2 = f1 + f2
		f1 = b

		if f2 > k {
			break
		}
		resultSlice = append(resultSlice, f2)
	}
	return resultSlice
}

/*
1414. 和为 K 的最少斐波那契数字数目
给你数字 k ，请你返回和为 k 的斐波那契数字的最少数目，其中，每个斐波那契数字都可以被使用多次。

斐波那契数字定义为：

F1 = 1
F2 = 1
Fn = Fn-1 + Fn-2 ， 其中 n > 2 。
数据保证对于给定的 k ，一定能找到可行解。



示例 1：

输入：k = 7
输出：2
解释：斐波那契数字为：1，1，2，3，5，8，13，……
对于 k = 7 ，我们可以得到 2 + 5 = 7 。
示例 2：

输入：k = 10
输出：2
解释：对于 k = 10 ，我们可以得到 2 + 8 = 10 。
示例 3：

输入：k = 19
输出：3
解释：对于 k = 19 ，我们可以得到 1 + 5 + 13 = 19 。

*/
