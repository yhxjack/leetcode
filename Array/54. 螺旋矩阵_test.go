package array

import "testing"

func TestSpiralOrder(t *testing.T) {
	cases := []struct {
		name   string
		input  [][]int
		expect []int
	}{
		{name: "normal",
			input:  [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			expect: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
	}
	for _, c := range cases {
		result := spiralOrder(c.input)
		if len(result) != len(c.expect) {
			t.Fatalf("error, result is not expect, input:%v, result:%v, expect:%v", c.input, result, c.expect)
		}
		for i, v := range result {
			if v != c.expect[i] {
				t.Fatalf("failed, result:%v is not equal expect:%v", result, c.expect)
			}
		}
		resultBylayer := spiralOrderByLayer(c.input)
		if len(resultBylayer) != len(c.expect) {
			t.Fatalf("error, result is not expect, input:%v, result:%v, expect:%v", c.input, resultBylayer, c.expect)
		}
		for i, v := range resultBylayer {
			if v != c.expect[i] {
				t.Fatalf("failed, result:%v is not equal expect:%v", resultBylayer, c.expect)
			}
		}
	}
}
