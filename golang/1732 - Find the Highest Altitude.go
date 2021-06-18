func largestAltitude(gain []int) int {
	var alt, maxAlt int
	for _, v := range gain {
		alt += v
		if alt > maxAlt {
			maxAlt = alt
		}
	}
	return maxAlt
}