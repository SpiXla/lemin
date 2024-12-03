package functions

import (
	"errors"

	input "lemin/Input"
)

func GetNumberOfConnections(connections map[string][]string) int {
	var Counter int
	for _, value := range connections {
		Counter += len(value)
	}
	return Counter
}

func FindUniquePaths(rooms map[string]input.Room, connections map[string][]string) ([][]string, [][]string, error) {
	start, end := FindStartEnd(rooms)

	if start == "" || end == "" {
		return nil, nil, errors.New("start or end room not defined")
	}
	var paths [][]string
	visited := map[string]bool{}

	if Num := GetNumberOfConnections(connections); Num > 200 {
		return nil, nil, errors.New("unsupported graph")
	}

	DepthFirstSearch([]string{start}, start, end, &paths, visited, connections)
	if len(paths) == 0 {
		return nil, nil, errors.New("invalid data format, no path found")
	}

	return filterUniquePaths(paths, true), filterUniquePaths(paths, false), nil
}

func DepthFirstSearch(path []string, start, end string, paths *[][]string, visited map[string]bool, connections map[string][]string) {

	room := path[len(path)-1]

	if room == end {
		*paths = append(*paths, append([]string{}, path...))
	}
	visited[room] = true
	for _, neighbor := range connections[room] {
		if !visited[neighbor] {
			DepthFirstSearch(append(path, neighbor), start, end, paths, visited, connections)
		}
	}
	visited[room] = false
}

func FindStartEnd(rooms map[string]input.Room) (string, string) {
	var start, end string
	for name, room := range rooms {
		if room.IsStart {
			start = name
		} else if room.IsEnd {
			end = name
		}
	}
	return start, end
}
