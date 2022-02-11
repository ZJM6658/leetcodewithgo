package main

import (
	"fmt"
	"leetcodewithgo/question"
)

func main() {
	//question.FindMinFibonacciNumbers(19)
	//a := question.TwoSum([]int{2, 7, 8, 11, 15}, 26))
	//a := question.LengthOfLongestSubstring("pwwkeabcpwwwwwwwww")
	//a := question.CountGoodRectangles([][]int{[]int{5, 8}, []int{3, 9}, []int{5, 12}, []int{16, 5}})

	//a := question.FindMedianSortedArrays([]int{1, 2}, []int{3, 4})
	//a := question.FindMedianSortedArrays([]int{}, []int{2})

	//a := question.SumOfUnique([]int{1, 1, 2, 3, 2, 5, 7, 7, 7})
	//a := question.Convert("ZHUJIAMIN123", 3)

	//a := question.SimplifiedFractions(4)

	//a := question.MinimumDifference([]int{9, 4, 1, 7}, 2)
	a := question.MinimumDifference([]int{87063, 61094, 44530, 21297, 95857, 93551, 9918}, 6)

	fmt.Println(a)

}
