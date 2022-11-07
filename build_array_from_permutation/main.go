package main

func main() {

}

// ans = nums[nums[i]] when i = 1 => ans[1] = nums[2]
func buildArray(nums []int) []int {
	ans := make([]int, len(nums)) // create slice length same num

	for i := 0; i < len(nums); i++ {
		ans[i] = nums[nums[i]]
	}

	return ans
}
