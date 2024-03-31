func removeDuplicate(nums []int) int {
	mymap := map[int]int{}
	count := 0
	for i := 0; i < len(nums); i++ {
		if mymap[nums[i]] < 2 {
			mymap[nums[i]] = mymap[nums[i]] + 1
			nums[count] = nums[i]
			count++
		}
	}
	return count
}
