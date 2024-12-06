package functions


func BFS(graph map[string][]string, start, end string) []string {
	queue := [][]string{{start}}
	visited := make(map[string]bool)
	visited[start] = true

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]

		node := path[len(path)-1]
		if node == end {
			return path
		}
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				newPath := append([]string{}, path...)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
			}
		}
	}
	return nil
}


func BfsGroups(graph map[string][]string, start, end string, globalVisited *map[string]bool) [][]string {
    localVisited := make(map[string]bool)
    for room, isVisited := range *globalVisited {
        localVisited[room] = isVisited
    }

    queue := [][]string{{start}}
    groups := [][]string{}

    for len(queue) > 0 {
        currentPath := queue[0]
        queue = queue[1:]

        currentRoom := currentPath[len(currentPath)-1]

        if currentRoom == end {
            isDuplicate := false
            visited := make(map[string]bool)
            
            for _, room := range currentPath {
                if room != end && visited[room] {
                    isDuplicate = true
                    break
                }
                visited[room] = true
            }

            if !isDuplicate {
                groups = append(groups, currentPath)
            }
            continue
        }

        for _, neighbor := range graph[currentRoom] {
            if !localVisited[neighbor] {
                newPath := append([]string{}, currentPath...) 
                newPath = append(newPath, neighbor)
                
                visited := make(map[string]bool)
                hasDuplicate := false
                for _, room := range newPath {
                    if room != end && visited[room] {
                        hasDuplicate = true
                        break
                    }
                    visited[room] = true
                }

                if !hasDuplicate {
                    queue = append(queue, newPath)
                    localVisited[neighbor] = true
                }
            }
        }
    }

    for _, path := range groups {
        for _, room := range path {
            if room != end {
                (*globalVisited)[room] = true
            }
        }
    }

    return groups
}
