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

func SliceToCatSet(s []string) (map[string]bool, error) {
	ans := make(map[string]bool)
	for _, item := range s {
		if !ItemTypes[item] {
			return nil, MegaItemTypeDoesNotExistError{InvalidType: item}
		}
		ans[item] = true
	}
	return ans, nil
}
