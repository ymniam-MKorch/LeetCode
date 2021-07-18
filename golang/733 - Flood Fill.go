func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	color := image[sr][sc]
	if newColor == color {
		return image
	}
	dfs733(image, sr, sc, newColor)
	return image
}

func dfs733(image [][]int, sr int, sc int, newColor int) {
    if image[sr][sc] == newColor {
		return
	}
    oldColor := image[sr][sc]
    image[sr][sc] = newColor
    if sr > 0 && image[sr-1][sc] == oldColor {
        dfs733(image, sr-1, sc, newColor)
    }
    if sc < len(image[0])-1 && image[sr][sc+1] == oldColor {
        dfs733(image, sr, sc+1, newColor)
    }
    if sc > 0 && image[sr][sc-1] == oldColor {
        dfs733(image, sr, sc-1, newColor)
    }
    if sr < len(image)-1 && image[sr+1][sc] == oldColor {
        dfs733(image, sr+1, sc, newColor)
    }
    return
}