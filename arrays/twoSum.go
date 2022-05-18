package arrays

func TowSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		p := target - nums[i];
		_, ok := m[p]
		if ok {
			return []int{ m[p], i }
		}
		m[nums[i]] = i
	}
	return []int{0, 0}
}