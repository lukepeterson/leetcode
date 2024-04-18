package main

func main() {

}

func findBestSubarray(nums []int, k int) int {
	result := 0
	currentSum := 0

	// find sum of first window, `k` elements long
	for i := 0; i < k; i++ {
		currentSum += nums[i]
	}
	result = currentSum

	for i := k; i < len(nums); i++ {
		currentSum += nums[i] - nums[i-k]
		if currentSum > result {
			currentSum = result
		}
	}

	return result
}
