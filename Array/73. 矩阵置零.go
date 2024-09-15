package array

/*
给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。
输入：matrix = [[1,1,1],[1,0,1],[1,1,1]]
输出：[[1,0,1],[0,0,0],[1,0,1]]
*/

// 模拟做法，使用标记数组，第一次遍历时将
func setZeroes(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	//第一次遍历使用标记数组将0所在的行和列全部设为true
	rows, cols := make([]bool, len(matrix)), make([]bool, len(matrix[0]))
	for i, v := range matrix {
		for j, val := range v {
			if val == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}
	// 二次遍历，将上边标记数组中标记为true的内容全部替换为0即可
	for i, v := range matrix {
		for j := range v {
			if rows[i] || cols[j] {
				v[j] = 0
			}
		}
	}
	return matrix
}
