package leetcode1886

func findRotation(mat [][]int, target [][]int) bool {
	n := len(mat)
	var res [4]bool = [4]bool{true, true, true, true}
	for i, _ := range mat {
		for j, _ := range target {
			if mat[i][j] != target[i][j] {
				res[0] = false
			}
			if mat[i][j] != target[j][n-i-1] {
				res[1] = false
			}
			if mat[i][j] != target[n-i-1][n-j-1] {
				res[2] = false
			}
			if mat[i][j] != target[n-j-1][i] {
				res[3] = false
			}
		}
	}
	return res[0] || res[1] || res[2] || res[3]
}

func main() {

}
