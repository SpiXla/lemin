package functions

// func filterUniquePaths(paths [][]string, isSorted bool) [][]string {
// 	var shortestPaths [][]string

// 	if !isSorted {
// 		for i := 0; i < len(paths); i++ {
// 			for j := i + 1; j < len(paths); j++ {
// 				if len(paths[i]) > len(paths[j]) {
// 					paths[i], paths[j] = paths[j], paths[i]
// 				}
// 			}
// 		}
// 	}

// 	usedRooms := map[string]bool{}

// 	for _, path := range paths {
// 		unique := true
// 		for _, room := range path[1 : len(path)-1] {
// 			if usedRooms[room] {
// 				unique = false
// 				break
// 			}
// 		}
// 		if unique {
// 			shortestPaths = append(shortestPaths, path)
// 			for _, room := range path[1 : len(path)-1] {
// 				usedRooms[room] = true
// 			}
// 		}
// 	}
// 	return shortestPaths
// }

func SortPaths(paths [][]string) [][]string{
	for i := 0; i < len(paths); i++ {
		for j := i + 1; j < len(paths); j++ {
			if len(paths[i]) > len(paths[j]) {
				paths[i] ,  paths[j] =  paths[j] ,  paths[i]
			}
		}
	}
	return paths
}