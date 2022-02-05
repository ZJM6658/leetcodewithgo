package question

import "fmt"

/**
思路：两个数组合并后判断长度，奇位取中位数 偶位取两位平均
*/
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	longArr := []int{}
	l1, l2 := len(nums1), len(nums2)
	i, j := 0, 0
	for i < l1 && j < l2 {
		x, y := nums1[i], nums2[j]
		if x < y {
			longArr = append(longArr, x)
			i++
		} else {
			longArr = append(longArr, y)
			j++
		}
	}
	for i < l1 {
		longArr = append(longArr, nums1[i])
		i++
	}
	for j < l2 {
		longArr = append(longArr, nums2[j])
		j++
	}
	fmt.Println(longArr)
	midIndex := len(longArr) / 2
	if len(longArr)%2 == 0 {
		sum := longArr[midIndex-1] + longArr[midIndex]
		return float64(sum) / 2.0
	}
	return float64(longArr[midIndex])
}

/**
4. 寻找两个正序数组的中位数
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

算法的时间复杂度应该为 O(log (m+n)) 。



示例 1：

输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2
示例 2：

输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
示例 3：

输入：nums1 = [0,0], nums2 = [0,0]
输出：0.00000
示例 4：

输入：nums1 = [], nums2 = [1]
输出：1.00000
示例 5：

输入：nums1 = [2], nums2 = []
输出：2.00000


提示：

nums1.length == m
nums2.length == n
0 <= m <= 1000
0 <= n <= 1000
1 <= m + n <= 2000
-106 <= nums1[i], nums2[i] <= 106
*/
