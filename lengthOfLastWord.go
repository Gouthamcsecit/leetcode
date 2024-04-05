func lengthOfLastWord(s string) int {
    s = strings.Trim(s, " ")
	sm := strings.Split(s, " ")
	count := len(string(sm[len(sm)-1]))
	return count
}
