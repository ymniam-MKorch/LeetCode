import "strings"

func findWords(words []string) []string {
	rows := []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
	res := make([]string, 0)
	for _, s := range words {
		if len(s) == 0 {
			continue
		}
		lowerS := strings.ToLower(s)
		oneRow := false
		for _, r := range rows {
			if strings.ContainsAny(lowerS, r) {
				oneRow = !oneRow
				if !oneRow {
					break
				}
			}
		}
		if oneRow {
			res = append(res, s)
		}
	}
	return res
}