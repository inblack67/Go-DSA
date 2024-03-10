package main

func TwoSum(nums []int, target int) []int {
	var res []int

	hash := make(map[int]int)

	for i, v := range nums {
		diff := target - v
		index, exists := hash[diff]
		if exists {
			res = append(res, index)
			res = append(res, i)
			return res
		}
		hash[v] = i
	}

	return res
}
