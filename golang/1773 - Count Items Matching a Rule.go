func getKeyIndex(key string) int {
	switch key {
	case "type":
		return 0
	case "color":
		return 1
	case "name":
		return 2
	}
	return -1
}

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	ans := 0
	keyIndex := getKeyIndex(ruleKey)
	for _, v := range items {
		if v[keyIndex] == ruleValue {
			ans++
		}
	}
	return ans
}