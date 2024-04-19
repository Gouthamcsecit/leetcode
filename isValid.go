func isValid(s string) bool {
    if len(s) == 0 || (len(s)%2) == 1 {
		return false
	}

	pairs := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stack := []rune{}
	for _, val := range s {
		if _, ok := pairs[val]; ok {
			stack = append(stack, val)
		} else if len(stack) == 0 || pairs[stack[len(stack)-1]] != val {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}