package array

import "testing"

func TestSubarraySum(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 7, 9}
	k := 9
	// [2,3,4], [4, 5], [9]
	expect := 3
	result := subarraySum(list, k)
	if result != expect {
		t.Fatalf("error,input:%v, k:%d, result:%v, expect:%d", list, k, result, expect)
	} else {
		t.Log("pass")
	}
	result1 := subarraySumByEnum(list, k)
	if result1 != expect {
		t.Fatalf("subarraySumByEnum,input:%v, k:%d, result:%v, expect:%d", list, k, result, expect)
	} else {
		t.Log("pass")
	}
}
