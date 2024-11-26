package functions

func filterUniquePaths(paths [][]string) [][]string {
	var shortestPaths [][]string
	usedRooms := map[string]bool{}

	for _, path := range paths {
		unique := true
		for _, room := range path[1 : len(path)-1] {
			if usedRooms[room] {
				unique = false
				break
			}
		}
		if unique {
			shortestPaths = append(shortestPaths, path)
			for _, room := range path[1 : len(path)-1] {
				usedRooms[room] = true
			}
		}
	}
	return shortestPaths
}
