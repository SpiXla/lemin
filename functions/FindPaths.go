package functions

import (
	"errors"
	"fmt"

	input "lemin/Input"
)

func FindUniquePaths(rooms map[string]input.Room, connections map[string][]string) ([][]string,[][][]string, error) {
	start, end := FindStartEnd(rooms)
	// var paths [][]string 
	visited := map[string]bool{}

	if start == "" || end == "" {
		return nil,nil, errors.New("start or end room not defined")
	}

	shortestPath := [][]string{}

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

	SortPaths(&shortestPath)
	AllGroups := [][][]string{}
	for _, path := range shortestPath {
		startConn := []string{}
		for _, room := range connections[start] {
			if room != path[0] {
				startConn = append(startConn, room)
			}
		}
		Groups := GroupsPaths(path, startConn, ConnWithoutStart, end)
		AllGroups = append(AllGroups, Groups)
		// fmt.Println(Groups)
	}
    for _, group := range AllGroups {
        roomSet := make(map[string]bool)
        for _, path := range group {
            for _, room := range path {
                if room != end {
                    if roomSet[room] {
                        fmt.Printf("Room %s is duplicated!\n", room)
                    }
                    roomSet[room] = true
                }
            }
        }
    }

	return shortestPath,AllGroups, nil
}

func GroupsPaths(path, roomsLinkedWStart []string, connWithoutStart map[string][]string, end string) [][]string {
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
