package question

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func NumEnclaves(grid [][]int) int {
	ans := 0
	rowCount, columnCount := len(grid), len(grid[0])
	markMap := make([][]bool, rowCount)
	for i := range markMap {
		markMap[i] = make([]bool, columnCount)
	}
	var dfs func(x, y int)
	dfs = func(x, y int) {
		// 边界问题+非陆地+避免重复执行
		if x < 0 || y < 0 || x >= rowCount || y >= columnCount || grid[x][y] == 0 || markMap[x][y] == true {
			return
		}
		markMap[x][y] = true
		for _, dir := range dirs {
			dfs(x+dir.x, y+dir.y)
		}
	}
	// 四条边
	for i := range grid {
		dfs(i, 0)
		dfs(i, columnCount-1) // 列尾
	}
	// 两头已经被处理过了
	for j := 1; j < columnCount-1; j++ {
		dfs(0, j)
		dfs(rowCount-1, j) // 行位
	}
	// 遍历非边界，去掉标记过的
	for i := 1; i < rowCount-1; i++ {
		for j := 1; j < columnCount-1; j++ {
			if grid[i][j] == 1 && !markMap[i][j] {
				ans++
			}
		}
	}
	return ans
}

/**
1020. 飞地的数量
https://leetcode-cn.com/problems/number-of-enclaves/

给你一个大小为 m x n 的二进制矩阵 grid ，其中 0 表示一个海洋单元格、1 表示一个陆地单元格。

一次 移动 是指从一个陆地单元格走到另一个相邻（上、下、左、右）的陆地单元格或跨过 grid 的边界。

返回网格中 无法 在任意次数的移动中离开网格边界的陆地单元格的数量。



示例 1：


输入：grid = [[0,0,0,0],[1,0,1,0],[0,1,1,0],[0,0,0,0]]
输出：3
解释：有三个 1 被 0 包围。一个 1 没有被包围，因为它在边界上。
示例 2：


输入：grid = [[0,1,1,0],[0,0,1,0],[0,0,1,0],[0,0,0,0]]
输出：0
解释：所有 1 都在边界上或可以到达边界。


提示：

m == grid.length
n == grid[i].length
1 <= m, n <= 500
grid[i][j] 的值为 0 或 1
*/
