package src

func twoSum(nums []int, target int) []int {
	numberIndexMapping := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		indexOfComplement, ok := numberIndexMapping[complement]
		if ok {
			return []int{indexOfComplement, i}
		}
		numberIndexMapping[num] = i
	}
	panic("No solution exists?")
}
