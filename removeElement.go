func removeElement(nums []int, val int) int {
    a:=nums
    t:=val
    for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if a[j] == t {
				// fmt.Println(a[i], "a[i]\n")
				// fmt.Println(a[j], "a[j]\n")
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	var result []int
	// k := 0

	for _, val := range a {
		if val != t {
			result = append(result, val)
		}
	}
	return len(result)
}
