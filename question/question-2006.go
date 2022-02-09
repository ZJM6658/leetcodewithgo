package question

func countKDifference(nums []int, k int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if abs(nums[i]-nums[j]) == k {
				ans++
			}
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func countKDifference2(nums []int, k int) int {
	ans := 0
	// |nums[i] - nums[j]| = k  => nums[i] = nums[j] + k or nums[i] = nums[j] - k
	// eg 1 - 0 = 1 or |1 - 2| = 1 只需要累加0和2出现的次数
	freqMap := map[int]int{}
	for _, num := range nums {
		// i < j 所以遍历一遍即可，后面的统计前面出现的
		ans += freqMap[num+k] + freqMap[num-k]
		// 统计每个数出现的频率
		freqMap[num]++
	}
	return ans
}

/*
2006. 差的绝对值为 K 的数对数目
给你一个整数数组 nums 和一个整数 k ，请你返回数对 (i, j) 的数目，满足 i < j 且 |nums[i] - nums[j]| == k 。

|x| 的值定义为：

如果 x >= 0 ，那么值为 x 。
如果 x < 0 ，那么值为 -x 。


示例 1：

输入：nums = [1,2,2,1], k = 1
输出：4
解释：差的绝对值为 1 的数对为：
- [1,2,2,1]
- [1,2,2,1]
- [1,2,2,1]
- [1,2,2,1]
示例 2：

输入：nums = [1,3], k = 3
输出：0
解释：没有任何数对差的绝对值为 3 。
示例 3：

输入：nums = [3,2,1,5,4], k = 2
输出：3
解释：差的绝对值为 2 的数对为：
- [3,2,1,5,4]
- [3,2,1,5,4]
- [3,2,1,5,4]


提示：

1 <= nums.length <= 200
1 <= nums[i] <= 100
1 <= k <= 99

*/
