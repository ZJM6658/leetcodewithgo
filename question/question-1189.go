package question

import "math"

func MaxNumberOfBalloons(text string) int {
	ans := math.MaxInt32
	countMap := map[rune]int{
		'b': 0,
		'a': 0,
		'l': 0,
		'o': 0,
		'n': 0,
	}
	for _, char := range text {
		if _, ok := countMap[char]; ok {
			countMap[char]++
		}
	}
	for r, i := range countMap {
		if r == 'l' || r == 'o' {
			if ans > i/2 {
				ans = i / 2
			}
		} else if ans > i {
			ans = i
		}
	}
	return ans
}

/*
1189. “气球” 的最大数量
给你一个字符串 text，你需要使用 text 中的字母来拼凑尽可能多的单词 "balloon"（气球）。

字符串 text 中的每个字母最多只能被使用一次。请你返回最多可以拼凑出多少个单词 "balloon"。



示例 1：



输入：text = "nlaebolko"
输出：1
示例 2：



输入：text = "loonbalxballpoon"
输出：2
示例 3：

输入：text = "leetcode"
输出：0


提示：

1 <= text.length <= 10^4
text 全部由小写英文字母组成
*/
