package array

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

func subarraySum(nums []int, k int) int {
	// 使用前缀和优化
	count, prev := 0, 0
	n := len(nums)
	if n == 0 {
		return count
	}
	st := map[int]int{}
	st[0] = 1
	for i := 0; i < n; i++ {
		prev += nums[i]
		if p, ok := st[prev-k]; ok {
			count += p
		}
		st[prev]++
	}
	return count
}

func subarraySumByEnum(nums []int, k int) int {
	// 使用枚举做法，暴力解法
	if len(nums) == 0 {
		return 0
	}
	count := 0
	res := [][]int{}
	for i := 0; i < len(nums); i++ {
		sum := 0
		//这里有个细节，这里可以将列表分为i之前的子表，进行反向遍历，查找是否存在和为k的子数组
		for j := i; j >= 0; j-- {
			sum += nums[j]
			if sum == k {
				count++
				res = append(res, nums[j:i+1])
			}
		}
	}
	return count
}
