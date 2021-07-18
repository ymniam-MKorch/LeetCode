func nextGreatestLetter(letters []byte, target byte) byte {
    l,r := 0,len(letters)-1
    for l <= r {
        mid := (l+r)/2
        if letters[mid] > target {
            r = mid - 1
        } else {
            l = mid + 1
        }
    }
    find := letters[l%len(letters)]
	if find <= target {
		return letters[0]
	}
	return find
}