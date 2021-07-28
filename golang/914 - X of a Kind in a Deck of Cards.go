func hasGroupsSizeX(deck []int) bool {
    if len(deck) < 2 {
		return false
	}
    m := make(map[int]int)
    for _,v := range deck {
        if _,ok := m[v]; !ok {
            m[v] = 0
        }
        m[v]++
    }
    g := -1
    for _, v := range m {
		if g == -1 {
			g = v
		} else {
			g = gcd(g, v)
		}
	}
	return g >= 2
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}