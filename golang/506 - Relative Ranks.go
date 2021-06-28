import (
	"strconv"
)

func findRelativeRanks(score []int) []string {
	sc := make([]int, len(score))
	copy(sc, score)
	sort.Sort(sort.Reverse(sort.IntSlice(sc)))
	m := make(map[int]int)
	for i := 0; i < len(sc); i++ {
		m[sc[i]] = i
	}
	var res []string
	for _, v := range score {
		if m[v] == 0 {
			res = append(res, "Gold Medal")
		} else if m[v] == 1 {
			res = append(res, "Silver Medal")
		} else if m[v] == 2 {
			res = append(res, "Bronze Medal")
		} else {
			res = append(res, strconv.Itoa(m[v]+1))
		}
	}
	return res
}