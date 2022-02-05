package question

// 学习动态规划解法
func longestPalindrome(s string) string {
	len := len(s)
	if len < 2 {
		return s
	}
	maxLen, begin := 1, 0

	// 这里要用到变量 所有使用slice
	dp := make([][]bool, len)
	for i := range dp {
		dp[i] = make([]bool, len)
	}

	for i := 0; i < len; i++ {
		dp[i][i] = true
	}
	// 左下角先填
	for j := 1; j < len; j++ {
		for i := 0; i < j; i++ {
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+maxLen]
}

/**
5. 最长回文子串
给你一个字符串 s，找到 s 中最长的回文子串。



示例 1：

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
示例 2：

输入：s = "cbbd"
输出："bb"
示例 3：

输入：s = "a"
输出："a"
示例 4：

输入：s = "ac"
输出："a"


提示：

1 <= s.length <= 1000
s 仅由数字和英文字母（大写和/或小写）组成
*/
