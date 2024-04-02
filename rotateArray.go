func rotate(nums []int, k int)  {
	k = k % len(nums)
	reverse(0, len(nums)-1, nums)
	reverse(0, k-1, nums)
	reverse(k, len(nums)-1, nums)
}

func reverse(start int, end int, nums []int) {
	for start < end {
		nums[start], nums[end], start,end = nums[end], nums[start], start+1, end -1
	}
}
