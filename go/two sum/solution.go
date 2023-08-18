func twoSum(nums []int, target int) []int {
	numIndices := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if index, found := numIndices[complement]; found {
			return []int{index, i}
		}
		numIndices[num] = i
	}

	return nil
}