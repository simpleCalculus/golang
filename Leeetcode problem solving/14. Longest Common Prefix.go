func longestCommonPrefix(strs []string) string {
	s := strs[0]
	var ans string
	for i, _ := range s {
		flag := false
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) <= i || s[i] != strs[j][i] {
				flag = true
				break
			}
		}
		if flag {
			return ans
		} else {
            ans = ans + string(s[i])
		}
	}
	return ans
}
