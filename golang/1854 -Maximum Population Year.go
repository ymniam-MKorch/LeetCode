package lc1854

func maximumPopulation(logs [][]int) int {
	var max int
	var res int
	years := make([]int, 100)
	for _, p := range logs {
		st := p[0]
		ed := p[1]

		for i := st; i < ed; i++ {
			years[i-1950]++
		}
	}

	for i, v := range years {
		if v > max {
			max = v
			res = i
		}
	}
	return res + 1950
}

func main() {

}
