package base

func Mono(nums []int) bool {
	size := len(nums)

	if size < 3 {
		return true
	}
	order := 0
	for i := 1; i < size; i++ {
		if order == 0 {
			order = nums[i] - nums[i-1]
		}
		if nums[i]-nums[i-1] > 0 && order < 0 || nums[i]-nums[i-1] < 0 && order > 0 {
			return false
		}
	}
	return true
}
