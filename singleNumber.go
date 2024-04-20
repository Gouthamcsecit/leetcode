func singleNumber(nums []int) int {
    myMap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if myMap[nums[i]] != 0 {
			myMap[nums[i]]++
		} else {
			myMap[nums[i]] = 1
		}
	}
	for j := 0; j < len(nums); j++ {
		// fmt.Println(j, val)
		if myMap[nums[j]] == 1 {
			return nums[j]
		}
	}
	return 0
}
