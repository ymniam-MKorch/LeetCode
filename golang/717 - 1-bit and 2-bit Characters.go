func isOneBitCharacter(bits []int) bool {
    var i int
    for ;i < len(bits)-1; i++ {
        if bits[i] == 1 {
            i++
        }
    }
    return i == len(bits)-1
}