package question

import "math"

var stateTable = make(map[string][]string, 4)

type AutoMachine struct {
	sign  int
	sum   int
	state string
}

func (am *AutoMachine) get(c rune) {
	am.state = stateTable[am.state][am.getColumn(c)]
	// 符合 和 数字处理
	if am.state == "signed" {
		if c == '+' {
			am.sign = 1
		} else {
			am.sign = -1
		}
	} else if am.state == "in_number" {
		am.sum = am.sum*10 + int(c-'0')
		// 边界 都是min 负数需要反一下比较 正数赋值
		if am.sign == 1 {
			if am.sum > math.MaxInt32 {
				am.sum = math.MaxInt32
			}
		} else {
			if am.sum > -1*math.MinInt32 {
				am.sum = -1 * math.MinInt32
			}
		}
	}
}

func (*AutoMachine) getColumn(c rune) int {
	if c == ' ' {
		return 0
	} else if c == '+' || c == '-' {
		return 1
	} else if c >= '0' && c <= '9' {
		return 2
	}
	return 3
}

// 自动状态机
func MyAtoiStateMachine(s string) int {
	// ' '    +/-   0-9   s
	stateTable["start"] = []string{"start", "signed", "in_number", "end"}
	stateTable["signed"] = []string{"end", "end", "in_number", "end"}
	stateTable["in_number"] = []string{"end", "end", "in_number", "end"}
	stateTable["end"] = []string{"end", "end", "end", "end"}

	am := AutoMachine{
		state: "start",
		sign:  1,
		sum:   0,
	}
	for _, char := range s {
		am.get(char)
	}
	return am.sign * am.sum
}

// 简单实现
func myAtoiSimple(s string) int {
	l, sign, sum := len(s), 1, 0
	i := 0
	// 空格跳过
	for i < l && s[i] == ' ' {
		i++
	}
	// 符号确定
	if i < l {
		if s[i] == '+' {
			sign = 1
			i++
		} else if s[i] == '-' {
			sign = -1
			i++
		}
	}

	for i < l && s[i] >= '0' && s[i] <= '9' {
		// ascii '0' = 48; '1' = 49;
		sum = sum*10 + int(s[i]-'0')
		if sum*sign < math.MinInt32 {
			return math.MinInt32
		} else if sum*sign > math.MaxInt32 {
			return math.MaxInt32
		}
		i++
	}
	return sign * sum
}

/**
8. 字符串转换整数 (atoi)
请你来实现一个 myAtoi(string s) 函数，使其能将字符串转换成一个 32 位有符号整数（类似 C/C++ 中的 atoi 函数）。

函数 myAtoi(string s) 的算法如下：

读入字符串并丢弃无用的前导空格
检查下一个字符（假设还未到字符末尾）为正还是负号，读取该字符（如果有）。 确定最终结果是负数还是正数。 如果两者都不存在，则假定结果为正。
读入下一个字符，直到到达下一个非数字字符或到达输入的结尾。字符串的其余部分将被忽略。
将前面步骤读入的这些数字转换为整数（即，"123" -> 123， "0032" -> 32）。如果没有读入数字，则整数为 0 。必要时更改符号（从步骤 2 开始）。
如果整数数超过 32 位有符号整数范围 [−231,  231 − 1] ，需要截断这个整数，使其保持在这个范围内。具体来说，小于 −231 的整数应该被固定为 −231 ，大于 231 − 1 的整数应该被固定为 231 − 1 。
返回整数作为最终结果。
注意：

本题中的空白字符只包括空格字符 ' ' 。
除前导空格或数字后的其余字符串外，请勿忽略 任何其他字符。


示例 1：

输入：s = "42"
输出：42
解释：加粗的字符串为已经读入的字符，插入符号是当前读取的字符。
第 1 步："42"（当前没有读入字符，因为没有前导空格）
         ^
第 2 步："42"（当前没有读入字符，因为这里不存在 '-' 或者 '+'）
         ^
第 3 步："42"（读入 "42"）
           ^
解析得到整数 42 。
由于 "42" 在范围 [-231, 231 - 1] 内，最终结果为 42 。
示例 2：

输入：s = "   -42"
输出：-42
解释：
第 1 步："   -42"（读入前导空格，但忽视掉）
            ^
第 2 步："   -42"（读入 '-' 字符，所以结果应该是负数）
             ^
第 3 步："   -42"（读入 "42"）
               ^
解析得到整数 -42 。
由于 "-42" 在范围 [-231, 231 - 1] 内，最终结果为 -42 。
示例 3：

输入：s = "4193 with words"
输出：4193
解释：
第 1 步："4193 with words"（当前没有读入字符，因为没有前导空格）
         ^
第 2 步："4193 with words"（当前没有读入字符，因为这里不存在 '-' 或者 '+'）
         ^
第 3 步："4193 with words"（读入 "4193"；由于下一个字符不是一个数字，所以读入停止）
             ^
解析得到整数 4193 。
由于 "4193" 在范围 [-231, 231 - 1] 内，最终结果为 4193 。
示例 4：

输入：s = "words and 987"
输出：0
解释：
第 1 步："words and 987"（当前没有读入字符，因为没有前导空格）
         ^
第 2 步："words and 987"（当前没有读入字符，因为这里不存在 '-' 或者 '+'）
         ^
第 3 步："words and 987"（由于当前字符 'w' 不是一个数字，所以读入停止）
         ^
解析得到整数 0 ，因为没有读入任何数字。
由于 0 在范围 [-231, 231 - 1] 内，最终结果为 0 。
示例 5：

输入：s = "-91283472332"
输出：-2147483648
解释：
第 1 步："-91283472332"（当前没有读入字符，因为没有前导空格）
         ^
第 2 步："-91283472332"（读入 '-' 字符，所以结果应该是负数）
          ^
第 3 步："-91283472332"（读入 "91283472332"）
                     ^
解析得到整数 -91283472332 。
由于 -91283472332 小于范围 [-231, 231 - 1] 的下界，最终结果被截断为 -231 = -2147483648 。


提示：

0 <= s.length <= 200
s 由英文字母（大写和小写）、数字（0-9）、' '、'+'、'-' 和 '.' 组成
*/
