package question

import (
	"fmt"
	"strconv"
	"strings"
)

func complexNumberMultiply(num1 string, num2 string) string {
	num1 = strings.ReplaceAll(num1, "i", "")
	num2 = strings.ReplaceAll(num2, "i", "")
	s1 := strings.Split(num1, "+")
	s2 := strings.Split(num2, "+")
	a, _ := strconv.Atoi(s1[0])
	b, _ := strconv.Atoi(s1[1])
	c, _ := strconv.Atoi(s2[0])
	d, _ := strconv.Atoi(s2[1])
	A, B := a*c-b*d, b*c+a*d
	//return string(A) + string(B) + "i"
	return fmt.Sprintf("%d%di", A, B)
}

/*
537. 复数乘法
复数 可以用字符串表示，遵循 "实部+虚部i" 的形式，并满足下述条件：

实部 是一个整数，取值范围是 [-100, 100]
虚部 也是一个整数，取值范围是 [-100, 100]
i2 == -1
给你两个字符串表示的复数 num1 和 num2 ，请你遵循复数表示形式，返回表示它们乘积的字符串。



示例 1：

输入：num1 = "1+1i", num2 = "1+1i"
输出："0+2i"
解释：(1 + i) * (1 + i) = 1 + i2 + 2 * i = 2i ，你需要将它转换为 0+2i 的形式。
示例 2：

输入：num1 = "1+-1i", num2 = "1+-1i"
输出："0+-2i"
解释：(1 - i) * (1 - i) = 1 + i2 - 2 * i = -2i ，你需要将它转换为 0+-2i 的形式。


提示：

num1 和 num2 都是有效的复数表示。
*/
