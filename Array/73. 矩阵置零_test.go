package array

import "testing"

func TestSetZeroes(t *testing.T) {
	// [[1,1,1],[1,0,1],[1,1,1]]
	input := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	// [[1,0,1],[0,0,0],[1,0,1]]
	exepect := [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}
	result := setZeroes(input)
	for i, v := range result {
		for j, val := range v {
			if val != exepect[i][j] {
				t.Fatalf("failed, input:%v, output:%v, expect:%v", input, result, exepect)
			}
		}
	}
	t.Log("pass")
}
