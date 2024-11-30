package functions

import (
	"fmt"


	input "lemin/Input"
)

func FindUniquePaths(rooms map[string]*input.Room, connections map[string][]string) ([][]string, [][]string, error) {
	start, end := FindStartEnd(rooms)
	if start == "" || end == "" {
		return nil, nil, fmt.Errorf("start or end room not defined")
	}

	var paths [][]string
	visited := map[string]bool{}
	var dfs func(path []string)
	
	dfs = func(path []string) {
		room := path[len(path)-1]
		if room == end {
			paths = append(paths, append([]string{}, path...))
			return
		}
		visited[room] = true
		for _, neighbor := range connections[room] {
			if !visited[neighbor] {
				dfs(append(path, neighbor))
			}
		}
		visited[room] = false
	}
	dfs([]string{start})
	return filterUniquePaths(paths, true), filterUniquePaths(paths, false), nil
}

func FindStartEnd(rooms map[string]*input.Room) (string, string) {
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
