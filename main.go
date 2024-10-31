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

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: could not open file %s\n", filename)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	numAnts, rooms, connections, err := parseInput(scanner)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Number of ants: %d\n", numAnts)

	fmt.Println("Rooms:")
	for _, room := range rooms {
		status := ""
		if room.IsStart {
			status = "(Start)"
		} else if room.IsEnd {
			status = "(End)"
		}
		fmt.Printf("Room: %s %s\n", room.Name, status)
	}

	fmt.Println("Rooms and their connections:")
	for room, links := range connections {
		fmt.Printf("Room %s: %v\n", room, links)
	}

	allPaths, err := findAllPaths(rooms, connections)
	if err != nil {
		fmt.Printf("Error finding paths: %v\n", err)
	} else {
		fmt.Println("All paths from start to end:")
		for _, path := range allPaths {
			fmt.Println(path)
		}
	}

	// Further logic for ant movement goes here...
}

func parseRoom(line string, isStart bool, isEnd bool) (*Room, error) {
	parts := strings.Split(line, " ")
	if len(parts) < 1 {
		return nil, fmt.Errorf("invalid room format: %s", line)
	}

	return &Room{
		Name:    parts[0],
		IsStart: isStart,
		IsEnd:   isEnd,
	}, nil
}

func parseInput(scanner *bufio.Scanner) (int, map[string]*Room, map[string][]string, error) {
	var numAnts int
	rooms := make(map[string]*Room)
	connections := make(map[string][]string)
	parsingRooms := true
	isStart := false
	isEnd := false

	for scanner.Scan() {
		line := scanner.Text()

		// Ignore comments
		if strings.HasPrefix(line, "#") {
			if line == "##start" {
				isStart = true
			} else if line == "##end" {
				isEnd = true
			}
			continue
		}

		// Parse number of ants
		if numAnts == 0 {
			ants, err := strconv.Atoi(line)
			if err != nil {
				return 0, nil, nil, fmt.Errorf("invalid number of ants: %v", err)
			}
			numAnts = ants
			continue
		}

		// Parse rooms and tunnels
		if parsingRooms {
			if strings.Contains(line, "-") {
				// Start parsing tunnels
				parsingRooms = false
			} else {
				// Parse room
				room, err := parseRoom(line, isStart, isEnd)
				if err != nil {
					return 0, nil, nil, err
				}
				rooms[room.Name] = room

				// Reset start/end flags
				isStart = false
				isEnd = false
				continue
			}
		}

		// Parse tunnels
		if !parsingRooms {
			err := parseTunnel(line, connections)
			if err != nil {
				return 0, nil, nil, err
			}
		}
	}

	return numAnts, rooms, connections, nil
}

func parseTunnel(line string, connections map[string][]string) error {
	parts := strings.Split(line, "-")
	if len(parts) != 2 {
		return fmt.Errorf("invalid tunnel format: %s", line)
	}

	room1 := parts[0]
	room2 := parts[1]

	// Add two-way connections
	connections[room1] = append(connections[room1], room2)
	connections[room2] = append(connections[room2], room1)

	return nil
}

func findAllPaths(rooms map[string]*Room, connections map[string][]string) ([][]string, error) {
	var startRoom, endRoom *Room

	// Locate the start and end rooms
	for _, room := range rooms {
		if room.IsStart {
			startRoom = room
		} else if room.IsEnd {
			endRoom = room
		}
	}

	if startRoom == nil || endRoom == nil {
		return nil, fmt.Errorf("start or end room not defined")
	}

	allPaths := [][]string{}
	visited := make(map[string]bool)

	var dfs func(current string, path []string)
	dfs = func(current string, path []string) {
		path = append(path, current)
		if current == endRoom.Name {
			// Found a complete path
			// Append a copy of the path to allPaths
			newPath := make([]string, len(path))
			copy(newPath, path)
			allPaths = append(allPaths, newPath)
			return
		}

		visited[current] = true
		for _, neighbor := range connections[current] {
			if !visited[neighbor] {
				dfs(neighbor, path)
			}
		}
		visited[current] = false
	}

	dfs(startRoom.Name, []string{})
	return allPaths, nil
}
