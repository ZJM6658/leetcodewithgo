package question

import (
	"math"
	"strconv"
)

// 内存2.5 击败78%
func CalPoints(ops []string) int {
	ans, validScoreList := 0, []int{}
	for i := 0; i < len(ops); i++ {
		ls := len(validScoreList)
		op := ops[i]
		score := math.MinInt64
		if op == "C" {
			// 取消成绩后，成绩往前移
			ans -= validScoreList[ls-1]
			validScoreList = validScoreList[:ls-1]
		} else if op == "D" {
			score = validScoreList[ls-1] * 2
		} else if op == "+" {
			score = validScoreList[ls-1] + validScoreList[ls-2]
		} else {
			score, _ = strconv.Atoi(op)
		}
		if score != math.MinInt64 {
			ans += score
			validScoreList = append(validScoreList, score)
		}
	}
	return ans
}

// 宫水三叶的  内存2.4 击败100%
func CalPoints3Leaf(ops []string) int {
	ans, validScoreList := 0, make([]int, 1010)
	idx := 0
	for i := 0; i < len(ops); i++ {
		idx++
		op := ops[i]
		if op == "C" {
			// 取消成绩后，成绩往前移
			idx -= 2
		} else if op == "D" {
			validScoreList[idx] = validScoreList[idx-1] * 2
		} else if op == "+" {
			validScoreList[idx] = validScoreList[idx-1] + validScoreList[idx-2]
		} else {
			validScoreList[idx], _ = strconv.Atoi(op)
		}
	}

	for i := 0; i < idx; i++ {
		ans += validScoreList[i]
	}

	return ans
}

/*
682. 棒球比赛
你现在是一场采用特殊赛制棒球比赛的记录员。这场比赛由若干回合组成，过去几回合的得分可能会影响以后几回合的得分。

比赛开始时，记录是空白的。你会得到一个记录操作的字符串列表 ops，其中 ops[i] 是你需要记录的第 i 项操作，ops 遵循下述规则：

整数 x - 表示本回合新获得分数 x
"+" - 表示本回合新获得的得分是前两次得分的总和。题目数据保证记录此操作时前面总是存在两个有效的分数。
"D" - 表示本回合新获得的得分是前一次得分的两倍。题目数据保证记录此操作时前面总是存在一个有效的分数。
"C" - 表示前一次得分无效，将其从记录中移除。题目数据保证记录此操作时前面总是存在一个有效的分数。
请你返回记录中所有得分的总和。



示例 1：

输入：ops = ["5","2","C","D","+"]
输出：30
解释：
"5" - 记录加 5 ，记录现在是 [5]
"2" - 记录加 2 ，记录现在是 [5, 2]
"C" - 使前一次得分的记录无效并将其移除，记录现在是 [5].
"D" - 记录加 2 * 5 = 10 ，记录现在是 [5, 10].
"+" - 记录加 5 + 10 = 15 ，记录现在是 [5, 10, 15].
所有得分的总和 5 + 10 + 15 = 30
示例 2：

输入：ops = ["5","-2","4","C","D","9","+","+"]
输出：27
解释：
"5" - 记录加 5 ，记录现在是 [5]
"-2" - 记录加 -2 ，记录现在是 [5, -2]
"4" - 记录加 4 ，记录现在是 [5, -2, 4]
"C" - 使前一次得分的记录无效并将其移除，记录现在是 [5, -2]
"D" - 记录加 2 * -2 = -4 ，记录现在是 [5, -2, -4]
"9" - 记录加 9 ，记录现在是 [5, -2, -4, 9]
"+" - 记录加 -4 + 9 = 5 ，记录现在是 [5, -2, -4, 9, 5]
"+" - 记录加 9 + 5 = 14 ，记录现在是 [5, -2, -4, 9, 5, 14]
所有得分的总和 5 + -2 + -4 + 9 + 5 + 14 = 27
示例 3：

输入：ops = ["1"]
输出：1


提示：

1 <= ops.length <= 1000
ops[i] 为 "C"、"D"、"+"，或者一个表示整数的字符串。整数范围是 [-3 * 104, 3 * 104]
对于 "+" 操作，题目数据保证记录此操作时前面总是存在两个有效的分数
对于 "C" 和 "D" 操作，题目数据保证记录此操作时前面总是存在一个有效的分数
*/
