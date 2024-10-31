package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Room struct {
	Name    string
	IsStart bool
	IsEnd   bool
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}

	numAnts, rooms, connections, err := parseInput(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Number of ants: %d\n", numAnts)
	for name, room := range rooms {
		fmt.Printf("Room: %s %s\n", name, roomStatus(room))
	}
	for room, links := range connections {
		fmt.Printf("Room %s: %v\n", room, links)
	}

	paths, err := findUniquePaths(rooms, connections)
	if err != nil {
		fmt.Println("Error finding paths:", err)
		return
	}

	fmt.Println("All unique paths from start to end:")
	for _, path := range paths {
		fmt.Println(path)
	}
}

func roomStatus(room *Room) string {
	switch {
	case room.IsStart:
		return "(Start)"
	case room.IsEnd:
		return "(End)"
	default:
		return ""
	}
}

func parseInput(filename string) (int, map[string]*Room, map[string][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, nil, nil, fmt.Errorf("could not open file %s", filename)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rooms := make(map[string]*Room)
	connections := make(map[string][]string)
	var numAnts int
	parsingRooms, isStart, isEnd := true, false, false

	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case strings.HasPrefix(line, "#"):
			isStart, isEnd = line == "##start", line == "##end"
		case numAnts == 0:
			numAnts, err = strconv.Atoi(line)
			if err != nil {
				return 0, nil, nil, fmt.Errorf("invalid number of ants: %v", err)
			}
		case parsingRooms && strings.Contains(line, "-"):
			parsingRooms = false
			parseTunnel(line, connections)
		case parsingRooms:
			room := &Room{Name: strings.Fields(line)[0], IsStart: isStart, IsEnd: isEnd}
			rooms[room.Name] = room
			isStart, isEnd = false, false
		default:
			parseTunnel(line, connections)
		}
	}

	return numAnts, rooms, connections, nil
}

func parseTunnel(line string, connections map[string][]string) {
	parts := strings.Split(line, "-")
	connections[parts[0]] = append(connections[parts[0]], parts[1])
	connections[parts[1]] = append(connections[parts[1]], parts[0])
}

func findUniquePaths(rooms map[string]*Room, connections map[string][]string) ([][]string, error) {
	start, end := findStartEnd(rooms)
	if start == "" || end == "" {
		return nil, fmt.Errorf("start or end room not defined")
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
	return filterUniquePaths(paths), nil
}

func findStartEnd(rooms map[string]*Room) (string, string) {
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
