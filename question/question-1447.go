package question

import "fmt"

func SimplifiedFractions(n int) []string {
	res := []string{}
	for i := 1; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if gcdSubstract(i, j) == 1 {
				res = append(res, fmt.Sprintf("%d/%d", i, j))
			}
		}
	}
	return res
}

func gcdSubstract(a, b int) int {
	if a == b {
		return a
	}
	if a > b {
		return gcdSubstract(a-b, b)
	}
	return gcdSubstract(a, b-a)
}

func gcdMod(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

/*
1447. 最简分数
给你一个整数 n ，请你返回所有 0 到 1 之间（不包括 0 和 1）满足分母小于等于  n 的 最简 分数 。分数可以以 任意 顺序返回。



示例 1：

输入：n = 2
输出：["1/2"]
解释："1/2" 是唯一一个分母小于等于 2 的最简分数。
示例 2：

输入：n = 3
输出：["1/2","1/3","2/3"]
示例 3：

输入：n = 4
输出：["1/2","1/3","1/4","2/3","3/4"]
解释："2/4" 不是最简分数，因为它可以化简为 "1/2" 。
示例 4：

输入：n = 1
输出：[]


提示：

1 <= n <= 100
*/
