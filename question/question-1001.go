package question

// 抄的 https://leetcode-cn.com/problems/grid-illumination/solution/rustgojava-ha-xi-biao-rust-100-by-kyushu-735o/
// 抄的题解，还没看懂
var DIRS = [][]int{{0, 0}, {1, 1}, {0, 1}, {-1, -1}, {0, -1}, {1, -1}, {1, 0}, {-1, 0}, {-1, 1}}

func gridIllumination(n int, lamps [][]int, queries [][]int) []int {
	axisX, axisY, axisXToY, axisYToX := map[int]int{}, map[int]int{}, map[int]int{}, map[int]int{}
	isBright, ret := map[int]bool{}, make([]int, len(queries))

	for _, lamp := range lamps {
		x, y := lamp[0], lamp[1]
		if isBright[x*n+y] {
			continue
		}
		axisX[x], axisY[y], axisXToY[x-y], axisYToX[x+y], isBright[x*n+y] = axisX[x]+1, axisY[y]+1, axisXToY[x-y]+1, axisYToX[x+y]+1, true
	}
	for i := 0; i < len(queries); i++ {
		x, y := queries[i][0], queries[i][1]
		if axisX[x] > 0 || axisY[y] > 0 || axisXToY[x-y] > 0 || axisYToX[x+y] > 0 {
			ret[i] = 1
		}
		for _, dir := range DIRS {
			tx, ty := x+dir[0], y+dir[1]
			if tx < 0 || tx >= n || ty < 0 || ty >= n || !isBright[tx*n+ty] {
				continue
			}
			delete(isBright, tx*n+ty)
			axisX[tx], axisY[ty], axisXToY[tx-ty], axisYToX[tx+ty] = axisX[tx]-1, axisY[ty]-1, axisXToY[tx-ty]-1, axisYToX[tx+ty]-1
		}
	}
	return ret
}

/**
1001. 网格照明
https://leetcode-cn.com/problems/grid-illumination/
在大小为 n x n 的网格 grid 上，每个单元格都有一盏灯，最初灯都处于 关闭 状态。

给你一个由灯的位置组成的二维数组 lamps ，其中 lamps[i] = [rowi, coli] 表示 打开 位于 grid[rowi][coli] 的灯。即便同一盏灯可能在 lamps 中多次列出，不会影响这盏灯处于 打开 状态。

当一盏灯处于打开状态，它将会照亮 自身所在单元格 以及同一 行 、同一 列 和两条 对角线 上的 所有其他单元格 。

另给你一个二维数组 queries ，其中 queries[j] = [rowj, colj] 。对于第 j 个查询，如果单元格 [rowj, colj] 是被照亮的，则查询结果为 1 ，否则为 0 。在第 j 次查询之后 [按照查询的顺序] ，关闭 位于单元格 grid[rowj][colj] 上及相邻 8 个方向上（与单元格 grid[rowi][coli] 共享角或边）的任何灯。

返回一个整数数组 ans 作为答案， ans[j] 应等于第 j 次查询 queries[j] 的结果，1 表示照亮，0 表示未照亮。



示例 1：


输入：n = 5, lamps = [[0,0],[4,4]], queries = [[1,1],[1,0]]
输出：[1,0]
解释：最初所有灯都是关闭的。在执行查询之前，打开位于 [0, 0] 和 [4, 4] 的灯。第 0 次查询检查 grid[1][1] 是否被照亮（蓝色方框）。该单元格被照亮，所以 ans[0] = 1 。然后，关闭红色方框中的所有灯。

第 1 次查询检查 grid[1][0] 是否被照亮（蓝色方框）。该单元格没有被照亮，所以 ans[1] = 0 。然后，关闭红色矩形中的所有灯。

示例 2：

输入：n = 5, lamps = [[0,0],[4,4]], queries = [[1,1],[1,1]]
输出：[1,1]
示例 3：

输入：n = 5, lamps = [[0,0],[0,4]], queries = [[0,4],[0,1],[1,4]]
输出：[1,1,0]


提示：

1 <= n <= 109
0 <= lamps.length <= 20000
0 <= queries.length <= 20000
lamps[i].length == 2
0 <= rowi, coli < n
queries[j].length == 2
0 <= rowj, colj < n
*/
