package main

func runningSum(nums []int) []int {
	ans := make([]int, len(nums))
	ans[0] = nums[0]
	n := len(nums)

	for i := 1; i <= n-1; i++ {
		ans[i] = ans[i-1] + nums[i]
	}

	return ans
}
