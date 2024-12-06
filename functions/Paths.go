package functions

import (
	"errors"
	input "lemin/Input"
)

func PathFindingLogic(rooms map[string]input.Room, connections map[string][]string) ([][][]string, error) {
	start, end := FindStartEnd(rooms)

	if start == "" || end == "" {
		return nil, errors.New("start or end room not defined")
	}

	shortestPath, ConnWithoutStart := FindShortestPaths(start, end, connections)
	SortPaths(&shortestPath)

	Groups := AllGroups(start,end,shortestPath,ConnWithoutStart,connections)

	return Groups , nil
}

func FindShortestPaths(start, end string, connections map[string][]string) ([][]string, map[string][]string) {
	shortestPath := [][]string{}
	visited := map[string]bool{}

	// DepthFirstSearch([]string{start}, start, end, &paths, visited, connections)
	ConnWithoutStart := RemoveStart(connections, start)
	// fmt.Println(connections[start])
	for _, StartConnection := range connections[start] {
		path := BFS(ConnWithoutStart, StartConnection, end)
		shortestPath = append(shortestPath, path)
	}
	SortPaths(&shortestPath)
	for _, path := range shortestPath {
		p := path[:len(path)-1]
		for i, room := range p {
			if !visited[room] {
				visited[room] = true
			} else {
				roomBefore := p[i-1]
				if len(connections[roomBefore]) > 2 {
					connections = RemoveConn(connections, roomBefore, room)
					path := BFS(ConnWithoutStart, p[0], end)
					shortestPath = append(shortestPath, path)
				}
			}
		}
	}
	return shortestPath, ConnWithoutStart
}
