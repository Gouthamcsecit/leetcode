func majorityElement(nums []int) int {
    mymap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		mymap[nums[i]] = mymap[nums[i]] + 1
	}
	count := 0
	// prev :=0
	for i, val := range mymap {
		if val > (len(nums) / 2) {
			count = i
		}
	}
	return count
}
