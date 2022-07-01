package sol

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxProduct(nums []int) int {
	res := nums[0]
	cur_max, cur_min := 1, 1

	for _, num := range nums {
		cur_max, cur_min = max(num, max(cur_max*num, cur_min*num)),
			min(num, min(cur_max*num, cur_min*num))
		res = max(res, cur_max)
	}
	return res
}
