package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Room struct {
	Name    string
	X, Y    string
	IsStart bool
	IsEnd   bool
}
type Ant struct {
	Current string
	Num     int
	path    []string
}

var (
	seencor = make(map[string]bool)
	seename = make(map[string]bool)
)

var linkstart, linkend = false, false

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
	start, end := findStartEnd(rooms)
	for a := range connections {
		if a == start {
			linkstart = true
		} else if a == end {
			linkend = true
		}
	}
	if !linkend || !linkstart {
		fmt.Println("start or end isnt linked")
		return
	}

	fmt.Printf("Number of ants: %d\n", numAnts)

	paths, err := findUniquePaths(rooms, connections)
	if err != nil {
		fmt.Println("Error finding paths:", err)
		return
	}

	fmt.Println("All unique paths from start to end:")
	for _, path := range paths {
		fmt.Println(path)
	}

	Antmovement(numAnts, start, end, paths)
}

func Antmovement(ants int, start, end string, paths [][]string) {
	Ants := make([]Ant, ants)

	for i := 1; i <= ants; i++ {
		Ants[i-1].Num = i
		Ants[i-1].path = paths[i%len(paths)]
		Ants[i-1].Current = Ants[i-1].path[0]
	}
	
	moves := []string{}
	
	for _, ant := range Ants {
		for _, p := range ant.path[1:] {
			num := ant.Num
			moves = append(moves, fmt.Sprintf("L%d-%v", num, p))
		}
	}
	fmt.Println(Ants)
	fmt.Println(moves)
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
	// var start, end string
	parsingRooms, isStart, isEnd := true, false, false
	foundstart, foundend := false, false
	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case line == "":
			return 0, nil, nil, errors.New("invalid data (empty lines)")
		case numAnts == 0:
			numAnts, err = strconv.Atoi(line)
			if err != nil {
				return 0, nil, nil, errors.New("invalid number of ants")
			}
		case strings.HasPrefix(line, "#"):
			if line == "##end" && !foundstart {
				foundstart = true
				isEnd = true
			} else if line == "##start" && !foundend {
				foundend = true
				isStart = true
			} else {
				return 0, nil, nil, errors.New("invalid start or end")
			}
		case parsingRooms && strings.Contains(line, "-"):
			parsingRooms = false
			parseTunnel(line, connections)
		case !parsingRooms && !strings.Contains(line, "-"):
			return 0, nil, nil, errors.New("invalid links")
		case parsingRooms:
			name, x, y, err := RoomParams(line)
			if err != nil {
				return 0, nil, nil, err
			}
			room := &Room{Name: name, X: x, Y: y, IsStart: isStart, IsEnd: isEnd}
			rooms[room.Name] = room
			isStart, isEnd = false, false
		default:
			err := parseTunnel(line, connections)
			if err != nil {
				return 0, nil, nil, err
			}
		}
	}

	return numAnts, rooms, connections, nil
}

func RoomParams(line string) (string, string, string, error) {
	info := strings.Fields(line)
	if len(info) != 3 {
		return "", "", "", errors.New("wrong room data")
	}
	name := info[0]
	if seename[name] {
		return "", "", "", errors.New("wrong name")
	} else if !seename[name] {
		seename[name] = true
	}
	coor := strings.Join(info[1:], " ")
	if seencor[coor] {
		return "", "", "", errors.New("wrong coordinates")
	} else if !seencor[coor] {
		seencor[coor] = true
	}
	x := info[1]
	y := info[2]
	return name, x, y, nil
}

func parseTunnel(line string, connections map[string][]string) error {
	if len(line) != 3 {
		return errors.New("invalid links")
	}
	parts := strings.Split(line, "-")
	if parts[0] == parts[1] {
		return errors.New("invalid links")
	}
	connections[parts[0]] = append(connections[parts[0]], parts[1])
	connections[parts[1]] = append(connections[parts[1]], parts[0])
	return nil
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
