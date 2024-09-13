package linklist

import "testing"

func TestReverseList(t *testing.T) {
	cases := []struct {
		Name   string
		Input  *ListNode
		Output []int
		Expect []int
	}{
		{
			Name:   "reverse",
			Input:  &ListNode{Val: 1, Next: &ListNode{Val: 2}},
			Output: []int{},
			Expect: []int{2, 1},
		},
		{
			Name:   "zero",
			Input:  nil,
			Output: nil,
			Expect: nil,
		},
	}
	for _, c := range cases {
		res := reverseList(c.Input)
		for res != nil {
			c.Output = append(c.Output, res.Val)
			res = res.Next
		}
		for i, v := range c.Output {
			if v != c.Expect[i] {
				t.Errorf("error reverse, input:%v, output:%v, expect:%v", c.Input, c.Output, c.Expect)
				return
			}
		}
	}
	t.Log("pass")
}
