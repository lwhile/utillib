package strs

// 通配符匹配

// WildcardMatch return a bool whether a str could match patter
func WildcardMatch(str, pattern string) bool {
	res := make([][]bool, len(pattern)+1)
	for i := range res {
		res[i] = make([]bool, len(str)+1)
	}

	res[0][0] = true

	for i := 1; i <= len(pattern); i++ {
		pt := pattern[i-1]
		res[i][0] = res[i-1][0] && pt == '*'
		for j := 1; j <= len(str); j++ {
			lt := str[j-1]
			if pt == '*' {
				res[i][j] = res[i][j-1] || res[i-1][j]
			} else {
				res[i][j] = res[i-1][j-1] && (pt == '?' || pt == lt)
			}
		}
	}
	return res[len(pattern)][len(str)]
}
