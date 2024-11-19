package functions



func filterUniquePaths(paths [][]string) [][]string {
	shortestLength := len(paths[0])
	for _, path := range paths {
		if len(path) < shortestLength {
			shortestLength = len(path)
		}
	}

	shortestPaths, usedRooms := [][]string{}, map[string]bool{}
	for _, path := range paths {
		if len(path) == shortestLength {
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
	}
	return shortestPaths
}
