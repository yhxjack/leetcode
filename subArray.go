package main

import "fmt"

//func main() {
//	list := []int{1, 2, 3, 4, 5, 7, 9}
//	k := 9
//	count := 0
//	//[2,3,4], [4, 5], [9]
//	res := [][]int{}
//	for i := 0; i < len(list); i++ {
//		sum := 0
//		//这里有个细节，这里可以将列表分为i之前的子表，进行反向遍历，查找是否存在和为k的子数组
//		for j := i; j >= 0; j-- {
//			sum += list[j]
//			if sum == k {
//				count++
//				res = append(res, list[j:i+1])
//			}
//		}
//	}
//	fmt.Println("count", count)
//}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	k := 5
	res := subarraySum(nums, k)
	fmt.Println(res)
}
func subarraySum(nums []int, k int) int {
	// 使用前缀和优化
	pre, count := 0, 0
	st := map[int]int{}
	st[0] = 1
	for _, v := range nums {
		pre += v
		if p, ok := st[k-pre]; ok {
			count += p
		}
		st[pre]++
	}
	return count
}
