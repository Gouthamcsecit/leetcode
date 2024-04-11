func isAnagram(s string, t string) bool {
    sNew := strings.Split(s, "")
	tNew := strings.Split(t, "")
	sort.Strings(sNew)
	sort.Strings(tNew)
	sNew1 := strings.Join(sNew, "")
	tNew1 := strings.Join(tNew, "")

n1 := len(sNew1)

n2:= len(tNew1)
length := 0
if n1 !=n2{
    return false
} else {
    length = n1
}
	fmt.Println(sNew1, tNew1)
	for i := 0; i < length; i++ {
		if string(sNew1[i]) != string(tNew1[i]) {
			fmt.Println(string(sNew1[i]), string(tNew1[i]))
			return false
		}
	}
	return true
}
