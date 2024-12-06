package functions

func GroupPaths(path, roomsLinkedWStart []string, connWithoutStart map[string][]string, end string) [][]string {
	globalVisited := make(map[string]bool)
	groups := [][]string{}

	for _, room := range path {
		if room != end {
			globalVisited[room] = true
		}
	}
	groups = append(groups, path)

	for _, pathStart := range roomsLinkedWStart {
		if !globalVisited[pathStart] {
			newGroup := BfsGroups(connWithoutStart, pathStart, end, &globalVisited)
			if len(newGroup) > 0 {
				groups = append(groups, newGroup...)
			}
		}
	}

	return groups
}

func AllGroups(start,end string, shortestPath [][]string,ConnWithoutStart,connections map[string][]string) [][][]string {
	AllGroups := [][][]string{}
	for _, path := range shortestPath {
		startConn := []string{}
		for _, room := range connections[start] {
			if room != path[0] {
				startConn = append(startConn, room)
			}
		}
		Groups := GroupPaths(path, startConn, ConnWithoutStart, end)
		AllGroups = append(AllGroups, Groups)
		// fmt.Println(Groups)
	}
	return AllGroups
}
