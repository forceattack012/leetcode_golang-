package main

func main() {

}

func getConcatenation(nums []int) []int {
	n := len(nums)
	ans := make([]int, 0, n*2)

	for i := 0; i < n; i++ {
		ans[i] = nums[i]
		ans[n+i] = nums[i]
	}

	return ans
}
