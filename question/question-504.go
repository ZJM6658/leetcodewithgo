package question

func ConvertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	// 负数
	isMinus := num < 0
	if isMinus {
		num = -num
	}
	ans := []byte{}
	for num != 0 {
		ans = append(ans, '0'+byte(num%7))
		num /= 7
	}
	if isMinus {
		ans = append(ans, '-')
	}

	for i, l := 0, len(ans); i < l/2; i++ {
		ans[i], ans[l-i-1] = ans[l-i-1], ans[i]
	}
	return string(ans)
}

/*
504. 七进制数
给定一个整数 num，将其转化为 7 进制，并以字符串形式输出。



示例 1:

输入: num = 100
输出: "202"
示例 2:

输入: num = -7
输出: "-10"


提示：

-107 <= num <= 107
*/
