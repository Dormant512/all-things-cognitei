package util

func MapToSlice(m map[string]bool) []string {
	ans := make([]string, 0)
	for key, val := range m {
		if val {
			ans = append(ans, key)
		}
	}
	return ans
}
