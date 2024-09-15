package array

// 使用模拟的做法，用矩阵表示遍历的记过
func spiralOrder(matrix [][]int) []int {
	// 如果矩阵为[]，直接返回
	n := len(matrix)
	if n == 0 {
		return nil
	}
	// 确定矩阵的行和列
	rows, cols := len(matrix), len(matrix[0])
	// visited表示矩阵中的每个元素是否遍历
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}
	var (
		// 总共元素个数，也是遍历的个数
		total = rows * cols
		// 返回的结果集
		order = make([]int, total)
		// 遍历时候的索引
		row, col = 0, 0
		// 遍历行走的方向，第一个元素表示行，第二个表示列，0表示当前行数不变，只变化列，1表示列右移一位，-1表示左移一位
		// [0, 1]表示第一行往右移动，行数不变，列是增加的，[0,-1]表示从右往左走，行数不变，列减少
		direction = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		// 用于标识是否需要切换方向
		directionIndex = 0
	)
	for i := 0; i < total; i++ {
		// 获取row和col对应的元素填充到结果中
		order[i] = matrix[row][col]
		// 标记当前元素已经记录，用于下边是否访问过的判断
		visited[row][col] = true
		// 通过direction和directionIndex获取下一个行和列的位置
		nextRow, nextCol := row+direction[directionIndex][0], col+direction[directionIndex][1]
		// 五种情况需要改变方向
		// extRow < 0 =>说明到达行最左端
		//nextRow >= rows  =>到达行的最右端
		// nextCol < 0 => 到达列的最上边
		// nextCol >= cols =>到达列的最下边
		// visited[nextRow][nextCol]  =>下一次遍历开始的判断
		if nextRow < 0 || nextRow >= rows || nextCol < 0 || nextCol >= cols || visited[nextRow][nextCol] {
			directionIndex = (directionIndex + 1) % 4
		}
		// row和col需要对应修改
		row += direction[directionIndex][0]
		col += direction[directionIndex][1]
	}

	return order
}

// 下一种做法，按照层级遍历，空间复杂度比第一种优异
func spiralOrderByLayer(matrix [][]int) []int {
	res := []int{}
	rows, cols := len(matrix), len(matrix[0])
	if rows == 0 || cols == 0 {
		return res
	}
	top, bottom, left, right := 0, rows-1, 0, cols-1
	for top <= bottom && left <= right {
		// 层级遍历上右下左
		// 上部遍历时，top是固定的，需要变化列值来获取当前的数值
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		// 第一行遍历完成后，需要top增加1，往第二层做遍历，需要最后与bottom比较，是终止条件的判断参数
		top++
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		// 此处更新了top和right，需要判断是否终止
		if top > bottom || left > right {
			return res
		}
		for i := right; i >= left; i-- {
			res = append(res, matrix[bottom][i])
		}
		bottom--
		for i := bottom; i >= top; i-- {
			res = append(res, matrix[i][left])
		}
		left++
	}
	return res
}
