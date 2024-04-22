func powerFunc(p int, c int) int {
	temp := 1
	for i := 0; i < c; i++ {
		temp = temp * p
	}
	return temp
}

func isPowerOfTwo(n int) bool {
	if n == 1 {
		return true
	} else if n ==2{
        return true
    }
	count := 2
	p := 2
	temp := powerFunc(p, count)
	for temp < n {
		temp = powerFunc(p, count)
		// fmt.Println(temp)
		count++
	}
	if temp == n {
		return true
	} else {
		return false
	}
	return false
}
