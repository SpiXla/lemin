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

type SHelpVaraibles struct {
	Line         string
	Foundstart   bool
	Foundend     bool
	ParsingRooms bool
	Name         string
	X, Y         int
	IsStart      bool
	IsEnd        bool
	Err          error
}

func ParseInput(filename string) (string, int, map[string]Room, map[string][]string, error) {
	var Tools SHelpVaraibles

	file, err := os.Open(filename)
	if err != nil {
		return "", 0, nil, nil, fmt.Errorf("could not open file %s", filename)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fileContent := ""
	rooms := make(map[string]Room)

	connections := make(map[string][]string)
	var numAnts int

	Tools.ParsingRooms, Tools.IsStart, Tools.IsEnd = true, false, false
	Tools.Foundstart, Tools.Foundend = false, false

	for scanner.Scan() {
		Tools.Line = scanner.Text()
		if Tools.Line == "" {
			continue
		}

		fileContent += Tools.Line + "\n"
		switch {

		case numAnts == 0:
			numAnts, err = strconv.Atoi(Tools.Line)
			if err != nil || numAnts < 0 {
				return "", 0, nil, nil, errors.New("invalid data format, wrong number of ants")
			}
		case strings.HasPrefix(Tools.Line, "#"):
			
			if err := CommentHandler(&Tools); err == nil {
				continue
			}
			
		case Tools.ParsingRooms && strings.Contains(Tools.Line, "-"):
			Tools.ParsingRooms = false
			parseTunnel(Tools.Line, connections)

		case !Tools.ParsingRooms && !strings.Contains(Tools.Line, "-"):
			return "", 0, nil, nil, errors.New("invalid data format, invalid links")

		case Tools.ParsingRooms:
			
			Tools.Name, Tools.X, Tools.Y, Tools.Err = RoomParams(Tools.Line)
			if Tools.Err != nil {
				return "", 0, nil, nil, Tools.Err
			}
			room := &Room{Name: Tools.Name, X: Tools.X, Y: Tools.Y, IsStart: Tools.IsStart, IsEnd: Tools.IsEnd}
			rooms[room.Name] = *room
			Tools.IsStart, Tools.IsEnd = false, false
			
		default:
			err := parseTunnel(Tools.Line, connections)
			if err != nil {
				return "", 0, nil, nil, err
			}
		}
	}
	return fileContent, numAnts, rooms, connections, nil
}

func CommentHandler(Tools *SHelpVaraibles) error {
	
	if Tools.Line == "##end" && !Tools.Foundstart {
		Tools.IsEnd = true
	} else if Tools.Line == "##start" && !Tools.Foundend {
		Tools.IsStart = true
	} else if Tools.Line != "##end" && Tools.Line != "##start" && len(strings.Fields(Tools.Line)) != 3 {
		return nil
	} else if len(strings.Fields(Tools.Line)) == 3 {
		return errors.New("invalid data format, invalid room name")
	}
	return errors.New("invalid data format, invalid start or end room")
}
