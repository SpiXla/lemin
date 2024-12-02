package input

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
	X, Y    int
	IsStart bool
	IsEnd   bool
}

func ParseInput(filename string) (string, int, map[string]*Room, map[string][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", 0, nil, nil, fmt.Errorf("could not open file %s", filename)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fileContent := ""
	rooms := make(map[string]*Room)
	connections := make(map[string][]string)
	var numAnts int

	parsingRooms, isStart, isEnd := true, false, false
	foundstart, foundend := false, false

	for scanner.Scan() {
		line := scanner.Text()
		fileContent += line + "\n"
		switch {
		case line == "":
			continue
		case numAnts == 0:
			numAnts, err = strconv.Atoi(line)
			if err != nil {
				return "", 0, nil, nil, errors.New("invalid data format, invalid number of ants")
			}
		case strings.HasPrefix(line, "#"):
			if line == "##end" && !foundstart {
				foundstart = true
				isEnd = true
			} else if line == "##start" && !foundend {
				foundend = true
				isStart = true
			} else if line != "##end" && line != "##start" && len(strings.Fields(line)) != 3{
				continue
			} else if len(strings.Fields(line)) == 3 {
				return "", 0, nil, nil, errors.New("invalid data format, invalid room name")
			} else {
				return "", 0, nil, nil, errors.New("invalid data format, invalid start or end room")
			}
		case parsingRooms && strings.Contains(line, "-"):
			parsingRooms = false
			parseTunnel(line, connections)
		case !parsingRooms && !strings.Contains(line, "-"):
			return "", 0, nil, nil, errors.New("invalid data format, invalid links")
		case parsingRooms:
			name, x, y, err := RoomParams(line)
			if err != nil {
				return "", 0, nil, nil, err
			}
			room := &Room{Name: name, X: x, Y: y, IsStart: isStart, IsEnd: isEnd}
			rooms[room.Name] = room
			isStart, isEnd = false, false
		default:
			err := parseTunnel(line, connections)
			if err != nil {
				return "", 0, nil, nil, err
			}
		}
	}

	return fileContent, numAnts, rooms, connections, nil
}
