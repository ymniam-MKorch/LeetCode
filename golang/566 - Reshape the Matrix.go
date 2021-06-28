func matrixReshape(mat [][]int, r int, c int) [][]int {
	m, n := len(mat), len(mat[0])
	if m*n != r*c {
		return mat
	}
	res := make([][]int, r)
	for index := range res {
		res[index] = make([]int, c)
	}
	for i := 0; i < m*n; i++ {
		res[i/c][i%c] = mat[i/n][i%n]
	}
	return res
}