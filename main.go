package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Ant struct {
	ID   int
	Room *Room
}

type Room struct {
	Name    string
	X, Y    int
	Links   []*Room
	IsStart bool
	IsEnd   bool
}

var rooms = make(map[string]*Room)

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

    numAnts, rooms, err := parseInput(scanner)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }

    fmt.Printf("Number of ants: %d\n", numAnts)

	fmt.Println("Rooms :")
	for _,room := range rooms {
		fmt.Printf("Room: %s %d %d\n", room.Name, room.X, room.Y)
	}


    // Display rooms and their connections
    fmt.Println("Rooms and their connections:")
    for name, room := range rooms {
        fmt.Printf("Room %s: ", name)
        for _, link := range room.Links {
            fmt.Printf("%s ", link.Name)
        }
        fmt.Println()
    }

    // Further logic for ant movement goes here...
}



func parseRoom(line string) (*Room, error) {
    parts := strings.Split(line, " ")
    if len(parts) != 3 {
        return nil, fmt.Errorf("invalid room format: %s", line)
    }

    x, err := strconv.Atoi(parts[1])
    if err != nil {
        return nil, fmt.Errorf("invalid X coordinate: %s", parts[1])
    }

    y, err := strconv.Atoi(parts[2])
    if err != nil {
        return nil, fmt.Errorf("invalid Y coordinate: %s", parts[2])
    }

    return &Room{
        Name:  parts[0],
        X:     x,
        Y:     y,
        Links: []*Room{}, // Initialize empty links
    }, nil
}

func parseInput(scanner *bufio.Scanner) (int, map[string]*Room, error) {
	var numAnts int
	rooms := make(map[string]*Room)
	parsingRooms := true

	for scanner.Scan() {
		line := scanner.Text()

		// Ignore comments
		if strings.HasPrefix(line, "#") {
			continue
		}

		// Parse number of ants
		if numAnts == 0 {
			ants, err := strconv.Atoi(line)
			if err != nil {
				return 0, nil, fmt.Errorf("invalid number of ants: %v", err)
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
				room, err := parseRoom(line)
				if err != nil {
					return 0, nil, err
				}
				rooms[room.Name] = room
				continue
			}
		}

		// Parse tunnels
		if !parsingRooms {
			err := parseTunnel(line, rooms)
			if err != nil {
				return 0, nil, err
			}
		}
	}

	return numAnts, rooms, nil
}

func parseTunnel(line string, rooms map[string]*Room) error {
    parts := strings.Split(line, "-")
    if len(parts) != 2 {
        return fmt.Errorf("invalid tunnel format: %s", line)
    }

    room1Name := parts[0]
    room2Name := parts[1]

    // Check if the rooms exist in the map
    room1, exists1 := rooms[room1Name]
    room2, exists2 := rooms[room2Name]

    if !exists1 || !exists2 {
        return fmt.Errorf("unknown room in tunnel: %s", line)
    }

    // Create two-way link
    room1.Links = append(room1.Links, room2)
    room2.Links = append(room2.Links, room1)

    return nil
}
