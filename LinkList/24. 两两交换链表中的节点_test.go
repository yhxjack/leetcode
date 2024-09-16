package linklist

import "testing"

func TestSwapPairs(t *testing.T) {
	cases := []struct {
		// [1,2,3,4]
		Input  *ListNode
		Output []int
		Expect []int
	}{
		{
			Input:  &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}},
			Output: []int{},
			Expect: []int{2, 1, 4, 3},
		},
		{
			Input:  nil,
			Output: []int{},
			Expect: []int{},
		},
		{
			Input:  &ListNode{Val: 1},
			Output: []int{},
			Expect: []int{1},
		},
	}
	for _, c := range cases {
		// result := swapPairs(c.Input)
		// if result != nil {
		// 	for result != nil {
		// 		c.Output = append(c.Output, result.Val)
		// 		result = result.Next
		// 	}
		// }
		// for i, v := range c.Output {
		// 	if v != c.Expect[i] {
		// 		t.Fatalf("failed, input:%v, output:%v, expect:%v", c.Input, c.Output, c.Expect)
		// 	}
		// }
		
		// recursion
		resultByRecursion := swapPairsByRecursion(c.Input)
		if resultByRecursion != nil {
			for resultByRecursion != nil {
				c.Output = append(c.Output, resultByRecursion.Val)
				resultByRecursion = resultByRecursion.Next
			}
		}
		for i, v := range c.Output {
			if v != c.Expect[i] {
				t.Fatalf("failed, input:%v, output:%v, expect:%v", c.Input, c.Output, c.Expect)
			}
		}
	}
	t.Log("pass")
}
